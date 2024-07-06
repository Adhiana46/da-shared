package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterLanguageTopic = "master.language"
)

// ==============================
// Event: LanguageCreatedEvent
// ==============================
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
	return MasterLanguageTopic
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

// ==============================
// Event: LanguageDeletedEvent
// ==============================
type LanguageDeletedEvent struct {
	LanguageDeleted string    `json:"language_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *LanguageDeletedEvent) Topic() string {
	return MasterLanguageTopic
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

// ==============================
// Event: LanguageUpdatedEvent
// ==============================
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
	return MasterLanguageTopic
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
