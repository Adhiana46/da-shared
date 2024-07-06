package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterEducationLevelTopic = "master.education_level"
)

// ==============================
// Event: EducationLevelCreatedEvent
// ==============================
type EducationLevelCreatedEvent struct {
	EducationLevelCreated struct{}  `json:"education_level_created"` // marking
	Id                    string    `json:"id"`
	Name                  string    `json:"name"`
	IsDeleted             bool      `json:"is_deleted"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *EducationLevelCreatedEvent) Topic() string {
	return MasterEducationLevelTopic
}

// GET KEY NAME, used for partition
func (e *EducationLevelCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *EducationLevelCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: EducationLevelDeletedEvent
// ==============================
type EducationLevelDeletedEvent struct {
	EducationLevelDeleted struct{}  `json:"education_level_deleted"` // marking
	Id                    string    `json:"id"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *EducationLevelDeletedEvent) Topic() string {
	return MasterEducationLevelTopic
}

// GET KEY NAME, used for partition
func (e *EducationLevelDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *EducationLevelDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}

// ==============================
// Event: EducationLevelUpdatedEvent
// ==============================
type EducationLevelUpdatedEvent struct {
	EducationLevelUpdated struct{}  `json:"education_level_updated"` // marking
	Id                    string    `json:"id"`
	Name                  string    `json:"name"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *EducationLevelUpdatedEvent) Topic() string {
	return MasterEducationLevelTopic
}

// GET KEY NAME, used for partition
func (e *EducationLevelUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *EducationLevelUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
