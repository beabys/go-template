package helloworld

import (
	"context"

	"github.com/beabys/go-template/internal/domain/model"
	"github.com/beabys/go-template/internal/hello_world/repository"
	"github.com/beabys/go-template/pkg/logger"
)

func NewHelloWorld(logger logger.Logger, repository repository.RepositoryIntereface) *HelloWorld {
	return &HelloWorld{
		logger:     logger,
		repository: repository,
	}
}

func (hw *HelloWorld) GetHelloWorld(ctx context.Context) (*model.HelloWorld, error) {
	hw.logger.Info("logging the Hello World get Method")
	helloWorld := &model.HelloWorld{}

	if err := hw.repository.SaveHelloWorld(ctx, helloWorld); err != nil {
		return nil, err
	}
	return helloWorld, nil
}
