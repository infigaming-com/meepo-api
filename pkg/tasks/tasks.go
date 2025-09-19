package tasks

const (
	REPORT_EXPORT_PENDING    = "pending"
	REPORT_EXPORT_FAILED     = "failed"
	REPORT_EXPORT_GENERATING = "generating"
	REPORT_EXPORT_COMPLETED  = "completed"
)

type ReportExportPayload struct {
}

type ReportExportResult struct {
}
