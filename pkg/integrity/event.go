package integrity

type FileInfo struct {
	FilePath string `json:"filePath"`
	Hash     string `json:"hash"`
}

type IntegrityEvent struct {
	LabelApp     string     `json:"labelApp"`
	PodName      string     `json:"podName"`
	PodNamespace string     `json:"podNamespace"`
	FileInfos    []FileInfo `json:"fileInfos"`
	CreatedAt    int64      `json:"createdAt"`
}
