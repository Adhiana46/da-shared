package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	ExerciseUpdatedTopic = "master.exercise"
)

type ExerciseUpdatedEvent struct {
	ExerciseUpdated struct{}  `json:"exercise_updated"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ExerciseUpdatedEvent) Topic() string {
	return ExerciseUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *ExerciseUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ExerciseUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
