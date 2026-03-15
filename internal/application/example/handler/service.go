package handler

import (
	"context"
	"errors"
	"time"

	"github.com/beabys/go-template/internal/application/example/command"
	"github.com/beabys/go-template/internal/application/example/repository"
	"github.com/beabys/go-template/internal/domain/example/model"
	"github.com/beabys/go-template/pkg/logger"
)

const defaultTimeout = 5 * time.Second

type ExampleService struct {
	logger     logger.Logger
	repository repository.HelloWorldRepository
}

func NewExampleService(logger logger.Logger, repository repository.HelloWorldRepository) *ExampleService {
	return &ExampleService{
		logger:     logger,
		repository: repository,
	}
}

func (s *ExampleService) GetHelloWorld(ctx context.Context, req *command.GetHelloWorldRequest) (*command.GetHelloWorldResponse, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	s.logger.Info("getting hello world")

	helloWorld := model.NewHelloWorld("Hello, World!")

	if err := s.repository.SaveHelloWorld(ctx, helloWorld); err != nil {
		s.logger.Error("failed to save hello world", err)
		return nil, err
	}

	return &command.GetHelloWorldResponse{
		ID:      helloWorld.ID,
		Message: helloWorld.Message,
	}, nil
}
