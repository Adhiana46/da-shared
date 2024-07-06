package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	MasterExerciseTopic = "master.exercise"
)

// ==============================
// Event: ExerciseCreatedEvent
// ==============================
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
	return MasterExerciseTopic
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

// ==============================
// Event: ExerciseDeletedEvent
// ==============================
type ExerciseDeletedEvent struct {
	ExerciseDeleted struct{}  `json:"exercise_deleted"` // marking
	Id              string    `json:"id"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ExerciseDeletedEvent) Topic() string {
	return MasterExerciseTopic
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

// ==============================
// Event: ExerciseUpdatedEvent
// ==============================
type ExerciseUpdatedEvent struct {
	ExerciseUpdated struct{}  `json:"exercise_updated"` // marking
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *ExerciseUpdatedEvent) Topic() string {
	return MasterExerciseTopic
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
