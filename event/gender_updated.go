package event

import (
	"time"
)

// Topic
const (
	GenderUpdatedTopic = "master.gender"
)

type GenderUpdatedEvent struct {
	GenderUpdated string    `json:"gender_updated"` // marking
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *GenderUpdatedEvent) Topic() string {
	return GenderUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *GenderUpdatedEvent) Key() string {
	return e.Id
}
