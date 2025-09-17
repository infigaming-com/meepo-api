package integrity

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	chash "github.com/infigaming-com/go-common/hash"
	system "github.com/infigaming-com/meepo-api/system/service/v1"
)

type IntegrityService struct {
	lg           *log.Helper
	interval     time.Duration
	podName      string
	podNamespace string
	labelApp     string
	filePaths    []string
	systemClient system.SystemClient
	wg           sync.WaitGroup
}

func NewIntegrityService(lg log.Logger, interval time.Duration, filePaths []string, systemClient system.SystemClient) (*IntegrityService, func()) {
	log := log.NewHelper(lg)

	if interval == 0 {
		interval = 5 * time.Minute
	}

	if len(filePaths) == 0 {
		log.Fatal("filePaths not provided")
	}

	if systemClient == nil {
		log.Fatal("systemClient not provided")
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
		systemClient: systemClient,
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

	// Run immediately
	is.checkAndReportIntegrity(ctx)

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

	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, min(is.interval, 30*time.Second))
	defer timeoutCancel()

	_, err = is.systemClient.AddIntegrityReport(timeoutCtx, &system.AddIntegrityReportRequest{
		LabelApp:     is.labelApp,
		PodName:      is.podName,
		PodNamespace: is.podNamespace,
		FileInfos:    fileInfos,
	})
	if err != nil {
		is.lg.Errorf("Failed to call system service to add integrity report: %v", err)
		return
	}
	is.lg.Debugf("Added integrity report: %+v", fileInfos)
}

func getFileInfos(filePaths []string) ([]*system.AddIntegrityReportRequest_IntegrityEventReportFileInfo, error) {
	fileInfos := make([]*system.AddIntegrityReportRequest_IntegrityEventReportFileInfo, 0, len(filePaths))
	for _, filePath := range filePaths {
		hash, err := getFileHash(filePath)
		if err != nil {
			return nil, err
		}
		fileInfos = append(
			fileInfos,
			&system.AddIntegrityReportRequest_IntegrityEventReportFileInfo{
				FilePath: filePath,
				Hash:     hash,
			},
		)
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
