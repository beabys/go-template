package helloworld

import (
	"net/http"

	"gitlab.com/beabys/quetzal"
)

type HelloWorldIntereface interface {
	GetHelloWorld(r *http.Request) error
}

type HelloWorld struct {
	logger quetzal.Logger
}
