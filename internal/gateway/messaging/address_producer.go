package messaging

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
	"main/internal/model"
)

type AddressProduser struct {
	Producer[*model.AddressEvent]
}

func NewAddressProducer(producer *kafka.Producer, log *logrus.Logger) *AddressProduser {
	return &AddressProduser{
		Producer: Producer[*model.AddressEvent]{
			Producer: producer,
			Topic:    "addresses",
			Log:      log,
		},
	}
}
