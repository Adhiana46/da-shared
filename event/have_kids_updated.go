package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	HaveKidsUpdatedTopic = "master.have_kids"
)

type HaveKidsUpdatedEvent struct {
	HaveKidsUpdated struct{}  `json:"have_kids_updated"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *HaveKidsUpdatedEvent) Topic() string {
	return HaveKidsUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *HaveKidsUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *HaveKidsUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
