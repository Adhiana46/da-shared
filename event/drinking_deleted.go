package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	DrinkingDeletedTopic = "master.drinking"
)

type DrinkingDeletedEvent struct {
	DrinkingDeleted struct{}  `json:"drinking_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *DrinkingDeletedEvent) Topic() string {
	return DrinkingDeletedTopic
}

// GET KEY NAME, used for partition
func (e *DrinkingDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *DrinkingDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
