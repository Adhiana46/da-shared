package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	SmokingUpdatedTopic = "master.smoking"
)

type SmokingUpdatedEvent struct {
	SmokingUpdated struct{}  `json:"smoking_updated"` // marking
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *SmokingUpdatedEvent) Topic() string {
	return SmokingUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *SmokingUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *SmokingUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
