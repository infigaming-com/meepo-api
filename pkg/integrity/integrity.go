package integrity

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	chash "github.com/infigaming-com/go-common/hash"
	pubsub "github.com/infigaming-com/meepo-api/pubsub/service/v1"
)

type IntegrityService struct {
	lg           *log.Helper
	wg           sync.WaitGroup
	podName      string
	podNamespace string
	labelApp     string
	filePaths    []string
	pubsubClient pubsub.PubsubClient
	reportTopic  string
}

func NewIntegrityService(lg log.Logger, filePaths []string, pubsubClient pubsub.PubsubClient, reportTopic string) (*IntegrityService, func()) {
	log := log.NewHelper(lg)

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

	ticker := time.NewTicker(5 * time.Minute)
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
	for _, filePath := range is.filePaths {
		is.checkAndReportIntegrityForFile(ctx, filePath)
	}
}

func (is *IntegrityService) checkAndReportIntegrityForFile(ctx context.Context, filePath string) {
	defer func() {
		if r := recover(); r != nil {
			is.lg.Errorf("integrity service panic processing file %s: %v", filePath, r)
		}
	}()

	file, err := os.Open(filePath)
	if err != nil {
		is.lg.Errorf("Failed to open file %s: %v", filePath, err)
		return
	}
	defer file.Close()

	hashResult, err := chash.HashWithReader(file, chash.SHA256)
	if err != nil {
		is.lg.Errorf("Failed to hash file %s: %v", filePath, err)
		return
	}

	is.lg.Debugf("Hash for file %s: %s", filePath, hashResult)

	event := IntegrityEvent{
		PodName:      is.podName,
		PodNamespace: is.podNamespace,
		LabelApp:     is.labelApp,
		File:         filePath,
		Hash:         hashResult,
		CreatedAt:    time.Now().UnixMilli(),
	}
	eventData, err := json.Marshal(event)
	if err != nil {
		is.lg.Errorf("Failed to marshal integrity event for file %s: %v", filePath, err)
		return
	}

	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 30*time.Second)
	defer timeoutCancel()

	_, err = is.pubsubClient.Pub(timeoutCtx, &pubsub.PubRequest{
		Topic:   is.reportTopic,
		Message: eventData,
	})
	if err != nil {
		is.lg.Errorf("Failed to publish integrity event for file %s: %v", filePath, err)
		return
	}
}
