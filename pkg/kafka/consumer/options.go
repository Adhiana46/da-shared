package consumer

import "time"

type Option func(*Consumer)

func WithBrokers(brokers []string) Option {
	return func(c *Consumer) {
		c.brokers = brokers
	}
}

func WithGroupID(groupID string) Option {
	return func(c *Consumer) {
		c.groupID = groupID
	}
}

func WithTopics(topics []string) Option {
	return func(c *Consumer) {
		c.topics = topics
	}
}

func WithShutdownTimeout(timeout time.Duration) Option {
	return func(c *Consumer) {
		c.shutdownTimeout = timeout
	}
}

func WithKafkaVersion(version string) Option {
	return func(c *Consumer) {
		c.version = version
	}
}
