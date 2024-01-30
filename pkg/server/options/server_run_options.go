package options

import "time"

type ServerRunOptions struct {
	RequestTimeout time.Duration
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{}
}

func (s *ServerRunOptions) Complete() {
	return
}
