package usecase

import (
	"context"
	"time"

	"github.com/beabys/go-template/internal/application/example/command"
	"github.com/beabys/go-template/internal/application/example/repository"
	"github.com/beabys/go-template/pkg/logger"
)

const defaultTimeout = 5 * time.Second

type ExampleServiceHandler interface {
	GetHelloWorld(ctx context.Context, req *command.GetHelloWorldRequest) (*command.GetHelloWorldResponse, error)
}

type ExampleService struct {
	logger     logger.Logger
	repository repository.HelloWorldRepository
}
