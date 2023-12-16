package messaging

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ConsumerHandler func(messgae *kafka.Message) error

// func ConsumeTopic(ctx context.Context, consumer *kafka.Consumer, topic string, log *logrus.Logger, ha)
