package tasks

import "fmt"

const (
	REPORT_EXPORT_STATUS_PENDING    = "pending"
	REPORT_EXPORT_STATUS_FAILED     = "failed"
	REPORT_EXPORT_STATUS_GENERATING = "generating"
	REPORT_EXPORT_STATUS_COMPLETED  = "completed"
)

const (
	REPORT_EXPORT_TYPE_CUSTOMER_RECORDS            = "customer_records"
	REPORT_EXPORT_TYPE_FICA_THRESHOLD_TRANSACTIONS = "fica_threshold_transactions"
	REPORT_EXPORT_TYPE_UNPAID_BETS                 = "unpaid_bets"
)

type ReportExportPayload struct {
}

type ReportExportResult struct {
}

// GetMimeTypeFromExtension returns the appropriate MIME type based on file extension
func GetMimeTypeFromExtension(extension string) string {
	switch extension {
	case "csv":
		return "text/csv"
	case "xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case "pdf":
		return "application/pdf"
	default:
		// Default to CSV for unknown extensions
		return "text/csv"
	}
}

// FormatFileSize formats file size in bytes to a human-readable string
func FormatFileSize(bytes int) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d bytes", bytes)
	}
}
