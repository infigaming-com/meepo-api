package main

import (
	"flag"
	"os"

	"github.com/infigaming-com/meepo-api/app/reward/service/internal/conf"
	"github.com/infigaming-com/meepo-api/pkg/zlogger"

	zaplog "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "meepo-api.reward.service"
	// Version is the version of the compiled software.
	Version string
	// Build is the build date of the compiled software.
	Build string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func initLogger(c *conf.Logger) (log.Logger, func()) {
	zlogger, close, err := zlogger.NewZLogger(c.Level)
	if err != nil {
		panic(err)
	}

	return log.With(zaplog.NewLogger(zlogger),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"service.build", Build,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	), close
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	logger, close := initLogger(bc.Logger)
	defer close()

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
