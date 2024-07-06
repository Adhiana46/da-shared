package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterGenderTopic = "master.gender"
)

// ==============================
// Event: GenderCreatedEvent
// ==============================
type GenderCreatedEvent struct {
	GenderCreated string    `json:"gender_created"` // marking
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	IsDeleted     bool      `json:"is_deleted"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *GenderCreatedEvent) Topic() string {
	return MasterGenderTopic
}

// GET KEY NAME, used for partition
func (e *GenderCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *GenderCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: GenderDeletedEvent
// ==============================
type GenderDeletedEvent struct {
	GenderDeleted string    `json:"gender_deleted"` // marking
	Id            string    `json:"id"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *GenderDeletedEvent) Topic() string {
	return MasterGenderTopic
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

// ==============================
// Event: GenderUpdatedEvent
// ==============================
type GenderUpdatedEvent struct {
	GenderUpdated string    `json:"gender_updated"` // marking
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *GenderUpdatedEvent) Topic() string {
	return MasterGenderTopic
}

// GET KEY NAME, used for partition
func (e *GenderUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *GenderUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
