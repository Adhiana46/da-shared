package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	KidsUpdatedTopic = "master.kids"
)

type KidsUpdatedEvent struct {
	KidsUpdated struct{}  `json:"kids_updated"` // marking
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *KidsUpdatedEvent) Topic() string {
	return KidsUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *KidsUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *KidsUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
