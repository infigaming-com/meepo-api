package integrity

type IntegrityEvent struct {
	PodName      string
	PodNamespace string
	LabelApp     string
	File         string
	Hash         string
	CreatedAt    int64
}
