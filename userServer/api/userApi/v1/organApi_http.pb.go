// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.2

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type OrganQueryHTTPServer interface {
	GetUser(context.Context, *GetOrganRequest) (*GetOrganReply, error)
}

func RegisterOrganQueryHTTPServer(s *http.Server, srv OrganQueryHTTPServer) {
	r := s.Route("/")
	r.GET("/OrganApi/{id}", _OrganQuery_GetUser0_HTTP_Handler(srv))
}

func _OrganQuery_GetUser0_HTTP_Handler(srv OrganQueryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOrganRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/userApi.v1.OrganQuery/GetUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetOrganRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOrganReply)
		return ctx.Result(200, reply)
	}
}

type OrganQueryHTTPClient interface {
	GetUser(ctx context.Context, req *GetOrganRequest, opts ...http.CallOption) (rsp *GetOrganReply, err error)
}

type OrganQueryHTTPClientImpl struct {
	cc *http.Client
}

func NewOrganQueryHTTPClient(client *http.Client) OrganQueryHTTPClient {
	return &OrganQueryHTTPClientImpl{client}
}

func (c *OrganQueryHTTPClientImpl) GetUser(ctx context.Context, in *GetOrganRequest, opts ...http.CallOption) (*GetOrganReply, error) {
	var out GetOrganReply
	pattern := "/OrganApi/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/userApi.v1.OrganQuery/GetUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
