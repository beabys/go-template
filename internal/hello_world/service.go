package helloworld

import (
	"context"

	"gitlab.com/beabys/go-template/internal/domain/model"
	"gitlab.com/beabys/go-template/pkg/logger"
)

func NewHelloWorld(logger logger.Logger) *HelloWorld {
	return &HelloWorld{
		logger: logger,
	}
}

func (hw *HelloWorld) GetHelloWorld(_ context.Context) (*model.HelloWorld, error) {
	hw.logger.Info("logging the Hello World get Method")
	return &model.HelloWorld{
		Hello: "world",
	}, nil
}
