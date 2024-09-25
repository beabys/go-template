package repository

import (
	"context"

	"gitlab.com/beabys/go-template/internal/domain/model"
	"gitlab.com/beabys/go-template/pkg/database"
	"gitlab.com/beabys/go-template/pkg/logger"
)

type RepositoryHelloWorld struct {
	logger logger.Logger
	Db     database.Database
}

type RepositoryIntereface interface {
	SaveHelloWorld(context.Context, *model.HelloWorld) error
}
