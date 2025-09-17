package integrity

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	chash "github.com/infigaming-com/go-common/hash"
	"github.com/infigaming-com/meepo-api/pkg/events"
	pubsub "github.com/infigaming-com/meepo-api/pubsub/service/v1"
)

type IntegrityService struct {
	lg           *log.Helper
	interval     time.Duration
	podName      string
	podNamespace string
	labelApp     string
	filePaths    []string
	pubsubClient pubsub.PubsubClient
	reportTopic  string
	wg           sync.WaitGroup
}

func NewIntegrityService(lg log.Logger, interval time.Duration, filePaths []string, pubsubClient pubsub.PubsubClient, reportTopic string) (*IntegrityService, func()) {
	log := log.NewHelper(lg)

	if interval == 0 {
		interval = 5 * time.Minute
	}

	if len(filePaths) == 0 {
		log.Fatal("filePaths not provided")
	}

	if pubsubClient == nil {
		log.Fatal("pubsubClient not provided")
	}

	if reportTopic == "" {
		log.Fatal("reportTopic not provided")
	}

	podName := os.Getenv("POD_NAME")
	if podName == "" {
		log.Warnf("POD_NAME environment variable not set")
	}

	podNamespace := os.Getenv("POD_NAMESPACE")
	if podNamespace == "" {
		log.Warnf("POD_NAMESPACE environment variable not set")
	}

	labelApp := os.Getenv("LABEL_APP")
	if labelApp == "" {
		log.Warnf("LABEL_APP environment variable not set")
	}

	log.Debugf("integrity service starting...")
	log.Debugf("podName: %s, podNamespace: %s, labelApp: %s, filePaths: %v", podName, podNamespace, labelApp, filePaths)

	cancelCtx, cancel := context.WithCancel(context.Background())

	is := &IntegrityService{
		lg:           log,
		interval:     interval,
		podName:      podName,
		podNamespace: podNamespace,
		labelApp:     labelApp,
		filePaths:    filePaths,
		pubsubClient: pubsubClient,
		reportTopic:  reportTopic,
	}

	is.wg.Add(1)
	go is.Start(cancelCtx)

	return is, func() {
		is.lg.Debugf("integrity service stopping...")
		cancel()
		is.wg.Wait()
	}
}

func (is *IntegrityService) Start(ctx context.Context) {
	defer is.wg.Done()

	defer func() {
		if r := recover(); r != nil {
			is.lg.Errorf("integrity service panic: %v", r)
		}
	}()

	ticker := time.NewTicker(is.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			is.checkAndReportIntegrity(ctx)
		}
	}
}

func (is *IntegrityService) checkAndReportIntegrity(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			is.lg.Errorf("integrity service panic: %v", r)
		}
	}()

	fileInfos, err := getFileInfos(is.filePaths)
	if err != nil {
		is.lg.Errorf("Failed to get file infos: %v", err)
		return
	}
	is.lg.Debugf("FileInfos: %+v", fileInfos)

	event := events.IntegrityReportEvent{
		LabelApp:     is.labelApp,
		PodName:      is.podName,
		PodNamespace: is.podNamespace,
		FileInfos:    fileInfos,
		CreatedAt:    time.Now().UnixMilli(),
	}
	eventData, err := json.Marshal(event)
	if err != nil {
		is.lg.Errorf("Failed to marshal integrity event: %v", err)
		return
	}

	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, min(is.interval, 30*time.Second))
	defer timeoutCancel()

	_, err = is.pubsubClient.Pub(timeoutCtx, &pubsub.PubRequest{
		Topic:   is.reportTopic,
		Message: eventData,
	})
	if err != nil {
		is.lg.Errorf("Failed to publish integrity event: %v", err)
		return
	}
	is.lg.Debugf("Published integrity event: %+v", event)
}

func getFileInfos(filePaths []string) ([]events.IntegrityEventReportFileInfo, error) {
	fileInfos := make([]events.IntegrityEventReportFileInfo, 0, len(filePaths))
	for _, filePath := range filePaths {
		hash, err := getFileHash(filePath)
		if err != nil {
			return nil, err
		}
		fileInfos = append(fileInfos, events.IntegrityEventReportFileInfo{FilePath: filePath, Hash: hash})
	}
	return fileInfos, nil
}

func getFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return chash.HashWithReader(file, chash.SHA256)
}
