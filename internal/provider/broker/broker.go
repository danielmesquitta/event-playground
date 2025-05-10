package broker

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Broker interface {
	Publish(ctx context.Context, topic string, payload any) error
	Subscribe(
		ctx context.Context,
		topic string,
	) (<-chan *message.Message, error)
}
