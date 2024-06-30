package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	ZodiacCreatedTopic = "master.zodiac"
)

type ZodiacCreatedEvent struct {
	ZodiacCreated struct{}  `json:"zodiac_created"` // marking
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	IsDeleted     bool      `json:"is_deleted"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ZodiacCreatedEvent) Topic() string {
	return ZodiacCreatedTopic
}

// GET KEY NAME, used for partition
func (e *ZodiacCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ZodiacCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
