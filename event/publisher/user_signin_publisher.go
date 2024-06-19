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

type UserSignInPublisher interface {
	Publish(e *event.UserSignInEvent) error
}

type userSignInPublisher struct {
	sarama.SyncProducer
}

func NewUserSignInPublisher(producer sarama.SyncProducer) UserSignInPublisher {
	return &userSignInPublisher{
		SyncProducer: producer,
	}
}

func (u *userSignInPublisher) path() string {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	chunks := strings.Split(fn.Name(), ".")
	return "userSignInPublisher:" + chunks[len(chunks)-1]
}

func (u *userSignInPublisher) Publish(e *event.UserSignInEvent) error {
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
