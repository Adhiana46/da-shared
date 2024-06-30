package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	LanguageCreatedTopic = "master.language"
)

type LanguageCreatedEvent struct {
	LanguageCreated string    `json:"language_created"` // marking
	Id              string    `json:"id"`
	Code            string    `json:"code"`
	Name            string    `json:"name"`
	NativeName      string    `json:"native_name"`
	IsDeleted       bool      `json:"is_deleted"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *LanguageCreatedEvent) Topic() string {
	return LanguageCreatedTopic
}

// GET KEY NAME, used for partition
func (e *LanguageCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *LanguageCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
