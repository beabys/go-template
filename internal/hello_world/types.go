package helloworld

import (
	"context"

	"github.com/beabys/go-template/internal/domain/model"
	"github.com/beabys/go-template/internal/hello_world/repository"
	"github.com/beabys/go-template/pkg/logger"
)

type HelloWorldIntereface interface {
	GetHelloWorld(context.Context) (*model.HelloWorld, error)
}

type HelloWorld struct {
	logger     logger.Logger
	repository repository.RepositoryIntereface
}
