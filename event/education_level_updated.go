package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	EducationLevelUpdatedTopic = "master.education_level"
)

type EducationLevelUpdatedEvent struct {
	EducationLevelUpdated struct{}  `json:"education_level_updated"` // marking
	Id                    string    `json:"id"`
	Name                  string    `json:"name"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *EducationLevelUpdatedEvent) Topic() string {
	return EducationLevelUpdatedTopic
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
