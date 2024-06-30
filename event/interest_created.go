package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	InterestCreatedTopic = "master.interest"
)

type InterestCreatedEvent struct {
	InterestCreated struct{}  `json:"interest_created"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	IsDeleted       bool      `json:"is_deleted"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *InterestCreatedEvent) Topic() string {
	return InterestCreatedTopic
}

// GET KEY NAME, used for partition
func (e *InterestCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *InterestCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
