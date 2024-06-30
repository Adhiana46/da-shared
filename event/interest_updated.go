package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	InterestUpdatedTopic = "master.interest"
)

type InterestUpdatedEvent struct {
	InterestUpdated struct{}  `json:"interest_updated"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *InterestUpdatedEvent) Topic() string {
	return InterestUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *InterestUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *InterestUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
