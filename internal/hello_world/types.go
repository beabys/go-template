package helloworld

import (
	"net/http"

	"gitlab.com/beabys/go-http-template/pkg/logger"
)

type HelloWorldIntereface interface {
	GetHelloWorld(r *http.Request) error
}

type HelloWorld struct {
	logger logger.Logger
}
