package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterPoliticTopic = "master.politic"
)

// ==============================
// Event: PoliticCreatedEvent
// ==============================
type PoliticCreatedEvent struct {
	PoliticCreated struct{}  `json:"politic_created"` // marking
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	IsDeleted      bool      `json:"is_deleted"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *PoliticCreatedEvent) Topic() string {
	return MasterPoliticTopic
}

// GET KEY NAME, used for partition
func (e *PoliticCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *PoliticCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: PoliticDeletedEvent
// ==============================
type PoliticDeletedEvent struct {
	PoliticDeleted struct{}  `json:"politic_deleted"` // marking
	Id             string    `json:"id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *PoliticDeletedEvent) Topic() string {
	return MasterPoliticTopic
}

// GET KEY NAME, used for partition
func (e *PoliticDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *PoliticDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: PoliticUpdatedEvent
// ==============================
type PoliticUpdatedEvent struct {
	PoliticUpdated struct{}  `json:"politic_updated"` // marking
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *PoliticUpdatedEvent) Topic() string {
	return MasterPoliticTopic
}

// GET KEY NAME, used for partition
func (e *PoliticUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *PoliticUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
