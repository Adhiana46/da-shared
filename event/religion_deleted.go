package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	ReligionDeletedTopic = "master.religion"
)

type ReligionDeletedEvent struct {
	ReligionDeleted string    `json:"religion_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ReligionDeletedEvent) Topic() string {
	return ReligionDeletedTopic
}

// GET KEY NAME, used for partition
func (e *ReligionDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ReligionDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
