package helloworld

import (
	"context"

	"gitlab.com/beabys/go-http-template/internal/domain/model"
	"gitlab.com/beabys/go-http-template/pkg/logger"
)

type HelloWorldIntereface interface {
	GetHelloWorld(context.Context) (*model.HelloWorld, error)
}

type HelloWorld struct {
	logger logger.Logger
}
