package main

import (
	"flag"
	"net/http" // 只用于静态文件服务
	"os"

	"helloworld/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, servers ...transport.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers...),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	// 启动静态文件服务（标准库 http）
	go func() {
		http.ListenAndServe(":8081", http.FileServer(http.Dir("web")))
	}()

	if err := c.Load(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		log.Fatalf("Failed to scan config: %v", err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}
}
