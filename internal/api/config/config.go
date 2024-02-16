package config

import "github.com/costa92/krm/internal/api/options"

type Config struct {
	*options.RunOptions
}

func CreateConfigFromOptions(opts *options.RunOptions) (*Config, error) {
	return &Config{opts}, nil
}
