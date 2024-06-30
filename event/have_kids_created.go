package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	HaveKidsCreatedTopic = "master.have_kids"
)

type HaveKidsCreatedEvent struct {
	HaveKidsCreated struct{}  `json:"have_kids_created"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	IsDeleted       bool      `json:"is_deleted"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *HaveKidsCreatedEvent) Topic() string {
	return HaveKidsCreatedTopic
}

// GET KEY NAME, used for partition
func (e *HaveKidsCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *HaveKidsCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
