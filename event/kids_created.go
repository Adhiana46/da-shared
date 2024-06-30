package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	KidsCreatedTopic = "master.kids"
)

type KidsCreatedEvent struct {
	KidsCreated struct{}  `json:"kids_created"` // marking
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *KidsCreatedEvent) Topic() string {
	return KidsCreatedTopic
}

// GET KEY NAME, used for partition
func (e *KidsCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *KidsCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
