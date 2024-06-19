package event

import (
	"testing"
)

func TestUserSignInEvent_Topic(t *testing.T) {
	tests := []struct {
		name string
		e    *UserSignInEvent
		want string
	}{
		{
			name: "UserSignInEvent.Topic() should return user.signin",
			e:    &UserSignInEvent{},
			want: "user.signin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Topic(); got != tt.want {
				t.Errorf("UserSignInEvent.Topic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSignInEvent_Key(t *testing.T) {
	tests := []struct {
		name string
		e    *UserSignInEvent
		want string
	}{
		{
			name: "UserSignInEvent.Key() should return event id",
			e: &UserSignInEvent{
				Id: "some-random-generated-id",
			},
			want: "some-random-generated-id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Key(); got != tt.want {
				t.Errorf("UserSignInEvent.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
