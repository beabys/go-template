package api

import (
	"context"
	"net/http"

	hwproto "github.com/beabys/go-template/proto/gen/go/hello_world/v1"
)

// Used to implement the methods of helloworld interface

// from openapi

// HelloWorld implements the method Hello World
func (hs *HttpServer) HelloWorld(w http.ResponseWriter, r *http.Request) {
	hello, err := hs.HelloWorldSvc.GetHelloWorld(r.Context())
	if err != nil {
		hs.Logger.Error("error from HelloWorldSvc", err)
		errorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}
	response := map[string]interface{}{
		"hello": hello.Hello,
	}
	successResponseJSON(w, response)
}

// from grpc

// GetHelloWorld implements the method Get Hello World
func (hg *GRPCServer) GetHelloWorld(ctx context.Context, _ *hwproto.HelloWorldRequest) (*hwproto.HelloWorldResponse, error) {
	hello, err := hg.HelloWorldSvc.GetHelloWorld(ctx)
	if err != nil {
		hg.Logger.Error("error from HelloWorldSvc", err)
		return nil, err
	}
	return &hwproto.HelloWorldResponse{
		Hello: hello.Hello,
	}, nil
}
