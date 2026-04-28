// Package notification holds template-name constants for the Meepo
// business-notification pipeline. See README.md in this directory for the
// pipeline contract.
//
// Template names are plain strings (not enums) so adding a new template
// does not force every business service to rebuild against a new
// meepo-api version. Each constant maps 1:1 to:
//
//   - meepo-push-service/templates/email/<name>_template.{txt,html}
//   - an entry in push-service's internal template registry
//
// Adding a new template requires three coordinated PRs:
//
//  1. push-service: drop the template files, register the name + subject
//     format string in the internal registry.
//  2. meepo-api: add the corresponding constant here. Cut a tag.
//  3. business service: bump meepo-api dep, build the params map,
//     publish notification.service.v1.BusinessAlertEvent.
package notification

const (
	// TemplateSubAccountBalanceAlert fires when an operator's sub-account
	// cash drops below the per-operator alert threshold configured in
	// wallet.operator_prediction_settings.balance_alert_threshold.
	// Required params: operatorName, productType, currency, balance,
	// threshold, eventTimeUTC.
	TemplateSubAccountBalanceAlert = "sub_account_balance_alert"
)
