package event

type Event interface {
	Topic() string
	Key() string
	Payload() ([]byte, error)
}
