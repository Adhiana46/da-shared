package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	InterestDeletedTopic = "master.interest"
)

type InterestDeletedEvent struct {
	InterestDeleted struct{}  `json:"interest_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *InterestDeletedEvent) Topic() string {
	return InterestDeletedTopic
}

// GET KEY NAME, used for partition
func (e *InterestDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *InterestDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
