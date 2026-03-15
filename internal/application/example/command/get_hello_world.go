package command

import (
	"github.com/beabys/go-template/internal/domain/example/model"
)

type GetHelloWorldRequest struct{}

type GetHelloWorldResponse struct {
	ID      model.HelloWorldID
	Message string
}
