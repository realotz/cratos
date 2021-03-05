package server

import (
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	gatewayPb "github.com/realotz/cratos/api/v1/gateway"
	nsPb "github.com/realotz/cratos/api/v1/namespace"
	"github.com/realotz/cratos/internal/conf"
	"github.com/realotz/cratos/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server,
	gateway *service.GatewayService,
	ns *service.NamespacesService,
) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	encoding.RegisterCodec(&codec{})
	s := http.NewServer(opts...)
	m := http.Middleware(
		middleware.Chain(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(),
		),
	)
	s.HanldePrefix("/cratos.api.v1.Gateway", gatewayPb.NewGatewaysHandler(gateway, m))
	s.HanldePrefix("/cratos.api.v1.Namespace", nsPb.NewNamespacesHandler(ns, m))
	return s
}
