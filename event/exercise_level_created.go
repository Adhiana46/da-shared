package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	ExerciseCreatedTopic = "master.exercise"
)

type ExerciseCreatedEvent struct {
	ExerciseCreated struct{}  `json:"exercise_created"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	IsDeleted       bool      `json:"is_deleted"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ExerciseCreatedEvent) Topic() string {
	return ExerciseCreatedTopic
}

// GET KEY NAME, used for partition
func (e *ExerciseCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *ExerciseCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
