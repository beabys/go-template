package helloworld

import (
	"net/http"

	"gitlab.com/beabys/go-http-template/pkg/logger"
)

func NewHelloWorld(logger logger.Logger) *HelloWorld {
	return &HelloWorld{
		logger: logger,
	}
}

func (hw *HelloWorld) GetHelloWorld(_ *http.Request) error {
	hw.logger.Info("logging the Hello World get Method")
	return nil
}
