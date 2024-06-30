package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	ExerciseDeletedTopic = "master.exercise"
)

type ExerciseDeletedEvent struct {
	ExerciseDeleted struct{}  `json:"exercise_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ExerciseDeletedEvent) Topic() string {
	return ExerciseDeletedTopic
}

// GET KEY NAME, used for partition
func (e *ExerciseDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ExerciseDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
