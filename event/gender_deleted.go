package event

import (
	"time"
)

// Topic
const (
	GenderDeletedTopic = "master.gender"
)

type GenderDeletedEvent struct {
	GenderDeleted string    `json:"gender_deleted"` // marking
	Id            string    `json:"id"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *GenderDeletedEvent) Topic() string {
	return GenderDeletedTopic
}

// GET KEY NAME, used for partition
func (e *GenderDeletedEvent) Key() string {
	return e.Id
}
