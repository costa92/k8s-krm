package options

import (
	controlplaneapiserver "github.com/costa92/krm/pkg/apiserver/options"
	"github.com/spf13/pflag"
)

type ServerRunOptions struct {
	*controlplaneapiserver.Options
}

func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{
		Options: controlplaneapiserver.NewOptions(),
	}
	return &s
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	s.GenericServerRunOptions.AddFlags(fs)
	s.SecureServing.AddFlags(fs)
}
