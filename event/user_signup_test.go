package event

import (
	"testing"
)

func TestUserSignUpEvent_Topic(t *testing.T) {
	tests := []struct {
		name string
		e    *UserSignUpEvent
		want string
	}{
		{
			name: "UserSignUpEvent.Topic() should return user.signup",
			e:    &UserSignUpEvent{},
			want: "user.signup",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Topic(); got != tt.want {
				t.Errorf("UserSignUpEvent.Topic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSignUpEvent_Key(t *testing.T) {
	tests := []struct {
		name string
		e    *UserSignUpEvent
		want string
	}{
		{
			name: "UserSignUpEvent.Key() should return event id",
			e: &UserSignUpEvent{
				Id: "some-random-generated-id",
			},
			want: "some-random-generated-id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Key(); got != tt.want {
				t.Errorf("UserSignUpEvent.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
