package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterKidsTopic = "master.kids"
)

// ==============================
// Event: KidsCreatedEvent
// ==============================
type KidsCreatedEvent struct {
	KidsCreated struct{}  `json:"kids_created"` // marking
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *KidsCreatedEvent) Topic() string {
	return MasterKidsTopic
}

// GET KEY NAME, used for partition
func (e *KidsCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *KidsCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: KidsDeletedEvent
// ==============================
type KidsDeletedEvent struct {
	KidsDeleted struct{}  `json:"kids_deleted"` // marking
	Id          string    `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *KidsDeletedEvent) Topic() string {
	return MasterKidsTopic
}

// GET KEY NAME, used for partition
func (e *KidsDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *KidsDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: KidsUpdatedEvent
// ==============================
type KidsUpdatedEvent struct {
	KidsUpdated struct{}  `json:"kids_updated"` // marking
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *KidsUpdatedEvent) Topic() string {
	return MasterKidsTopic
}

// GET KEY NAME, used for partition
func (e *KidsUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *KidsUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
