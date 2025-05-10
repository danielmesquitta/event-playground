package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/danielmesquitta/event-playground/internal/app/listener/topic"
	"github.com/danielmesquitta/event-playground/internal/domain/entity"
	"github.com/danielmesquitta/event-playground/internal/provider/broker/aws"
)

func main() {
	broker := aws.NewAWSBroker()

	user := entity.User{
		ID:        uuid.NewString(),
		Email:     "john.doe@example.com",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
	}

	payload, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	broker.Publish(
		context.Background(),
		string(topic.TopicUserCreated),
		message.NewMessage(
			uuid.NewString(),
			payload,
		),
	)
}
