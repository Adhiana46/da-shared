package publisher

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"runtime"
	"strings"

	"github.com/IBM/sarama"
	"github.com/adhiana46/da-shared/event"
	"github.com/pkg/errors"
)

type GenderCreatedPublisher interface {
	Publish(e *event.GenderCreatedEvent) error
}

type genderCreatedPublisher struct {
	sarama.SyncProducer
}

func NewGenderCreatedPublisher(producer sarama.SyncProducer) GenderCreatedPublisher {
	return &genderCreatedPublisher{
		SyncProducer: producer,
	}
}

func (u *genderCreatedPublisher) path() string {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	chunks := strings.Split(fn.Name(), ".")
	return "genderCreatedPublisher:" + chunks[len(chunks)-1]
}

func (u *genderCreatedPublisher) Publish(e *event.GenderCreatedEvent) error {
	path := u.path()

	topic := e.Topic()
	key := e.Key()
	payload, err := json.Marshal(e)
	if err != nil {
		return errors.Wrap(err, path)
	}

	msg := sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(payload),
	}

	partition, offset, err := u.SendMessage(&msg)
	if err != nil {
		return errors.Wrap(err, path)
	}
	slog.Info(fmt.Sprintf("Event published to topic: '%s' partition %d with offset %d", topic, partition, offset))

	return nil
}
