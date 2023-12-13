package messaging

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type ConsumerHandler func(messgae *kafka.Message) error


// func ConsumeTopic(ctx context.Context, consumer *kafka.Consumer, topic string, log *logrus.Logger, ha)