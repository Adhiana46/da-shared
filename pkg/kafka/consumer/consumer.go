package consumer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/IBM/sarama"
)

var (
	ErrNoBrokers        = errors.New("no Kafka bootstrap brokers defined")
	ErrUndefinedGroupID = errors.New("group id not defined")
	ErrNoTopics         = errors.New("no Kafka topics defined")
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

type Consumer struct {
	consumerGroup   sarama.ConsumerGroup
	notify          chan error
	shutdownTimeout time.Duration
	groupID         string
	brokers         []string
	topics          []string
	version         string
	ready           chan bool
	handlers        map[string]ConsumerHandler
}

func New(options ...Option) (*Consumer, error) {
	c := Consumer{
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		version:         sarama.DefaultVersion.String(),
		ready:           make(chan bool),
		handlers:        map[string]ConsumerHandler{},
	}

	// apply options
	for _, opt := range options {
		opt(&c)
	}

	// Validate configuration
	if len(c.brokers) == 0 {
		return nil, ErrNoBrokers
	}
	if c.groupID == "" {
		return nil, ErrUndefinedGroupID
	}

	// parse sarama version
	saramaVersion, err := sarama.ParseKafkaVersion(c.version)
	if err != nil {
		return nil, err
	}

	// config sarama
	config := sarama.NewConfig()
	config.Version = saramaVersion
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create sarama consumer group
	consumerGroup, err := sarama.NewConsumerGroup(c.brokers, c.groupID, config)
	if err != nil {
		return nil, err
	}

	c.consumerGroup = consumerGroup

	return &c, nil
}

func (c *Consumer) Start() {
	if len(c.topics) == 0 {
		c.notify <- ErrNoTopics
		return
	}

	slog.Info(fmt.Sprintf("Starting consumer for topics %v", c.topics))
	ctx := context.Background()
	go func() {
		for {
			if err := c.consumerGroup.Consume(ctx, c.topics, c); err != nil {
				c.notify <- err
				return
			}
			if ctx.Err() != nil {
				return
			}
			c.ready = make(chan bool)
		}
	}()
	<-c.ready // Await till the consumer has been set up
}

func (c *Consumer) Notify() <-chan error {
	return c.notify
}

func (c *Consumer) Shutdown() error {

	if err := c.consumerGroup.Close(); err != nil {
		return err
	}
	return nil
}

// sarama.ConsumerGroupHandler
func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

// sarama.ConsumerGroupHandler
func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) RegisterHandler(topic string, h ConsumerHandler) {
	_, exists := c.handlers[topic]

	if !exists {
		c.topics = append(c.topics, topic)
	}

	c.handlers[topic] = h
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				slog.Info("message channel was closed")
				return nil
			}

			if err := c.processMessage(session, msg); err != nil {
				slog.Error("Failed to process message", slog.Any("details", map[string]any{
					"partition": msg.Partition,
					"topic":     msg.Topic,
					"offset":    msg.Offset,
					"payload":   string(msg.Value),
					"error":     err.Error(),
				}))
			}
		case <-session.Context().Done():
			return nil
		}
	}
}

func (c *Consumer) processMessage(session sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage) error {
	h, ok := c.handlers[msg.Topic]
	if !ok {
		return fmt.Errorf("no consumer handler available for topic = %s", msg.Topic)
	}

	if err := h.Handle(msg.Value); err != nil {
		return err
	}

	// Mark messsage as processed
	session.MarkMessage(msg, "")

	return nil
}
