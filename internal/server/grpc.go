package server

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	gatewayPb "github.com/realotz/cratos/api/v1/gateway"
	nsPb "github.com/realotz/cratos/api/v1/namespace"
	"github.com/realotz/cratos/internal/conf"
	"github.com/realotz/cratos/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server,
	gateway *service.GatewayService,
	namespace *service.NamespacesService,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				status.Server(),
				tracing.Server(),
				logging.Server(),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	gatewayPb.RegisterGatewaysServer(srv, gateway)
	nsPb.RegisterNamespacesServer(srv, namespace)
	return srv
}
