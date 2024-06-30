package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	GenderCreatedTopic = "master.gender"
)

type GenderCreatedEvent struct {
	GenderCreated string    `json:"gender_created"` // marking
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	IsDeleted     bool      `json:"is_deleted"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *GenderCreatedEvent) Topic() string {
	return GenderCreatedTopic
}

// GET KEY NAME, used for partition
func (e *GenderCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *GenderCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
