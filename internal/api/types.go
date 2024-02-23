package api

import (
	"gitlab.com/beabys/go-http-template/internal/app/config"
	helloworld "gitlab.com/beabys/go-http-template/internal/hello_world"
	"gitlab.com/beabys/quetzal"
)

// HttpServer is a struct of an Http Server
type HttpServer struct {
	Config        *config.Config
	Logger        quetzal.Logger
	HelloWorldSvc helloworld.HelloWorldIntereface
}
