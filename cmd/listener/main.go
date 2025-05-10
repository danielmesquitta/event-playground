package main

import (
	"context"

	"github.com/danielmesquitta/event-playground/internal/app/listener"
	"github.com/danielmesquitta/event-playground/internal/app/listener/handler"
	"github.com/danielmesquitta/event-playground/internal/domain/usecase"
	"github.com/danielmesquitta/event-playground/internal/pkg/broker"
	"github.com/danielmesquitta/event-playground/internal/pkg/gracefulshutdown"
)

func main() {
	ctx := gracefulshutdown.WithShutdownSignal(context.Background())

	broker := broker.NewBroker()

	sendWelcomeEmailUseCase := usecase.NewSendWelcomeEmail()

	userCreatedHandler := handler.NewUserCreatedHandler(sendWelcomeEmailUseCase)

	listener := listener.NewListener(
		broker,
		userCreatedHandler,
	)
	if err := listener.Run(ctx); err != nil {
		panic(err)
	}
}
