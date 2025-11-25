package events

import (
	"encoding/json"

	"github.com/infigaming-com/meepo-api/common"
)

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

const (
	WEBSOCKET_MESSAGE_CLIENT   = 1
	WEBSOCKET_MESSAGE_USER     = 2
	WEBSOCKET_MESSAGE_OPERATOR = 3
)

type WebsocketMessageEvent struct {
	ID      int64  `json:"ID"`
	Channel string `json:"Channel"`
	Type    int    `json:"Type"`
	Message []byte `json:"Message"`
}

type WebsocketMessageMessage[T any] struct {
	M string `json:"m"`
	V string `json:"v"`
	T string `json:"t"`
	P T      `json:"p"`
}

func BuildWebsocketEvent[T any](id int64, command string, mssageType int, data T) (*WebsocketMessageEvent, error) {
	message := WebsocketMessageMessage[T]{
		M: command,
		V: "1.0.0",
		T: "",
		P: data,
	}

	if jsonMessage, err := json.Marshal(&message); err != nil {
		return nil, err
	} else {
		return &WebsocketMessageEvent{
			ID:      id,
			Channel: "",
			Type:    mssageType,
			Message: jsonMessage,
		}, nil
	}
}

type ClientOnlineStatusEvent struct {
	ClientId   int64 `json:"clientId"`
	OperatorId int64 `json:"operatorId"`
}

type UserOnlineStatusEvent struct {
	UserId     int64 `json:"userId"`
	OperatorId int64 `json:"operatorId"`
	Online     bool  `json:"online"`
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
	OperatorId                   int64       `json:"operator_id,omitempty"`
	OperatorName                 string      `json:"operator_name,omitempty"`
	ProviderId                   string      `json:"provider_id,omitempty"`
	ProviderName                 string      `json:"provider_name,omitempty"`
	UserId                       int64       `json:"user_id,omitempty"`
	GameId                       string      `json:"game_id,omitempty"`
	GameName                     string      `json:"game_name,omitempty"`
	GameAction                   string      `json:"game_action,omitempty"` // "rollback"
	RoundId                      int64       `json:"round_id,omitempty"`
	BetId                        int64       `json:"bet_id,omitempty"`
	TransactionId                int64       `json:"transaction_id,omitempty"`
	FeeGroup                     string      `json:"fee_group,omitempty"`
	Currency                     string      `json:"currency,omitempty"`
	SettlementCurrency           string      `json:"settlement_currency,omitempty"`
	ExchangeRate                 string      `json:"exchange_rate,omitempty"`
	AmountCurrency               string      `json:"amount_currency,omitempty"`
	AmountSettlementCurrency     string      `json:"amount_settlement_currency,omitempty"`
	AmountUsd                    string      `json:"amount_usd,omitempty"`
	AmountReportingCurrency      string      `json:"amount_reporting_currency,omitempty"`
	CashAmountCurrency           string      `json:"cash_amount_currency,omitempty"`
	CashAmountSettlementCurrency string      `json:"cash_amount_settlement_currency,omitempty"`
	CashAmountUsd                string      `json:"cash_amount_usd,omitempty"`
	CashAmountReportingCurrency  string      `json:"cash_amount_reporting_currency,omitempty"`
	Turnover                     string      `json:"turnover,omitempty"`
	TurnoverUsd                  string      `json:"turnover_usd,omitempty"`
	TurnoverReportingCurrency    string      `json:"turnover_reporting_currency,omitempty"`
	CreatedAt                    int64       `json:"created_at,omitempty"`
	SystemOperatorId             int64       `json:"system_operator_id,omitempty"`
	RetailerOperatorId           int64       `json:"retailer_operator_id,omitempty"`
	CompanyOperatorId            int64       `json:"company_operator_id,omitempty"`
	UserTotalBetCount            int64       `json:"user_total_bet_count,omitempty"`
	EventInfos                   []EventInfo `json:"event_infos,omitempty"`
	Type                         string      `json:"type,omitempty"`
	RollbackType                 string      `json:"rollback_type,omitempty"`
	UpdateBalance                bool        `json:"update_balance,omitempty"`
	RollbackOriginalAction       string      `json:"rollback_original_action,omitempty"`
}

type GameBetEvent struct {
	BetID                                int64          `json:"bet_id"`
	ProviderBetID                        string         `json:"provider_bet_id"`
	RoundID                              int64          `json:"round_id"`
	ProviderRoundID                      string         `json:"provider_round_id"`
	OperatorID                           int64          `json:"operator_id"`
	OperatorName                         string         `json:"operator_name"`
	ProviderID                           string         `json:"provider_id"`
	ProviderName                         string         `json:"provider_name"`
	UserID                               int64          `json:"user_id"`
	GameID                               string         `json:"game_id"`
	GameName                             string         `json:"game_name"`
	GameCategory                         string         `json:"game_category"`
	Currency                             string         `json:"currency"`
	SettlementCurrency                   string         `json:"settlement_currency"`
	ReportingCurrency                    string         `json:"reporting_currency"`
	Status                               string         `json:"status"`
	BetAmount                            string         `json:"bet_amount"`
	BetAmountUSD                         string         `json:"bet_amount_usd"`
	BetAmountReportingCurrency           string         `json:"bet_amount_reporting_currency"`
	SettlementBetAmount                  string         `json:"settlement_bet_amount"`
	SettlementBetAmountUSD               string         `json:"settlement_bet_amount_usd"`
	SettlementBetAmountReportingCurrency string         `json:"settlement_bet_amount_reporting_currency"`
	WinAmount                            string         `json:"win_amount"`
	WinAmountUSD                         string         `json:"win_amount_usd"`
	WinAmountReportingCurrency           string         `json:"win_amount_reporting_currency"`
	SettlementWinAmount                  string         `json:"settlement_win_amount"`
	SettlementWinAmountUSD               string         `json:"settlement_win_amount_usd"`
	SettlementWinAmountReportingCurrency string         `json:"settlement_win_amount_reporting_currency"`
	Turnover                             string         `json:"turnover"`
	TurnoverUSD                          string         `json:"turnover_usd"`
	TurnoverReportingCurrency            string         `json:"turnover_reporting_currency"`
	WinCount                             int32          `json:"win_count"`
	WagerType                            string         `json:"wager_type"`
	Selections                           map[string]any `json:"selections"`
	SettleTime                           int64          `json:"settle_time"`
	CreatedAt                            int64          `json:"created_at"`
	UpdatedAt                            int64          `json:"updated_at"`
	SystemOperatorID                     int64          `json:"system_operator_id"`
	RetailerOperatorID                   int64          `json:"retailer_operator_id"`
	CompanyOperatorID                    int64          `json:"company_operator_id"`
	GameTransactionIDs                   []int64        `json:"game_transaction_ids"`
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
	BetOdds      float32 `json:"bet_odds"`
	UserID       int64   `json:"user_id"`
	UserName     string  `json:"user_name"`
}

type WebsocketOperatorBoardcastEvent struct {
	OperatorID int64  `json:"operator_id"`
	Command    string `json:"command"`
	Data       any    `json:"data"`
}

type WebsocketOperatorUserEvent struct {
	UserID  int64  `json:"user_id"`
	Command string `json:"command"`
	Data    any    `json:"data"`
}

type WebsocketOperatorClientEvent struct {
	ClientID int64  `json:"client_id"`
	Command  string `json:"command"`
	Data     any    `json:"data"`
}

type SessionCreatedEvent struct {
	SessionID          int64  `json:"session_id"`
	UserID             string `json:"user_id"`
	GameID             string `json:"game_id"`
	GameName           string `json:"game_name"`
	ProviderID         string `json:"provider_id"`
	ProviderName       string `json:"provider_name"`
	Currency           string `json:"currency"`
	SettlementCurrency string `json:"settlement_currency"`
	OperatorID         int64  `json:"operator_id"`
	SystemOperatorID   int64  `json:"system_operator_id"`
	RetailerOperatorID int64  `json:"retailer_operator_id"`
	CompanyOperatorID  int64  `json:"company_operator_id"`
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
const WEBSOCKET_PUSH_MESSAGE = "websocket.push.message"

const GameBetUserTopic = "game.bet.user"

const GameBetClientTopic = "game.bet.client"

const CreateSessionTopic = "game.create.session"

const GameBetStatusSettled = "settled"

// ------------------------------------------------------------
// Affiliate events
// ------------------------------------------------------------

const AffiliateBillSettlementTopic = "affiliate.bill.settlement"

type AffiliateBillSettlementEvent struct {
	UserID                  int64                   `json:"user_id"`
	ReportingCurrency       string                  `json:"reporting_currency"`
	AmountReportingCurrency string                  `json:"amount_reporting_currency"`
	AmountUSD               string                  `json:"amount_usd"`
	BillID                  int64                   `json:"bill_id"`
	OperatorContext         *common.OperatorContext `json:"operator_context,omitempty"`
	CreatedAt               int64                   `json:"created_at"`
}

const AffiliateUserActionEventTopic = "affiliate.user.action"

type AffiliateUserActionEvent struct {
	UserID          int64                   `json:"user_id"`
	EventID         int64                   `json:"event_id"`
	EventType       string                  `json:"event_type"`
	ChannelEvent    *string                 `json:"channel_event,omitempty"`
	CustomName      *string                 `json:"custom_name,omitempty"`
	OperatorContext *common.OperatorContext `json:"operator_context,omitempty"`
	CreatedAt       int64                   `json:"created_at"`
}

type OperatorSuspendEvent struct {
	OperatorContext *common.OperatorContext `json:"operator_context,omitempty"`
}

const OperatorSuspendTopic = "operator.suspend"