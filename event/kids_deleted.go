package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Topic
const (
	KidsDeletedTopic = "master.kids"
)

type KidsDeletedEvent struct {
	KidsDeleted struct{}  `json:"kids_deleted"` // marking
	Id          string    `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *KidsDeletedEvent) Topic() string {
	return KidsDeletedTopic
}

// GET KEY NAME, used for partition
func (e *KidsDeletedEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *KidsDeletedEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
