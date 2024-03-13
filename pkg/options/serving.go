package options

import (
	"fmt"
	"net"
	"strconv"

	"github.com/costa92/krm/pkg/server"
	"github.com/spf13/pflag"
)

type SecureServingOptions struct {
	// BindAddress is the IP address on which to listen for the --secure-port port.
	BindAddress string `json:"bind-address" mapstructure:"bind-address" yaml:"bind-address"`
	// BindPort is the port on which to listen for secure HTTPS connections.
	BindPort int `json:"bind-port" mapstructure:"bind-port" yaml:"bind-port"`
}

func DefaultServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindAddress: "0.0.0.0",
		BindPort:    8000,
	}
}

func NewSecureServingOptions() *SecureServingOptions {
	secureServingOptions := DefaultServingOptions()
	return secureServingOptions
}

func (s *SecureServingOptions) Address() string {
	return net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort))
}

// Validate checks the SecureServingOptions and returns a slice of errors if there are any.
func (s *SecureServingOptions) Validate() []error {
	var errors []error
	if s.BindPort == 0 {
		if s.BindPort < 0 || s.BindPort > 65535 {
			errors = append(
				errors,
				fmt.Errorf(
					"--insecure.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port",
					s.BindPort,
				),
			)
		}
	}
	return errors
}

func (s *SecureServingOptions) ApplyTo(cfg *server.Config) error {
	cfg.SecureServing = &server.SecureServingInfo{
		Address: s.Address(),
	}
	return nil
}

func (s *SecureServingOptions) Complete() error {
	if s == nil || s.BindPort == 0 {
		return nil
	}
	return nil
}

func (s *SecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	if s == nil {
		return
	}
	fs.StringVar(&s.BindAddress, "bind-address", s.BindAddress, "The IP address on which to serve the secure server")
	fs.IntVar(&s.BindPort, "bind-port", s.BindPort, "The port on which to serve the secure server")
}
