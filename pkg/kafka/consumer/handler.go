package consumer

type ConsumerHandler interface {
	Handle(payload []byte) error
}
