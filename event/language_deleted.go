package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	LanguageDeletedTopic = "master.language"
)

type LanguageDeletedEvent struct {
	LanguageDeleted string    `json:"language_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *LanguageDeletedEvent) Topic() string {
	return LanguageDeletedTopic
}

// GET KEY NAME, used for partition
func (e *LanguageDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *LanguageDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
