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

type UserPasswordRetiresExceededEvent struct {
	UserId          int64                  `json:"userId"`
	OperatorContext common.OperatorContext `json:"operatorContext"`
	Retries         int32                  `json:"retries"`
	MaxRetries      int32                  `json:"maxRetries"`
	Timestamp       int64                  `json:"timestamp"`
}

type UserPasswordResetEvent struct {
	UserId          int64                  `json:"userId"`
	OperatorContext common.OperatorContext `json:"operatorContext"`
}
