package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	GenderDeletedTopic = "master.gender"
)

type GenderDeletedEvent struct {
	GenderDeleted string    `json:"gender_deleted"` // marking
	Id            string    `json:"id"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *GenderDeletedEvent) Topic() string {
	return GenderDeletedTopic
}

// GET KEY NAME, used for partition
func (e *GenderDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *GenderDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
