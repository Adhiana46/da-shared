package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	DrinkingUpdatedTopic = "master.drinking"
)

type DrinkingUpdatedEvent struct {
	DrinkingUpdated struct{}  `json:"drinking_updated"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *DrinkingUpdatedEvent) Topic() string {
	return DrinkingUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *DrinkingUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *DrinkingUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
