package main

import (
	"context"

	"github.com/danielmesquitta/event-playground/internal/app/listener"
	"github.com/danielmesquitta/event-playground/internal/app/listener/handler"
	"github.com/danielmesquitta/event-playground/internal/domain/usecase"
	"github.com/danielmesquitta/event-playground/internal/pkg/gracefulshutdown"
	"github.com/danielmesquitta/event-playground/internal/provider/broker/aws"
)

func main() {
	ctx := gracefulshutdown.WithShutdownSignal(context.Background())

	broker := aws.NewAWSBroker()

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
