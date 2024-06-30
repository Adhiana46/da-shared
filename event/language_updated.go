package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	LanguageUpdatedTopic = "master.language"
)

type LanguageUpdatedEvent struct {
	LanguageUpdated string    `json:"language_updated"` // marking
	Id              string    `json:"id"`
	Code            string    `json:"code"`
	Name            string    `json:"name"`
	NativeName      string    `json:"native_name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *LanguageUpdatedEvent) Topic() string {
	return LanguageUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *LanguageUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *LanguageUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
