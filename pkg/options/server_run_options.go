package options

import (
	"time"

	"github.com/costa92/krm/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

type ServerRunOptions struct {
	Mode           string        `json:"mode" mapstructure:"mode" yaml:"mode"`
	Healthz        bool          `json:"healthz" mapstructure:"healthz" yaml:"healthz"`
	Middlewares    []string      `json:"middlewares" mapstructure:"middlewares" yaml:"middlewares"`
	RequestTimeout time.Duration `json:"request-timeout" mapstructure:"request-timeout" yaml:"request-timeout"`
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		RequestTimeout: 10 * time.Second,
		Healthz:        true,
		Mode:           gin.ReleaseMode,
	}
}

func (s *ServerRunOptions) Complete() error {
	return nil
}

func (s *ServerRunOptions) Validate() []error {
	var errs []error
	return errs
}

func (s *ServerRunOptions) ApplyTo(cfg *server.Config) error {
	cfg.Mode = s.Mode
	cfg.Healthz = s.Healthz
	cfg.RequestTimeout = s.RequestTimeout
	cfg.Middlewares = s.Middlewares
	return nil
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	if s == nil {
		return
	}
	fs.DurationVar(&s.RequestTimeout, "request-timeout", s.RequestTimeout, "The default request timeout")
	fs.BoolVar(&s.Healthz, "healthz", s.Healthz, "Enable healthz endpoint")
	fs.StringVar(&s.Mode, "mode", s.Mode, "Set gin mode")

	fs.StringSliceVar(&s.Middlewares, "server.middlewares", s.Middlewares, ""+
		"List of allowed middlewares for server, comma separated. If this list is empty default middlewares will be used.")
}
