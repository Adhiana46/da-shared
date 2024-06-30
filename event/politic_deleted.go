package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	PoliticDeletedTopic = "master.politic"
)

type PoliticDeletedEvent struct {
	PoliticDeleted struct{}  `json:"politic_deleted"` // marking
	Id             string    `json:"id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *PoliticDeletedEvent) Topic() string {
	return PoliticDeletedTopic
}

// GET KEY NAME, used for partition
func (e *PoliticDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *PoliticDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
