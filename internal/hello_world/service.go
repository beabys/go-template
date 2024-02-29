package helloworld

import (
	"context"

	"gitlab.com/beabys/go-http-template/internal/domain/models"
	"gitlab.com/beabys/go-http-template/pkg/logger"
)

func NewHelloWorld(logger logger.Logger) *HelloWorld {
	return &HelloWorld{
		logger: logger,
	}
}

func (hw *HelloWorld) GetHelloWorld(_ context.Context) (*models.HelloWorld, error) {
	hw.logger.Info("logging the Hello World get Method")
	return &models.HelloWorld{
		Hello: "world",
	}, nil
}
