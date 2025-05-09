package broker

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Broker struct {
	publisher  message.Publisher
	subscriber message.Subscriber
}

func NewBroker() *Broker {
	return newAWSBroker()
}

func (b *Broker) Publish(
	ctx context.Context,
	topic string,
	message *message.Message,
) error {
	return b.publisher.Publish(topic, message)
}

func (b *Broker) Subscribe(
	ctx context.Context,
	topic string,
) (<-chan *message.Message, error) {
	return b.subscriber.Subscribe(ctx, topic)
}
