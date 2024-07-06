package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterReligionTopic = "master.religion"
)

// ==============================
// Event: ReligionCreatedEvent
// ==============================
type ReligionCreatedEvent struct {
	ReligionCreated string    `json:"religion_created"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	IsDeleted       bool      `json:"is_deleted"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ReligionCreatedEvent) Topic() string {
	return MasterReligionTopic
}

// GET KEY NAME, used for partition
func (e *ReligionCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ReligionCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: ReligionDeletedEvent
// ==============================
type ReligionDeletedEvent struct {
	ReligionDeleted string    `json:"religion_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ReligionDeletedEvent) Topic() string {
	return MasterReligionTopic
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

// ==============================
// Event: ReligionUpdatedEvent
// ==============================
type ReligionUpdatedEvent struct {
	ReligionUpdated string    `json:"religion_updated"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ReligionUpdatedEvent) Topic() string {
	return MasterReligionTopic
}

// GET KEY NAME, used for partition
func (e *ReligionUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ReligionUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
