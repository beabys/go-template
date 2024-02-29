package api

import (
	"context"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"

	hwproto "gitlab.com/beabys/go-http-template/proto/gen/go/hello_world/v1"
)

// Used to implement the methods of helloworld interface

// from openapi

// HelloWorld implements the method Hello World
func (hs *HttpServer) HelloWorld(w http.ResponseWriter, r *http.Request) {
	hs.HelloWorldSvc.GetHelloWorld(r)
	SuccessResponseJSON(w, nil)
}

// from grpc

// GetHelloWorld implements the method Get Hello World
func (hg *GRPCServer) GetHelloWorld(_ context.Context, _ *emptypb.Empty) (*hwproto.HelloWorldResponse, error) {
	return &hwproto.HelloWorldResponse{
		Data: "Hello World",
	}, nil
}
