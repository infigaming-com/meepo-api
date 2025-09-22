package events

type IntegrityEventReportFileInfo struct {
	FilePath string `json:"filePath"`
	Hash     string `json:"hash"`
}

type IntegrityReportEvent struct {
	LabelApp     string                         `json:"labelApp"`
	PodName      string                         `json:"podName"`
	PodNamespace string                         `json:"podNamespace"`
	FileInfos    []IntegrityEventReportFileInfo `json:"fileInfos"`
	CreatedAt    int64                          `json:"createdAt"`
}

type UserPasswordRetriesExceededEvent struct {
	UserId     int64 `json:"userId"`
	OperatorId int64 `json:"operatorId"`
	Retries    int32 `json:"retries"`
	MaxRetries int32 `json:"maxRetries"`
	Timestamp  int64 `json:"timestamp"`
}

type UserPasswordResetEvent struct {
	UserId     int64 `json:"userId"`
	OperatorId int64 `json:"operatorId"`
	Timestamp  int64 `json:"timestamp"`
}

type EventInfo struct {
	EventId                   string `json:"event_id,omitempty"`
	EventName                 string `json:"event_name,omitempty"`
	Venue                     string `json:"venue,omitempty"`
	Tournament                string `json:"tournament,omitempty"`
	OddsInUserStyle           string `json:"odds_in_user_style,omitempty"`
	BetTypeName               string `json:"bet_type_name,omitempty"`
	WinMultiplier             string `json:"win_multiplier,omitempty"`
	BetDetails                string `json:"bet_details,omitempty"`
	EventDescription          string `json:"event_description,omitempty"`
	LeagueName                string `json:"league_name,omitempty"`
	HomeTeam                  string `json:"home_team,omitempty"`
	AwayTeam                  string `json:"away_team,omitempty"`
	CurrentResult             string `json:"current_result,omitempty"`
	EventDateUtc              int64  `json:"event_date_utc,omitempty"`
	SettlementStatus          string `json:"settlement_status,omitempty"`
	SettlementDescription     string `json:"settlement_description,omitempty"`
	BetSettledDate            int64  `json:"bet_settled_date,omitempty"`
	TriggeredResult           string `json:"triggered_result,omitempty"`
	CardedTime                int64  `json:"carded_time,omitempty"`
	FirstBetTime              int64  `json:"first_bet_time,omitempty"`
	LastBetTime               int64  `json:"last_bet_time,omitempty"`
	NoMoreBetsTime            int64  `json:"no_more_bets_time,omitempty"`
	ResultProcessingTime      int64  `json:"result_processing_time,omitempty"`
	RevisedOfficialCardedTime int64  `json:"revised_official_carded_time,omitempty"`
	Status                    string `json:"status,omitempty"`
}

type GameRollbackEvent struct {
	OperatorId                int64     `json:"operator_id,omitempty"`
	OperatorName              string    `json:"operator_name,omitempty"`
	ProviderId                string    `json:"provider_id,omitempty"`
	ProviderName              string    `json:"provider_name,omitempty"`
	UserId                    int64     `json:"user_id,omitempty"`
	GameId                    string    `json:"game_id,omitempty"`
	GameName                  string    `json:"game_name,omitempty"`
	GameAction                string    `json:"game_action,omitempty"` // "rollback"
	RoundId                   int64     `json:"round_id,omitempty"`
	BetId                     int64     `json:"bet_id,omitempty"`
	TransactionId             int64     `json:"transaction_id,omitempty"`
	FeeGroup                  string    `json:"fee_group,omitempty"`
	Currency                  string    `json:"currency,omitempty"`
	SettlementCurrency        string    `json:"settlement_currency,omitempty"`
	AmountCurrency            string    `json:"amount_currency,omitempty"`
	AmountSettlementCurrency  string    `json:"amount_settlement_currency,omitempty"`
	AmountUsd                 string    `json:"amount_usd,omitempty"`
	AmountReportingCurrency   string    `json:"amount_reporting_currency,omitempty"`
	Turnover                  string    `json:"turnover,omitempty"`
	TurnoverUsd               string    `json:"turnover_usd,omitempty"`
	TurnoverReportingCurrency string    `json:"turnover_reporting_currency,omitempty"`
	CreatedAt                 int64     `json:"created_at,omitempty"`
	SystemOperatorId          int64     `json:"system_operator_id,omitempty"`
	RetailerOperatorId        int64     `json:"retailer_operator_id,omitempty"`
	CompanyOperatorId         int64     `json:"company_operator_id,omitempty"`
	EventInfo                 EventInfo `json:"event_info,omitempty"`
}

type GameAdjustmentEvent struct {
	OperatorId                int64     `json:"operator_id,omitempty"`
	OperatorName              string    `json:"operator_name,omitempty"`
	ProviderId                string    `json:"provider_id,omitempty"`
	ProviderName              string    `json:"provider_name,omitempty"`
	UserId                    int64     `json:"user_id,omitempty"`
	GameId                    string    `json:"game_id,omitempty"`
	GameName                  string    `json:"game_name,omitempty"`
	GameAction                string    `json:"game_action,omitempty"` // "rollback"
	RoundId                   int64     `json:"round_id,omitempty"`
	BetId                     int64     `json:"bet_id,omitempty"`
	TransactionId             int64     `json:"transaction_id,omitempty"`
	FeeGroup                  string    `json:"fee_group,omitempty"`
	Currency                  string    `json:"currency,omitempty"`
	SettlementCurrency        string    `json:"settlement_currency,omitempty"`
	AmountCurrency            string    `json:"amount_currency,omitempty"`
	AmountSettlementCurrency  string    `json:"amount_settlement_currency,omitempty"`
	AmountUsd                 string    `json:"amount_usd,omitempty"`
	AmountReportingCurrency   string    `json:"amount_reporting_currency,omitempty"`
	Turnover                  string    `json:"turnover,omitempty"`
	TurnoverUsd               string    `json:"turnover_usd,omitempty"`
	TurnoverReportingCurrency string    `json:"turnover_reporting_currency,omitempty"`
	CreatedAt                 int64     `json:"created_at,omitempty"`
	SystemOperatorId          int64     `json:"system_operator_id,omitempty"`
	RetailerOperatorId        int64     `json:"retailer_operator_id,omitempty"`
	CompanyOperatorId         int64     `json:"company_operator_id,omitempty"`
	EventInfo                 EventInfo `json:"event_info,omitempty"`
}
