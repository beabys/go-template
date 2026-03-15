package repository

import (
	"context"

	"github.com/beabys/go-template/internal/domain/example/model"
)

type HelloWorldRepository interface {
	SaveHelloWorld(ctx context.Context, helloWorld *model.HelloWorld) error
	GetHelloWorld(ctx context.Context, id model.HelloWorldID) (*model.HelloWorld, error)
}
