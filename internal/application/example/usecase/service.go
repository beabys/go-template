package usecase

import (
	"context"
	"errors"

	"github.com/beabys/go-template/internal/application/example/command"
	"github.com/beabys/go-template/internal/application/example/repository"
	"github.com/beabys/go-template/internal/domain/example/model"
	"github.com/beabys/go-template/pkg/logger"
)

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

	helloWorld, err := model.NewHelloWorld("Hello, World!")
	if err != nil {
		s.logger.Error("failed to create hello world", err)
		return nil, err
	}

	if err := s.repository.SaveHelloWorld(ctx, helloWorld); err != nil {
		s.logger.Error("failed to save hello world", err)
		return nil, err
	}

	return &command.GetHelloWorldResponse{
		ID:      helloWorld.ID,
		Message: helloWorld.Message,
	}, nil
}
