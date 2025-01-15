package repository

import (
	"context"

	"github.com/beabys/go-template/internal/domain/model"
	"github.com/beabys/go-template/pkg/database"
	"github.com/beabys/go-template/pkg/logger"
)

type RepositoryHelloWorld struct {
	logger logger.Logger
	Db     database.Database
}

type RepositoryIntereface interface {
	SaveHelloWorld(context.Context, *model.HelloWorld) error
}
