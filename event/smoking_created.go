package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	SmokingCreatedTopic = "master.smoking"
)

type SmokingCreatedEvent struct {
	SmokingCreated struct{}  `json:"smoking_created"` // marking
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	IsDeleted      bool      `json:"is_deleted"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *SmokingCreatedEvent) Topic() string {
	return SmokingCreatedTopic
}

// GET KEY NAME, used for partition
func (e *SmokingCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *SmokingCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
