// Code generated by protoc-gen-go-http. DO NOT EDIT.

package book

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
// context./http.
const _ = http1.SupportPackageIsVersion1

type IstioServiceHTTPServer interface {
	CheckInfo(context.Context, *Request) (*Response, error)
}

func RegisterIstioServiceHTTPServer(s http1.ServiceRegistrar, srv IstioServiceHTTPServer) {
	s.RegisterService(&_HTTP_IstioService_serviceDesc, srv)
}

func _HTTP_IstioService_CheckInfo_0(srv interface{}, ctx context.Context, req *http.Request, dec func(interface{}) error) (interface{}, error) {
	var in Request

	if err := http1.BindForm(req, &in); err != nil {
		return nil, err
	}

	out, err := srv.(IstioServiceServer).CheckInfo(ctx, &in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _HTTP_IstioService_serviceDesc = http1.ServiceDesc{
	ServiceName: "cratos.service.v1.IstioService",
	Methods: []http1.MethodDesc{

		{
			Path:    "/cratos.service.v1.IstioService/CheckInfo",
			Method:  "POST",
			Handler: _HTTP_IstioService_CheckInfo_0,
		},
	},
	Metadata: "api/v1/istio.proto",
}