package broker

import (
	"context"
	"fmt"
	"net/url"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-aws/sns"
	"github.com/ThreeDotsLabs/watermill-aws/sqs"
	"github.com/aws/aws-sdk-go-v2/aws"
	amazonsns "github.com/aws/aws-sdk-go-v2/service/sns"
	amazonsqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	transport "github.com/aws/smithy-go/endpoints"
	"github.com/samber/lo"
)

func newAWSBroker() *Broker {
	logger := watermill.NewStdLogger(false, false)

	snsOpts := []func(*amazonsns.Options){
		amazonsns.WithEndpointResolverV2(sns.OverrideEndpointResolver{
			Endpoint: transport.Endpoint{
				URI: *lo.Must(url.Parse("http://localhost:4566")),
			},
		}),
	}

	sqsOpts := []func(*amazonsqs.Options){
		amazonsqs.WithEndpointResolverV2(sqs.OverrideEndpointResolver{
			Endpoint: transport.Endpoint{
				URI: *lo.Must(url.Parse("http://localhost:4566")),
			},
		}),
	}

	topicResolver, err := sns.NewGenerateArnTopicResolver(
		"000000000000",
		"us-east-1",
	)
	if err != nil {
		panic(err)
	}

	subscriberConfig := sns.SubscriberConfig{
		AWSConfig: aws.Config{
			Credentials: aws.AnonymousCredentials{},
		},
		OptFns:        snsOpts,
		TopicResolver: topicResolver,
		GenerateSqsQueueName: func(ctx context.Context, snsTopic sns.TopicArn) (string, error) {
			topic, err := sns.ExtractTopicNameFromTopicArn(snsTopic)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("%v-subscriber", topic), nil
		},
	}

	sqsSubscriberConfig := sqs.SubscriberConfig{
		AWSConfig: aws.Config{
			Credentials: aws.AnonymousCredentials{},
		},
		OptFns: sqsOpts,
	}

	subscriber, err := sns.NewSubscriber(
		subscriberConfig,
		sqsSubscriberConfig,
		logger,
	)
	if err != nil {
		panic(err)
	}

	publisherConfig := sns.PublisherConfig{
		AWSConfig: aws.Config{
			Credentials: aws.AnonymousCredentials{},
		},
		OptFns:        snsOpts,
		TopicResolver: topicResolver,
	}

	publisher, err := sns.NewPublisher(publisherConfig, logger)
	if err != nil {
		panic(err)
	}

	return &Broker{
		publisher:  publisher,
		subscriber: subscriber,
	}
}
