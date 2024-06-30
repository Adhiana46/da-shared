package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	ReligionUpdatedTopic = "master.religion"
)

type ReligionUpdatedEvent struct {
	ReligionUpdated string    `json:"religion_updated"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ReligionUpdatedEvent) Topic() string {
	return ReligionUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *ReligionUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ReligionUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
