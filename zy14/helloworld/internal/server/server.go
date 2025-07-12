package server

import (
	"github.com/go-kratos/kratos/v2/log"
	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, service.NewUserService, service.NewDriverService)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *kratoshttp.Server {
	var opts = []kratoshttp.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, kratoshttp.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, kratoshttp.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, kratoshttp.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := kratoshttp.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	// 注册UserService和DriverService（如有需要，可按实际依赖注入参数添加）
	// v1.RegisterUserHTTPServer(srv, userService)
	// v1.RegisterDriverHTTPServer(srv, driverService)
	return srv
}
