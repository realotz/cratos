package main

import (
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	kc "github.com/go-kratos/kube/config"
	pb "github.com/realotz/cratos/api/v1"
	"github.com/realotz/cratos/internal/conf"
	"github.com/realotz/cratos/internal/service"
	"gopkg.in/yaml.v2"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf kube config file ")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, greeter *service.IstioService) *kratos.App {
	pb.RegisterIstioServiceServer(gs, greeter)
	pb.RegisterIstioServiceHTTPServer(hs, greeter)
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func getEnv(key string, def string) string {
	if env := os.Getenv(key); env != "" {

		return env
	}
	return def
}

func main() {
	flag.Parse()
	logger := log.NewStdLogger()
	cfg := config.New(
		config.WithSource(
			kc.NewSource(kc.Namespace(getEnv("POD_NAMESPACE", "mesh")),
				kc.LabelSelector("app=comm"),
				kc.KubeConfig(flagconf),
			),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := cfg.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := cfg.Scan(&bc); err != nil {
		panic(err)
	}

	app, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}

}
