package helloworld

import (
	"context"

	"gitlab.com/beabys/go-http-template/internal/domain/models"
	"gitlab.com/beabys/go-http-template/pkg/logger"
)

type HelloWorldIntereface interface {
	GetHelloWorld(context.Context) (*models.HelloWorld, error)
}

type HelloWorld struct {
	logger logger.Logger
}
