package events

import "github.com/infigaming-com/meepo-api/common"

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
	UserId          int64                   `json:"userId"`
	OperatorContext *common.OperatorContext `json:"operatorContext"`
	Retries         int32                   `json:"retries"`
	MaxRetries      int32                   `json:"maxRetries"`
	Timestamp       int64                   `json:"timestamp"`
}

type UserPasswordResetEvent struct {
	UserId          int64                   `json:"userId"`
	OperatorContext *common.OperatorContext `json:"operatorContext"`
	Timestamp       int64                   `json:"timestamp"`
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

type GameTransactionEvent struct {
	OperatorId                int64       `json:"operator_id,omitempty"`
	OperatorName              string      `json:"operator_name,omitempty"`
	ProviderId                string      `json:"provider_id,omitempty"`
	ProviderName              string      `json:"provider_name,omitempty"`
	UserId                    int64       `json:"user_id,omitempty"`
	GameId                    string      `json:"game_id,omitempty"`
	GameName                  string      `json:"game_name,omitempty"`
	GameAction                string      `json:"game_action,omitempty"` // "rollback"
	RoundId                   int64       `json:"round_id,omitempty"`
	BetId                     int64       `json:"bet_id,omitempty"`
	TransactionId             int64       `json:"transaction_id,omitempty"`
	FeeGroup                  string      `json:"fee_group,omitempty"`
	Currency                  string      `json:"currency,omitempty"`
	SettlementCurrency        string      `json:"settlement_currency,omitempty"`
	AmountCurrency            string      `json:"amount_currency,omitempty"`
	AmountSettlementCurrency  string      `json:"amount_settlement_currency,omitempty"`
	AmountUsd                 string      `json:"amount_usd,omitempty"`
	AmountReportingCurrency   string      `json:"amount_reporting_currency,omitempty"`
	Turnover                  string      `json:"turnover,omitempty"`
	TurnoverUsd               string      `json:"turnover_usd,omitempty"`
	TurnoverReportingCurrency string      `json:"turnover_reporting_currency,omitempty"`
	CreatedAt                 int64       `json:"created_at,omitempty"`
	SystemOperatorId          int64       `json:"system_operator_id,omitempty"`
	RetailerOperatorId        int64       `json:"retailer_operator_id,omitempty"`
	CompanyOperatorId         int64       `json:"company_operator_id,omitempty"`
	EventInfos                []EventInfo `json:"event_infos,omitempty"`
	Type                      string      `json:"type,omitempty"`
}

type GameBetEvent struct {
	BetID                                int64                  `json:"bet_id"`
	ProviderBetID                        string                 `json:"provider_bet_id"`
	RoundID                              int64                  `json:"round_id"`
	ProviderRoundID                      string                 `json:"provider_round_id"`
	OperatorID                           int64                  `json:"operator_id"`
	OperatorName                         string                 `json:"operator_name"`
	ProviderID                           string                 `json:"provider_id"`
	ProviderName                         string                 `json:"provider_name"`
	UserID                               int64                  `json:"user_id"`
	GameID                               string                 `json:"game_id"`
	GameName                             string                 `json:"game_name"`
	GameCategory                         string                 `json:"game_category"`
	Currency                             string                 `json:"currency"`
	SettlementCurrency                   string                 `json:"settlement_currency"`
	ReportingCurrency                    string                 `json:"reporting_currency"`
	Status                               string                 `json:"status"`
	BetAmount                            string                 `json:"bet_amount"`
	BetAmountUSD                         string                 `json:"bet_amount_usd"`
	BetAmountReportingCurrency           string                 `json:"bet_amount_reporting_currency"`
	SettlementBetAmount                  string                 `json:"settlement_bet_amount"`
	SettlementBetAmountUSD               string                 `json:"settlement_bet_amount_usd"`
	SettlementBetAmountReportingCurrency string                 `json:"settlement_bet_amount_reporting_currency"`
	WinAmount                            string                 `json:"win_amount"`
	WinAmountUSD                         string                 `json:"win_amount_usd"`
	WinAmountReportingCurrency           string                 `json:"win_amount_reporting_currency"`
	SettlementWinAmount                  string                 `json:"settlement_win_amount"`
	SettlementWinAmountUSD               string                 `json:"settlement_win_amount_usd"`
	SettlementWinAmountReportingCurrency string                 `json:"settlement_win_amount_reporting_currency"`
	Turnover                             string                 `json:"turnover"`
	TurnoverUSD                          string                 `json:"turnover_usd"`
	TurnoverReportingCurrency            string                 `json:"turnover_reporting_currency"`
	WinCount                             int32                  `json:"win_count"`
	WagerType                            string                 `json:"wager_type"`
	Selections                           map[string]interface{} `json:"selections"`
	SettleTime                           int64                  `json:"settle_time"`
	CreatedAt                            int64                  `json:"created_at"`
	UpdatedAt                            int64                  `json:"updated_at"`
	SystemOperatorID                     int64                  `json:"system_operator_id"`
	RetailerOperatorID                   int64                  `json:"retailer_operator_id"`
	CompanyOperatorID                    int64                  `json:"company_operator_id"`
}

type GameBetEventBoardcastItem struct {
	BetID        int64   `json:"bet_id"`
	OperatorID   int64   `json:"operator_id"`
	OperatorName string  `json:"operator_name"`
	GameName     string  `json:"game_name"`
	GameCategory string  `json:"game_category"`
	Currency     string  `json:"currency"`
	Status       string  `json:"status"`
	BetAmount    string  `json:"bet_amount"`
	WinAmount    string  `json:"win_amount"`
	Turnover     string  `json:"turnover"`
	WinCount     int32   `json:"win_count"`
	WagerType    string  `json:"wager_type"`
	SettleTime   int64   `json:"settle_time"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
	BetMutiple   float32 `json:"bet_mutiple"`
	UserName     string  `json:"user_name"`
}

type GameBetEventBoardcast struct {
	OperatorID    int64                       `json:"operator_id"`
	AllBetEvents  []GameBetEventBoardcastItem `json:"all_bet_events"`
	HighWinEvents []GameBetEventBoardcastItem `json:"high_win_events"`
}

type GameTransactionType string

const (
	GameTransactionTypeRollback   GameTransactionType = "rollback"
	GameTransactionTypeAdjust     GameTransactionType = "adjust"
	GameTransactionTypeCancelled  GameTransactionType = "cancelled"
	GameTransactionTypeNoMoreBets GameTransactionType = "no_more_bets"
)

const GameEventTopic = "game.events"

const GameBetTopic = "game.bet"

const GameBetBoardcastTopic = "game.bet.boardcast"
