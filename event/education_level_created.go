package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	EducationLevelCreatedTopic = "master.education_level"
)

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
	return EducationLevelCreatedTopic
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
