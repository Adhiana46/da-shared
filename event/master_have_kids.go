package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterHaveKidsTopic = "master.have_kids"
)

// ==============================
// Event: HaveKidsCreatedEvent
// ==============================
type HaveKidsCreatedEvent struct {
	HaveKidsCreated struct{}  `json:"have_kids_created"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	IsDeleted       bool      `json:"is_deleted"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *HaveKidsCreatedEvent) Topic() string {
	return MasterHaveKidsTopic
}

// GET KEY NAME, used for partition
func (e *HaveKidsCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *HaveKidsCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: HaveKidsDeletedEvent
// ==============================
type HaveKidsDeletedEvent struct {
	HaveKidsDeleted struct{}  `json:"have_kids_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *HaveKidsDeletedEvent) Topic() string {
	return MasterHaveKidsTopic
}

// GET KEY NAME, used for partition
func (e *HaveKidsDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *HaveKidsDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: HaveKidsUpdatedEvent
// ==============================
type HaveKidsUpdatedEvent struct {
	HaveKidsUpdated struct{}  `json:"have_kids_updated"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *HaveKidsUpdatedEvent) Topic() string {
	return MasterHaveKidsTopic
}

// GET KEY NAME, used for partition
func (e *HaveKidsUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *HaveKidsUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
