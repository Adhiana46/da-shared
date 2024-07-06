package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterLookingForTopic = "master.looking_for"
)

// ==============================
// Event: LookingForCreatedEvent
// ==============================
type LookingForCreatedEvent struct {
	LookingForCreated struct{}  `json:"looking_for_created"` // marking
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	IsDeleted         bool      `json:"is_deleted"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *LookingForCreatedEvent) Topic() string {
	return MasterLookingForTopic
}

// GET KEY NAME, used for partition
func (e *LookingForCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *LookingForCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: LookingForDeletedEvent
// ==============================
type LookingForDeletedEvent struct {
	LookingForDeleted struct{}  `json:"looking_for_deleted"` // marking
	Id                string    `json:"id"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *LookingForDeletedEvent) Topic() string {
	return MasterLookingForTopic
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

// ==============================
// Event: LookingForUpdatedEvent
// ==============================
type LookingForUpdatedEvent struct {
	LookingForUpdated struct{}  `json:"looking_for_updated"` // marking
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *LookingForUpdatedEvent) Topic() string {
	return MasterLookingForTopic
}

// GET KEY NAME, used for partition
func (e *LookingForUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *LookingForUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
