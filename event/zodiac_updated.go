package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	ZodiacUpdatedTopic = "master.zodiac"
)

type ZodiacUpdatedEvent struct {
	ZodiacUpdated struct{}  `json:"zodiac_updated"` // marking
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ZodiacUpdatedEvent) Topic() string {
	return ZodiacUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *ZodiacUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ZodiacUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
