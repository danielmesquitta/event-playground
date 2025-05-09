package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/danielmesquitta/event-playground/internal/app/listener/handler"
	"github.com/danielmesquitta/event-playground/internal/app/listener/topic"
	"github.com/danielmesquitta/event-playground/internal/pkg/broker"
)

type Listener struct {
	broker             *broker.Broker
	userCreatedHandler *handler.UserCreatedHandler
}

func NewListener(
	broker *broker.Broker,
	userCreatedHandler *handler.UserCreatedHandler,
) *Listener {
	return &Listener{
		broker:             broker,
		userCreatedHandler: userCreatedHandler,
	}
}

func (l *Listener) Run(ctx context.Context) error {
	topics := map[topic.Topic]handler.Handler{
		topic.TopicUserCreated: l.userCreatedHandler,
	}

	for topic, handler := range topics {
		if err := l.subscribe(ctx, string(topic), handler); err != nil {
			return fmt.Errorf("failed to subscribe to topic %s: %w", topic, err)
		}
	}

	return nil
}

func (l *Listener) subscribe(
	ctx context.Context,
	topic string,
	handler handler.Handler,
) error {
	messages, err := l.broker.Subscribe(
		ctx,
		topic,
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic %s: %w", topic, err)
	}

	go func() {
		for message := range messages {
			l.handle(ctx, topic, message, handler)
		}
	}()

	return nil
}

func (l *Listener) handle(
	ctx context.Context,
	topic string,
	message *message.Message,
	handler handler.Handler,
) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := handler.Handle(
		ctx,
		json.RawMessage(message.Payload),
	); err != nil {
		slog.Error(
			"failed to handle message",
			"topic", topic,
			"error", err,
		)
	}

	message.Ack()
}
