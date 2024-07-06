package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterDrinkingTopic = "master.drinking"
)

// ==============================
// Event: DrinkingCreatedEvent
// ==============================
type DrinkingCreatedEvent struct {
	DrinkingCreated struct{}  `json:"drinking_created"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	IsDeleted       bool      `json:"is_deleted"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *DrinkingCreatedEvent) Topic() string {
	return MasterDrinkingTopic
}

// GET KEY NAME, used for partition
func (e *DrinkingCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *DrinkingCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: DrinkingDeletedEvent
// ==============================
type DrinkingDeletedEvent struct {
	DrinkingDeleted struct{}  `json:"drinking_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *DrinkingDeletedEvent) Topic() string {
	return MasterDrinkingTopic
}

// GET KEY NAME, used for partition
func (e *DrinkingDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *DrinkingDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: DrinkingUpdatedEvent
// ==============================
type DrinkingUpdatedEvent struct {
	DrinkingUpdated struct{}  `json:"drinking_updated"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *DrinkingUpdatedEvent) Topic() string {
	return MasterDrinkingTopic
}

// GET KEY NAME, used for partition
func (e *DrinkingUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *DrinkingUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
