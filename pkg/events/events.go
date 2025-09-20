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


type GameRollbackEvent struct {
	OperatorID                   int64  `json:"operator_id"`
	OperatorName                 string `json:"operator_name"`
	ProviderID                   string `json:"provider_id"`
	ProviderName                 string `json:"provider_name"`
	UserID                       int64  `json:"user_id"`
	GameID                       string `json:"game_id"`
	GameName                     string `json:"game_name"`
	RoundID                      int64  `json:"round_id"`
	BetID                        int64  `json:"bet_id"`
	TransactionID                int64  `json:"transaction_id"`
	Currency                     string `json:"currency"`
	SettlementCurrency           string `json:"settlement_currency"`
	
	GameAction                   string `json:"game_action"`           // "rollback"
	OriginalTransactionID        int64  `json:"original_transaction_id"` // 被回滚的原始交易ID
	OriginalProviderTransactionID string `json:"original_provider_transaction_id"` // 原始provider交易ID
	RollbackType                 string `json:"rollback_type"`         // "game_bet_rollback" 或 "game_win_rollback"
	
	// 金额信息
	AmountCurrency               string `json:"amount_currency"`
	AmountSettlementCurrency     string `json:"amount_settlement_currency"`
	AmountUSD                    string `json:"amount_usd"`
	Turnover                     string `json:"turnover"`
	TurnoverUSD                  string `json:"turnover_usd"`
	
	// 时间信息
	OriginalBetTime              int64  `json:"original_bet_time"`     // 原始投注时间
	OriginalSettleTime           int64  `json:"original_settle_time"`  // 原始结算时间
	RollbackTime                 int64  `json:"rollback_time"`         // 回滚时间
	ProcessedAt                  int64  `json:"processed_at"`          // 处理时间
	
	// 操作员信息
	SystemOperatorID             int64  `json:"system_operator_id"`
	RetailerOperatorID           int64  `json:"retailer_operator_id"`
	CompanyOperatorID            int64  `json:"company_operator_id"`
	
	// 回滚原因和上下文
	RollbackReason               string `json:"rollback_reason,omitempty"` // 回滚原因
	SmResult                     string `json:"sm_result,omitempty"`       // SM结果
	Finished                     bool   `json:"finished"`                  // 是否完成
}