package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterZodiacTopic = "master.zodiac"
)

// ==============================
// Event: ZodiacCreatedEvent
// ==============================
type ZodiacCreatedEvent struct {
	ZodiacCreated struct{}  `json:"zodiac_created"` // marking
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	IsDeleted     bool      `json:"is_deleted"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ZodiacCreatedEvent) Topic() string {
	return MasterZodiacTopic
}

// GET KEY NAME, used for partition
func (e *ZodiacCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ZodiacCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: ZodiacDeletedEvent
// ==============================
type ZodiacDeletedEvent struct {
	ZodiacDeleted struct{}  `json:"zodiac_deleted"` // marking
	Id            string    `json:"id"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ZodiacDeletedEvent) Topic() string {
	return MasterZodiacTopic
}

// GET KEY NAME, used for partition
func (e *ZodiacDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ZodiacDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: ZodiacUpdatedEvent
// ==============================
type ZodiacUpdatedEvent struct {
	ZodiacUpdated struct{}  `json:"zodiac_updated"` // marking
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ZodiacUpdatedEvent) Topic() string {
	return MasterZodiacTopic
}

// GET KEY NAME, used for partition
func (e *ZodiacUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ZodiacUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
