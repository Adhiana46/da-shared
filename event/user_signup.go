package event

import (
	"time"
)

const (
	UserSignupTopic = "user.signup"
)

type UserSignUpEvent struct {
	Id           string    `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Package      string    `json:"package"`
	LastActiveAt time.Time `json:"last_active_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// GET TOPIC NAME
func (e *UserSignUpEvent) Topic() string {
	return UserSignupTopic
}

// GET KEY NAME, used for partition
func (e *UserSignUpEvent) Key() string {
	return e.Id
}
