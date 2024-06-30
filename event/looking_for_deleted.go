package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	LookingForDeletedTopic = "master.looking_for"
)

type LookingForDeletedEvent struct {
	LookingForDeleted struct{}  `json:"looking_for_deleted"` // marking
	Id                string    `json:"id"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *LookingForDeletedEvent) Topic() string {
	return LookingForDeletedTopic
}

// GET KEY NAME, used for partition
func (e *LookingForDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *LookingForDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
