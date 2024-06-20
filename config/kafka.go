package config

type KafkaConfig struct {
	ClientId string   `env:"KAFKA_CLIENT_ID" yaml:"client_id"`
	Brokers  []string `env-separator:"," env:"KAFKA_BROKERS" yaml:"brokers"`
}
