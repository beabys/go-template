package api

import (
	"context"
	"net/http"

	v1 "gitlab.com/beabys/go-http-template/internal/api/v1"
	hwproto "gitlab.com/beabys/go-http-template/proto/gen/go/hello_world/v1"
)

// Used to implement the methods of helloworld interface

// from openapi

// HelloWorld implements the method Hello World
func (hs *HttpServer) HelloWorld(w http.ResponseWriter, r *http.Request) *v1.Response {
	response := v1.CommonResponse{}
	hello, err := hs.HelloWorldSvc.GetHelloWorld(r.Context())
	if err != nil {
		hs.Logger.Error(err)
		response.Data = map[string]interface{}{
			"error": err.Error(),
		}
		return v1.HelloWorldJSON5xxResponse(response)
	}
	response.Success = true
	response.Data = map[string]interface{}{
		"hello": hello.Hello,
	}
	return v1.HelloWorldJSON200Response(response)
}

// from grpc

// GetHelloWorld implements the method Get Hello World
func (hg *GRPCServer) GetHelloWorld(ctx context.Context, _ *hwproto.HelloWorldRequest) (*hwproto.HelloWorldResponse, error) {
	hello, err := hg.HelloWorldSvc.GetHelloWorld(ctx)
	if err != nil {
		hg.Logger.Error(err)
		return nil, err
	}
	return &hwproto.HelloWorldResponse{
		Hello: hello.Hello,
	}, nil
}
