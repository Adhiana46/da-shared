package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	GenderUpdatedTopic = "master.gender"
)

type GenderUpdatedEvent struct {
	GenderUpdated string    `json:"gender_updated"` // marking
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *GenderUpdatedEvent) Topic() string {
	return GenderUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *GenderUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *GenderUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
