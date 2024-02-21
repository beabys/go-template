package api

import (
	"gitlab.com/beabys/go-http-template/internal/app/config"
	helloworld "gitlab.com/beabys/go-http-template/internal/hello_world"
	"gitlab.com/beabys/quetzal"
)

// NewHttpServer returns a new pointer of HttpServer
func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

// SetConfig is a setter function to set Configs
func (hs *HttpServer) SetConfig(c *config.Config) *HttpServer {
	hs.Config = c
	return hs
}

// SetLogger is a setter function to set the Logger
func (hs *HttpServer) SetLogger(l quetzal.Logger) *HttpServer {
	hs.Logger = l
	return hs
}

func (hs *HttpServer) SetHelloWorldService(hw helloworld.HelloWorldIntereface) *HttpServer {
	hs.HelloWorldSvc = hw
	return hs
}
