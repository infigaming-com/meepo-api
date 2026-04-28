# Meepo Business Notification Pipeline v1

This document is the contract for the Meepo business-notification
pipeline. PR reviewers should reference it when reviewing any change
that sends an email from a backend-triggered code path.

## Scope

All "backend-triggered operational / alert" emails MUST go through this
pipeline. Examples:

- balance / cost-rate alerts (operator, sub-account)
- large-withdrawal review notifications
- monthly-invoice-ready notifications
- KYC pending-review notifications
- system anomaly notifications

### Out of scope (continue to send synchronously)

- OTP / 2FA codes — a user is waiting on the wire.
- Account-creation / activation emails — same.

The rule of thumb: if a human is staring at a UI waiting for the result,
sync send is fine. If the trigger is a background event and the user is
not actively waiting, it goes through this pipeline.

## How it works

```
business service                      meepo-push-service
─────────────────                     ──────────────────
  detect condition                       subscribe topic
  look up emails                            │
  build params map                          ▼
  publish ─────────► pubsub ───────► dedup (Redis SET NX EX)
                                            │
                                            ▼
                                       render template
                                            │
                                            ▼
                                       send via internal SendEmail
                                            │
                                            ▼
                                       audit + retry + dead-letter
```

## How to send a notification (business side)

```go
import (
    "github.com/infigaming-com/meepo-api/pkg/notification"
    notification_pb "github.com/infigaming-com/meepo-api/notification/service/v1"
)

emails, err := userSvc.GetCompanyAdminEmails(ctx, ...) // or whatever lookup applies
// ... look up operatorName / other params ...

useSystemDefault := true
event := &notification_pb.BusinessAlertEvent{
    Template: notification.TemplateSubAccountBalanceAlert,
    Params: map[string]string{
        "operatorName": operatorName,
        "balance":      balance.String(),
        // ... whatever the template requires
    },
    To:               emails,
    OperatorContext:  operatorContext,
    UseSystemDefault: &useSystemDefault,
    DedupKey:         fmt.Sprintf("sub_account_balance_alert:%d:%s:%s", realOpId, productType, currency),
    DedupWindow:      notification_pb.DedupWindow_DEDUP_WINDOW_UNTIL_UTC_MIDNIGHT,
}
pubsubSvc.Pub(ctx, "notifications.business_alert", event)
```

## Adding a new template (3 coordinated PRs)

1. **push-service**: drop `templates/email/<name>_template.txt` and
   `<name>_template.html`, register the name in the internal template
   registry with its subject format string.
2. **meepo-api**: add the corresponding constant in
   `pkg/notification/templates.go`. Cut a tag per the meepo-api release
   process.
3. **business service**: bump the meepo-api dep, build the params map,
   publish the event.

## Required event fields

- `template` — must be a constant from `meepo-api/pkg/notification`. Do
  not inline the literal string.
- `params` — every placeholder declared in the template files MUST be
  provided. Missing keys are passed through verbatim and will appear
  literally in the rendered email; push-service logs a warning.
- `to` — non-empty. Business side resolves recipients (no resolver
  abstraction in v1).
- `operator_context` — at minimum the `operator_id` push-service binds
  the sender on. push-service uses this for sender-binding selection
  and audit; missing fields fall back to system defaults.

## Optional event fields

- `from` — empty defaults to system sender.
- `user_id` — independent semantic from operator_context; pass 0 when
  N/A.
- `use_system_default` — defaults to true. Set false to use the
  operator's bound sender configuration.
- `dedup_key` + `dedup_window` — see below.

## Dedup contract

- `dedup_key` is the business identity of "one alert occurrence."
  Include every dimension that should fire independently (operator,
  product, currency, etc.). Exclude time and changing values (balance,
  amount, percentage) — those would defeat dedup.
- `dedup_window` selects how long the key stays claimed.
- Semantics: **claim-on-attempt** (`SET NX EX` before sending). Exactly
  one send attempt per window per key. A send failure inside a window
  means the alert is lost for that window — see the dead-letter table
  to replay.
- This is intentionally at-most-once per window. If a future alert
  family needs at-least-once, we add a `dedup_strategy` enum field;
  v1 does not need it.

`dedup_key` empty OR `dedup_window` `UNSPECIFIED` ⇒ no dedup, every
event sends.

### Choosing a dedup_window

| Window | Use when |
|---|---|
| `UNSPECIFIED` (no dedup) | one-shot events (welcome email, single-row business decision) |
| `FIVE_MINUTES` | guard against threshold-crossing chatter |
| `ONE_HOUR` | important alerts that should not spam |
| `UNTIL_UTC_MIDNIGHT` | "tell me once today" — most operational alerts |
| `SLIDING_24H` | rolling 24h instead of calendar day |

## Forbidden patterns

- Embedding HTML/text email bodies inside a business service repo.
- Mounting `/email_templates` as a k8s volume from a business service.
- Calling `push.SendEmail` directly for business notifications. (OTP /
  account activation are the only exceptions.)
- Inlining template literal strings in `event.Template`. Always
  reference a constant from this package.

## Outliers tracked for migration

These predate this pipeline and are listed as backlog. They will be
migrated in their own PRs.

- `meepo-operator-service/templates/email/balance_alert.html` +
  `internal/biz/balance_alert_helpers.go` — migrate to a
  `TemplateOperatorBalanceAlert` constant + `BusinessAlertEvent`
  publish; delete the embedded HTML and Redis dedup helpers.
- `meepo-backoffice-service`, `meepo-affiliate-service`,
  `meepo-user-service` callers of `email_util.GetEmailTemplate` + sync
  `push.SendEmail` for business notifications — migrate to publish.
  Remove the `/email_templates` volume mount on each of these
  services once cutover completes.
- `meepo-push-service/templates/email/` retains its files but the
  volume is no longer exposed externally; external mounts are
  deprecated.
