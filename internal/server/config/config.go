package config

import (
	"fmt"
)

type (
	Option func(cfg Config) Config
)

type Config struct {
	NatsURL string
}

func NewConfig(opts ...Option) Config {
	config := Config{}

	for _, opt := range opts {
		config = opt(config)
	}

	return config
}

func (c Config) IsValid() (bool, error) {
	if c.NatsURL == "" {
		return false, missingConfigField("NATS url")
	}

	return true, nil
}

func NatsAddress(addr string) Option {
	return func(cfg Config) Config {
		cfg.NatsURL = addr

		return cfg
	}
}

func missingConfigField(name string) error {
	return fmt.Errorf("configuration field '%s' is required", name)
}
