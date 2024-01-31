package options

import (
	"github.com/gin-gonic/gin"
	"time"
)

type ServerRunOptions struct {
	RequestTimeout time.Duration
	Healthz        bool
	Mode           string
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		RequestTimeout: 10 * time.Second,
		Healthz:        true,
		Mode:           gin.ReleaseMode,
	}
}

func (s *ServerRunOptions) Complete() {
	return
}
