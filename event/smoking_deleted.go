package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	SmokingDeletedTopic = "master.smoking"
)

type SmokingDeletedEvent struct {
	SmokingDeleted struct{}  `json:"smoking_deleted"` // marking
	Id             string    `json:"id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *SmokingDeletedEvent) Topic() string {
	return SmokingDeletedTopic
}

// GET KEY NAME, used for partition
func (e *SmokingDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *SmokingDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
