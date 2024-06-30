package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	LookingForUpdatedTopic = "master.looking_for"
)

type LookingForUpdatedEvent struct {
	LookingForUpdated struct{}  `json:"looking_for_updated"` // marking
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *LookingForUpdatedEvent) Topic() string {
	return LookingForUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *LookingForUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *LookingForUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
