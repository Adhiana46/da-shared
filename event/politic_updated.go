package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	PoliticUpdatedTopic = "master.politic"
)

type PoliticUpdatedEvent struct {
	PoliticUpdated struct{}  `json:"politic_updated"` // marking
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *PoliticUpdatedEvent) Topic() string {
	return PoliticUpdatedTopic
}

// GET KEY NAME, used for partition
func (e *PoliticUpdatedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *PoliticUpdatedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
