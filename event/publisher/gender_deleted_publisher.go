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

type GenderDeletedPublisher interface {
	Publish(e *event.GenderDeletedEvent) error
}

type genderDeletedPublisher struct {
	sarama.SyncProducer
}

func NewGenderDeletedPublisher(producer sarama.SyncProducer) GenderDeletedPublisher {
	return &genderDeletedPublisher{
		SyncProducer: producer,
	}
}

func (u *genderDeletedPublisher) path() string {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	chunks := strings.Split(fn.Name(), ".")
	return "genderDeletedPublisher:" + chunks[len(chunks)-1]
}

func (u *genderDeletedPublisher) Publish(e *event.GenderDeletedEvent) error {
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
