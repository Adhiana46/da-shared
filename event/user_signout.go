package event

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

const (
	UserSignoutTopic = "user.signout"
)

type UserSignOutEvent struct {
	Id           string    `json:"id"`
	Email        string    `json:"email"`
	Package      string    `json:"package"`
	LastActiveAt time.Time `json:"last_active_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *UserSignOutEvent) Topic() string {
	return UserSignoutTopic
}

// GET KEY NAME, used for partition
func (e *UserSignOutEvent) Key() string {
	return e.Id
}

// GET Payload
func (e *UserSignOutEvent) Payload() ([]byte, error) {
	payload, err := json.Marshal(e)
	if err != nil {
		return nil, errors.Wrap(err, "json marshall failed")
	}

	return payload, nil
}
