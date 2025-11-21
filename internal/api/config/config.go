package config

import (
	"fmt"
)

type (
	Option func(cfg Config) Config
)

type Config struct {
	Addr string
}

func NewConfig(opts ...Option) Config {
	config := Config{}

	for _, opt := range opts {
		config = opt(config)
	}

	return config
}

func (c Config) IsValid() (bool, error) {
	if c.Addr == "" {
		return false, missingConfigField("Address")
	}

	return true, nil
}

func Address(addr string) Option {
	return func(cfg Config) Config {
		cfg.Addr = addr

		return cfg
	}
}

func missingConfigField(name string) error {
	return fmt.Errorf("configuration field '%s' is required", name)
}
