package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	EducationLevelDeletedTopic = "master.education_level"
)

type EducationLevelDeletedEvent struct {
	EducationLevelDeleted struct{}  `json:"education_level_deleted"` // marking
	Id                    string    `json:"id"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *EducationLevelDeletedEvent) Topic() string {
	return EducationLevelDeletedTopic
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
