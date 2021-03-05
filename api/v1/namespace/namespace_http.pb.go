// Code generated by protoc-gen-go-http. DO NOT EDIT.

package namespace

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	v1 "k8s.io/api/core/v1"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type NamespacesHandler interface {
	Create(context.Context, *v1.Namespace) (*Response, error)

	Delete(context.Context, *DeleteKind) (*Response, error)

	Get(context.Context, *GetKind) (*v1.Namespace, error)

	List(context.Context, *ListOption) (*NamespaceList, error)

	Update(context.Context, *v1.Namespace) (*Response, error)
}

func NewNamespacesHandler(srv NamespacesHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/cratos.api.v1.Namespace/List", func(w http.ResponseWriter, r *http.Request) {
		var in ListOption

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*ListOption))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, out); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/cratos.api.v1.Namespace/Get", func(w http.ResponseWriter, r *http.Request) {
		var in GetKind

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*GetKind))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, out); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/cratos.api.v1.Namespace/Create", func(w http.ResponseWriter, r *http.Request) {
		var in v1.Namespace

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*v1.Namespace))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, out); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/cratos.api.v1.Namespace/Update", func(w http.ResponseWriter, r *http.Request) {
		var in v1.Namespace

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*v1.Namespace))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, out); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/cratos.api.v1.Namespace/Delete", func(w http.ResponseWriter, r *http.Request) {
		var in DeleteKind

		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*DeleteKind))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		if err := h.Encode(w, r, out); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("DELETE")

	return r
}