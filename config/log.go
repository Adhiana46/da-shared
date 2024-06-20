package config

type LogConfig struct {
	Level string `env-required:"true" env:"LOG_LEVEL" yaml:"level"`
}
