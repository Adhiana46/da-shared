package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterInterestTopic = "master.interest"
)

// ==============================
// Event: InterestCreatedEvent
// ==============================
type InterestCreatedEvent struct {
	InterestCreated struct{}  `json:"interest_created"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	IsDeleted       bool      `json:"is_deleted"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *InterestCreatedEvent) Topic() string {
	return MasterInterestTopic
}

// GET KEY NAME, used for partition
func (e *InterestCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *InterestCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: InterestDeletedEvent
// ==============================
type InterestDeletedEvent struct {
	InterestDeleted struct{}  `json:"interest_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *InterestDeletedEvent) Topic() string {
	return MasterInterestTopic
}

// GET KEY NAME, used for partition
func (e *InterestDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *InterestDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: InterestUpdatedEvent
// ==============================
type InterestUpdatedEvent struct {
	InterestUpdated struct{}  `json:"interest_updated"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *InterestUpdatedEvent) Topic() string {
	return MasterInterestTopic
}

// GET KEY NAME, used for partition
func (e *InterestUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *InterestUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
