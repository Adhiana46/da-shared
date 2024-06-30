package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	LookingForCreatedTopic = "master.looking_for"
)

type LookingForCreatedEvent struct {
	LookingForCreated struct{}  `json:"looking_for_created"` // marking
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	IsDeleted         bool      `json:"is_deleted"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *LookingForCreatedEvent) Topic() string {
	return LookingForCreatedTopic
}

// GET KEY NAME, used for partition
func (e *LookingForCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *LookingForCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
