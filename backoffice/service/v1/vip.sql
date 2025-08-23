-- VIP system schema (PostgreSQL)
-- This file defines enums, settings, levels, benefits, members, ledgers, and settlement state

CREATE TYPE vip_display_rule AS ENUM ('cumulative','current');
-- Reward kinds aligned with UI blocks
-- upgrade_base / upgrade_incremental map to the two upgrade sections
-- weekly_reward / monthly_reward map to weekly/monthly blocks
CREATE TYPE vip_reward_kind AS ENUM (
  'upgrade_base',
  'upgrade_incremental',
  'rakeback_instant',
  'rakeback_daily',
  'rakeback_weekly',
  'rakeback_monthly'
);
CREATE TYPE vip_payout_type AS ENUM ('cash','bonus');
CREATE TYPE vip_reset_policy AS ENUM ('none','reset_on_upgrade','reset_next_week','reset_next_month');
CREATE TYPE vip_reward_status AS ENUM ('pending','issued','expired','revoked');

-- 1) General Settings / Reward Type / XP Config / Settlement
CREATE TABLE vip_setting (
  id                BIGINT PRIMARY KEY,
  -- flattened scope: system -> retailer -> company -> operator
  system_operator_id   BIGINT NOT NULL DEFAULT 0,
  retailer_operator_id BIGINT NOT NULL DEFAULT 0,
  company_operator_id  BIGINT NOT NULL DEFAULT 0,
  operator_id          BIGINT NOT NULL DEFAULT 0,

  base_currency     VARCHAR(16) NOT NULL DEFAULT 'USDT',
  display_rule      vip_display_rule NOT NULL DEFAULT 'cumulative',
  rewards_slider    BOOLEAN   NOT NULL DEFAULT TRUE,

  -- reward type selection
  upgrade_payout_type          vip_payout_type NOT NULL DEFAULT 'bonus',
  rakeback_instant_payout_type vip_payout_type NOT NULL DEFAULT 'cash',
  rakeback_daily_payout_type   vip_payout_type NOT NULL DEFAULT 'cash',
  weekly_reward_payout_type    vip_payout_type NOT NULL DEFAULT 'bonus',
  monthly_reward_payout_type   vip_payout_type NOT NULL DEFAULT 'bonus',

  -- XP config
  deposit_xp_rate              NUMERIC(20,8) NOT NULL DEFAULT 1,
  wagering_xp_rate             NUMERIC(20,8) NOT NULL DEFAULT 1,
  max_house_edge_limit_percent NUMERIC(10,4) NOT NULL DEFAULT 30,
  standard_xp_multiplier       NUMERIC(10,4) NOT NULL DEFAULT 1,

  -- settlement schedule
  timezone                     VARCHAR(64) NOT NULL DEFAULT 'UTC',
  daily_issue_at               TIME        NOT NULL DEFAULT '00:00',
  weekly_issue_dow             SMALLINT    NOT NULL DEFAULT 1,  -- 1=Mon..7=Sun
  weekly_issue_at              TIME        NOT NULL DEFAULT '00:00',
  monthly_issue_dom            SMALLINT    NOT NULL DEFAULT 28, -- 1..31/28
  monthly_issue_at             TIME        NOT NULL DEFAULT '00:00',

  -- reward expiry settings (moved from vip_reward_expiry table)
  upgrade_base_expiry_days         INTEGER NOT NULL DEFAULT 0,
  upgrade_base_expiry_hours        INTEGER NOT NULL DEFAULT 0,
  upgrade_base_never_expire        BOOLEAN NOT NULL DEFAULT FALSE,
  upgrade_base_reset_policy        vip_reset_policy NOT NULL DEFAULT 'none',
  
  upgrade_incremental_expiry_days  INTEGER NOT NULL DEFAULT 0,
  upgrade_incremental_expiry_hours INTEGER NOT NULL DEFAULT 0,
  upgrade_incremental_never_expire BOOLEAN NOT NULL DEFAULT FALSE,
  upgrade_incremental_reset_policy vip_reset_policy NOT NULL DEFAULT 'none',
  
  rakeback_instant_expiry_days     INTEGER NOT NULL DEFAULT 0,
  rakeback_instant_expiry_hours    INTEGER NOT NULL DEFAULT 0,
  rakeback_instant_never_expire    BOOLEAN NOT NULL DEFAULT FALSE,
  rakeback_instant_reset_policy    vip_reset_policy NOT NULL DEFAULT 'none',
  
  rakeback_daily_expiry_days       INTEGER NOT NULL DEFAULT 0,
  rakeback_daily_expiry_hours      INTEGER NOT NULL DEFAULT 0,
  rakeback_daily_never_expire      BOOLEAN NOT NULL DEFAULT FALSE,
  rakeback_daily_reset_policy      vip_reset_policy NOT NULL DEFAULT 'none',
  
  rakeback_weekly_expiry_days      INTEGER NOT NULL DEFAULT 0,
  rakeback_weekly_expiry_hours     INTEGER NOT NULL DEFAULT 0,
  rakeback_weekly_never_expire     BOOLEAN NOT NULL DEFAULT FALSE,
  rakeback_weekly_reset_policy     vip_reset_policy NOT NULL DEFAULT 'none',
  
  rakeback_monthly_expiry_days     INTEGER NOT NULL DEFAULT 0,
  rakeback_monthly_expiry_hours    INTEGER NOT NULL DEFAULT 0,
  rakeback_monthly_never_expire    BOOLEAN NOT NULL DEFAULT FALSE,
  rakeback_monthly_reset_policy    vip_reset_policy NOT NULL DEFAULT 'none',

  created_at BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,
  updated_at BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,

  UNIQUE(system_operator_id, retailer_operator_id, company_operator_id, operator_id)
);


-- 5) Members
-- TODO: 修改 vip 等级所需要的经验值后，玩家是否降级或升级, 后台是否需要相应API重算玩家 VIP 等级
CREATE TABLE vip_member (
  id                BIGINT PRIMARY KEY,
  system_operator_id   BIGINT NOT NULL DEFAULT 0,
  retailer_operator_id BIGINT NOT NULL DEFAULT 0,
  company_operator_id  BIGINT NOT NULL DEFAULT 0,
  operator_id          BIGINT NOT NULL DEFAULT 0,
  user_id           BIGINT    NOT NULL,

  current_level_id  BIGINT REFERENCES vip_level_config_template(id),
  current_xp        NUMERIC(20,4) NOT NULL DEFAULT 0,
  total_xp          NUMERIC(20,4) NOT NULL DEFAULT 0,
  last_level_up_at  BIGINT,

  last_instant_upgrade_at BIGINT,

  created_at BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,
  updated_at BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,

  UNIQUE(system_operator_id, retailer_operator_id, company_operator_id, operator_id, user_id)
);

-- 6) XP Ledger
CREATE TABLE vip_xp_ledger (
  id                BIGINT PRIMARY KEY,
  system_operator_id   BIGINT NOT NULL DEFAULT 0,
  retailer_operator_id BIGINT NOT NULL DEFAULT 0,
  company_operator_id  BIGINT NOT NULL DEFAULT 0,
  operator_id          BIGINT NOT NULL DEFAULT 0,
  user_id           BIGINT    NOT NULL,

  source_type       VARCHAR(32) NOT NULL,            -- deposit/wagering/adjust
  base_amount       NUMERIC(20,4) NOT NULL DEFAULT 0,
  currency          VARCHAR(16) NOT NULL DEFAULT 'USDT',
  xp_earned         NUMERIC(20,8) NOT NULL DEFAULT 0,
  applied_rate      NUMERIC(20,8) NOT NULL DEFAULT 0,
  house_edge_used   NUMERIC(10,4) NOT NULL DEFAULT 0,
  std_xp_multiplier NUMERIC(10,4) NOT NULL DEFAULT 1,

  event_ts          BIGINT NOT NULL,
  created_at        BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT
);
CREATE INDEX idx_vip_xp_user_ts ON vip_xp_ledger(system_operator_id, retailer_operator_id, company_operator_id, operator_id, user_id, event_ts);

-- 7) Reward Ledger
-- TODO: 是否是每笔投注都有对应的奖励记录, 奖励记录会特别大，相当于另一个投注记录表
CREATE TABLE vip_reward_ledger (
  id                BIGINT PRIMARY KEY,
  system_operator_id   BIGINT NOT NULL DEFAULT 0,
  retailer_operator_id BIGINT NOT NULL DEFAULT 0,
  company_operator_id  BIGINT NOT NULL DEFAULT 0,
  operator_id          BIGINT NOT NULL DEFAULT 0,
  user_id           BIGINT    NOT NULL,
  level_id          BIGINT REFERENCES vip_level(id),

  reward_type       vip_reward_type NOT NULL,
  payout_type       vip_payout_type NOT NULL,
  amount            NUMERIC(20,4) NOT NULL DEFAULT 0,
  currency          VARCHAR(16) NOT NULL DEFAULT 'USDT',

  claim_amount      NUMERIC(20,4) NOT NULL DEFAULT 0,
  claim_currency    VARCHAR(16) NOT NULL DEFAULT 'USDT',

  status            vip_reward_status NOT NULL DEFAULT 'pending',
  issue_at          BIGINT,
  expire_at         BIGINT,

  created_at        BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,
  updated_at        BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT
);
CREATE INDEX idx_vip_reward_user ON vip_reward_ledger(system_operator_id, retailer_operator_id, company_operator_id, operator_id, user_id, reward_kind, status);

-- 8) Settlement State (cursors for schedulers)
CREATE TABLE vip_settlement_state (
  id                BIGINT PRIMARY KEY,
  system_operator_id   BIGINT NOT NULL DEFAULT 0,
  retailer_operator_id BIGINT NOT NULL DEFAULT 0,
  company_operator_id  BIGINT NOT NULL DEFAULT 0,
  operator_id          BIGINT NOT NULL DEFAULT 0,

  daily_last_run_at   BIGINT,
  weekly_last_run_at  BIGINT,
  monthly_last_run_at BIGINT,

  created_at BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,
  updated_at BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,

  UNIQUE(system_operator_id, retailer_operator_id, company_operator_id, operator_id)
);


-- 11) VIP Config: inheritance switches per scope
-- One row per scope combination; controls whether a scope follows its parent config
-- Parent chain (default): system -> retailer -> company -> operator
CREATE TABLE vip_config (
  id                          BIGINT PRIMARY KEY,
  system_operator_id          BIGINT NOT NULL DEFAULT 0,
  retailer_operator_id        BIGINT NOT NULL DEFAULT 0,
  company_operator_id         BIGINT NOT NULL DEFAULT 0,
  operator_id                 BIGINT NOT NULL DEFAULT 0,
  enabled                     BOOLEAN NOT NULL DEFAULT TRUE,

  -- follow parent switches per config domain
  follow_parent_setting       BOOLEAN NOT NULL DEFAULT TRUE,
  follow_parent_level_tpl     BOOLEAN NOT NULL DEFAULT TRUE,

  created_at BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,
  updated_at BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,

  UNIQUE(system_operator_id, retailer_operator_id, company_operator_id, operator_id)
);

-- 9) VIP Level Config Template (bulk rule for a range of sub-levels WITHIN a level)
--   start_level/end_level here mean sub-level ranks, scoped by a specific level_id
--   Matches fields shown in "Edit VIP" → Upgrade Settings / Maintenance / Rakeback / Weekly Reward
CREATE TABLE vip_level_config_template (
  id                        BIGINT PRIMARY KEY,
  system_operator_id        BIGINT NOT NULL DEFAULT 0,
  retailer_operator_id      BIGINT NOT NULL DEFAULT 0,
  company_operator_id       BIGINT NOT NULL DEFAULT 0,
  operator_id               BIGINT NOT NULL DEFAULT 0,

  -- base info
  name                      VARCHAR(64) NOT NULL DEFAULT 'default',
  start_level               BIGINT NOT NULL,
  end_level                 BIGINT NOT NULL,

  -- Upgrade Settings
  base_level_upgrade_xp     NUMERIC(20,4) NOT NULL DEFAULT 0,
  incr_level_upgrade_xp     NUMERIC(20,4) NOT NULL DEFAULT 0,
  base_upgrade_reward_amt   NUMERIC(20,4) NOT NULL DEFAULT 0,
  incr_upgrade_reward_amt   NUMERIC(20,4) NOT NULL DEFAULT 0,
  upgrade_reward_wagering_x NUMERIC(10,2) NOT NULL DEFAULT 0, -- e.g. 30x

  -- Level Maintenance Requirements
  weekly_xp_loss            NUMERIC(20,4) NOT NULL DEFAULT 0,
  monthly_xp_loss           NUMERIC(20,4) NOT NULL DEFAULT 0,

  -- Rakeback (Instant)
  rakeback_instant_enabled  BOOLEAN NOT NULL DEFAULT TRUE,
  rakeback_instant_rate     NUMERIC(10,4) NOT NULL DEFAULT 0, -- percent
  rakeback_instant_req      NUMERIC(10,2) NOT NULL DEFAULT 0, -- wagering requirements

  -- Rakeback (Daily)
  rakeback_daily_enabled    BOOLEAN NOT NULL DEFAULT FALSE,
  rakeback_daily_rate       NUMERIC(10,4) NOT NULL DEFAULT 0,
  rakeback_daily_req        NUMERIC(10,2) NOT NULL DEFAULT 0,

  -- Weekly Reward
  weekly_reward_enabled     BOOLEAN NOT NULL DEFAULT TRUE,
  weekly_fixed_reward_amt   NUMERIC(20,4) NOT NULL DEFAULT 0,
  weekly_turnover_rate      NUMERIC(10,4) NOT NULL DEFAULT 0,  -- percent
  weekly_net_loss_amt       NUMERIC(20,4) NOT NULL DEFAULT 0,
  weekly_adjust_range_pct   NUMERIC(10,4) NOT NULL DEFAULT 0,  -- adjustable range percent
  weekly_active_days_reward SMALLINT NOT NULL DEFAULT 0,
  weekly_wagering_req       NUMERIC(10,2) NOT NULL DEFAULT 0,

  created_at BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,
  updated_at BIGINT NOT NULL DEFAULT (extract(epoch from now())*1000)::BIGINT,

  UNIQUE(system_operator_id, retailer_operator_id, company_operator_id, operator_id, level_id, start_level, end_level)
);

