package messaging

import (
	"main/internal/model"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type ContactProducer struct {
	Producer[*model.ContactEvent]
}

func NewContactProducer(producer *kafka.Producer, log *logrus.Logger) *ContactProducer {
	return &ContactProducer{
		Producer: Producer[*model.ContactEvent]{
			Producer: producer,
			Topic:    "contacs",
			Log:      log,
		},
	}
}
