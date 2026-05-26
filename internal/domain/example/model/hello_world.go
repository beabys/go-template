package model

import (
	"fmt"
	"time"

	"github.com/beabys/go-template/internal/domain/example"
)

type HelloWorldID string

type HelloWorld struct {
	ID      HelloWorldID
	Message string
	Timestamps
}

func NewHelloWorldID() HelloWorldID {
	return HelloWorldID(time.Now().Format("20060102150405"))
}

func NewHelloWorld(message string) (*HelloWorld, error) {
	if message == "" {
		return nil, example.ErrInvalidMessage
	}
	now := time.Now().UTC()
	return &HelloWorld{
		ID:         NewHelloWorldID(),
		Message:    message,
		Timestamps: Timestamps{CreatedAt: now, UpdatedAt: now},
	}, nil
}

func (h *HelloWorld) UpdateMessage(message string) error {
	if message == "" {
		return example.ErrInvalidMessage
	}
	h.Message = message
	h.UpdatedAt = time.Now().UTC()
	return nil
}

func (h *HelloWorld) String() string {
	return fmt.Sprintf("HelloWorld(ID: %s, Message: %s)", h.ID, h.Message)
}
