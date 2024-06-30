package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	PoliticCreatedTopic = "master.politic"
)

type PoliticCreatedEvent struct {
	PoliticCreated struct{}  `json:"politic_created"` // marking
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	IsDeleted      bool      `json:"is_deleted"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *PoliticCreatedEvent) Topic() string {
	return PoliticCreatedTopic
}

// GET KEY NAME, used for partition
func (e *PoliticCreatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *PoliticCreatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
