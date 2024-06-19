package event

import (
	"time"
)

type UserSignInEventAction string

// Topic
const (
	UserSigninTopic = "user.signin"
)

// Action
const (
	UserSignInEventActionSignin       UserSignInEventAction = "signin"
	UserSignInEventActionRefreshToken UserSignInEventAction = "refresh-token"
)

type UserSignInEvent struct {
	Id           string                `json:"id"`
	Email        string                `json:"email"`
	Package      string                `json:"package"`
	LastActiveAt time.Time             `json:"last_active_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	IpAddr       string                `json:"ip_address"`
	UserAgent    string                `json:"user_agent"`
	Action       UserSignInEventAction `json:"action"`
}

// GET TOPIC NAME
func (e *UserSignInEvent) Topic() string {
	return UserSigninTopic
}

// GET KEY NAME, used for partition
func (e *UserSignInEvent) Key() string {
	return e.Id
}
