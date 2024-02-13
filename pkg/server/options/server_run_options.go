package options

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ServerRunOptions struct {
	RequestTimeout  time.Duration
	Healthz         bool
	Mode            string
	EnableMetrics   bool
	EnableProfiling bool
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		RequestTimeout:  10 * time.Second,
		Healthz:         true,
		Mode:            gin.ReleaseMode,
		EnableMetrics:   true,
		EnableProfiling: true,
	}
}

func (s *ServerRunOptions) Complete() {
	return
}
