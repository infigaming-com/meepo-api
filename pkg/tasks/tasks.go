package tasks

const (
	REPORT_EXPORT_STATUS_PENDING    = "pending"
	REPORT_EXPORT_STATUS_FAILED     = "failed"
	REPORT_EXPORT_STATUS_GENERATING = "generating"
	REPORT_EXPORT_STATUS_COMPLETED  = "completed"
)

const (
	REPORT_EXPORT_TYPE_CUSTOMER_RECORDS            = "customer_records"
	REPORT_EXPORT_TYPE_FICA_THRESHOLD_TRANSACTIONS = "fica_threshold_transactions"
	REPORT_EXPORT_TYPE_UNPAID_BET                  = "unpaid_bets"
)

type ReportExportPayload struct {
}

type ReportExportResult struct {
}
