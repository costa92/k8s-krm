package options

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
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

func (s *ServerRunOptions) Validate() []error {
	var errs []error
	return errs
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	if s == nil {
		return
	}
	fs.DurationVar(&s.RequestTimeout, "request-timeout", s.RequestTimeout, "The default request timeout")
	fs.BoolVar(&s.Healthz, "healthz", s.Healthz, "Enable healthz endpoint")
	fs.StringVar(&s.Mode, "mode", s.Mode, "Set gin mode")
	fs.BoolVar(&s.EnableMetrics, "enable-metrics", s.EnableMetrics, "Enable metrics")
	fs.BoolVar(&s.EnableProfiling, "enable-profiling", s.EnableProfiling, "Enable profiling")
}
