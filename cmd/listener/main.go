package main

import (
	"context"

	"github.com/danielmesquitta/event-playground/internal/app/listener"
	"github.com/danielmesquitta/event-playground/internal/app/listener/handler"
	"github.com/danielmesquitta/event-playground/internal/pkg/broker"
)

func main() {
	broker := broker.NewBroker()
	listener := listener.NewListener(
		broker,
		handler.NewUserCreatedHandler(),
	)
	if err := listener.Run(context.Background()); err != nil {
		panic(err)
	}
}
