package event

import (
	"testing"
)

func TestUserSignOutEvent_Topic(t *testing.T) {
	tests := []struct {
		name string
		e    *UserSignOutEvent
		want string
	}{
		{
			name: "UserSignOutEvent.Topic() should return user.signout",
			e:    &UserSignOutEvent{},
			want: "user.signout",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Topic(); got != tt.want {
				t.Errorf("UserSignOutEvent.Topic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSignOutEvent_Key(t *testing.T) {
	tests := []struct {
		name string
		e    *UserSignOutEvent
		want string
	}{
		{
			name: "UserSignOutEvent.Key() should return event id",
			e: &UserSignOutEvent{
				Id: "some-random-generated-id",
			},
			want: "some-random-generated-id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Key(); got != tt.want {
				t.Errorf("UserSignOutEvent.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
