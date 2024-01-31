package options

import "time"

type ServerRunOptions struct {
	RequestTimeout time.Duration
	Healthz        bool
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		RequestTimeout: 10 * time.Second,
		Healthz:        true,
	}
}

func (s *ServerRunOptions) Complete() {
	return
}
