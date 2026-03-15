package model

import (
	"fmt"
	"time"
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

func NewHelloWorld(message string) *HelloWorld {
	now := time.Now().UTC()
	return &HelloWorld{
		ID:         NewHelloWorldID(),
		Message:    message,
		Timestamps: Timestamps{CreatedAt: now, UpdatedAt: now},
	}
}

func (h *HelloWorld) String() string {
	return fmt.Sprintf("HelloWorld(ID: %s, Message: %s)", h.ID, h.Message)
}
