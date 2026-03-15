package handler

import (
	"context"

	"github.com/beabys/go-template/internal/application/example/command"
)

type ExampleServiceHandler interface {
	GetHelloWorld(ctx context.Context, req *command.GetHelloWorldRequest) (*command.GetHelloWorldResponse, error)
}
