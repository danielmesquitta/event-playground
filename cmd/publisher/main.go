package main

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/danielmesquitta/event-playground/internal/app/listener/topic"
	"github.com/danielmesquitta/event-playground/internal/domain/entity"
	"github.com/danielmesquitta/event-playground/internal/provider/broker/aws"
)

func main() {
	broker := aws.NewBroker()

	user := entity.User{
		ID:        uuid.NewString(),
		Email:     "john.doe@example.com",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
	}

	if err := broker.Publish(
		context.Background(),
		string(topic.TopicUserCreated),
		user,
	); err != nil {
		panic(err)
	}
}
