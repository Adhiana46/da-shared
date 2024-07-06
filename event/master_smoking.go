package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterSmokingTopic = "master.smoking"
)

// ==============================
// Event: SmokingCreatedEvent
// ==============================
type SmokingCreatedEvent struct {
	SmokingCreated struct{}  `json:"smoking_created"` // marking
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	IsDeleted      bool      `json:"is_deleted"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *SmokingCreatedEvent) Topic() string {
	return MasterSmokingTopic
}

// GET KEY NAME, used for partition
func (e *SmokingCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *SmokingCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: SmokingDeletedEvent
// ==============================
type SmokingDeletedEvent struct {
	SmokingDeleted struct{}  `json:"smoking_deleted"` // marking
	Id             string    `json:"id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *SmokingDeletedEvent) Topic() string {
	return MasterSmokingTopic
}

// GET KEY NAME, used for partition
func (e *SmokingDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *SmokingDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: SmokingUpdatedEvent
// ==============================
type SmokingUpdatedEvent struct {
	SmokingUpdated struct{}  `json:"smoking_updated"` // marking
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *SmokingUpdatedEvent) Topic() string {
	return MasterSmokingTopic
}

// GET KEY NAME, used for partition
func (e *SmokingUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *SmokingUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
