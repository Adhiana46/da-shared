package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	HaveKidsDeletedTopic = "master.have_kids"
)

type HaveKidsDeletedEvent struct {
	HaveKidsDeleted struct{}  `json:"have_kids_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *HaveKidsDeletedEvent) Topic() string {
	return HaveKidsDeletedTopic
}

// GET KEY NAME, used for partition
func (e *HaveKidsDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *HaveKidsDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
