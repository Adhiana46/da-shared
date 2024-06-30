package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	ZodiacDeletedTopic = "master.zodiac"
)

type ZodiacDeletedEvent struct {
	ZodiacDeleted struct{}  `json:"zodiac_deleted"` // marking
	Id            string    `json:"id"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ZodiacDeletedEvent) Topic() string {
	return ZodiacDeletedTopic
}

// GET KEY NAME, used for partition
func (e *ZodiacDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ZodiacDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
