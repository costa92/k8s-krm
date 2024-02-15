package options

import (
	"net"
	"strconv"

	"github.com/spf13/pflag"
)

type SecureServingOptions struct {
	// BindAddress is the IP address on which to listen for the --secure-port port.
	BindAddress net.IP
	// BindPort is the port on which to listen for secure HTTPS connections.
	BindPort int
}

func DefaultServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindAddress: net.ParseIP("0.0.0.0"),
		BindPort:    8000,
	}
}

func NewSecureServingOptions() *SecureServingOptions {
	secureServingOptions := DefaultServingOptions()
	return secureServingOptions
}

func (s *SecureServingOptions) Address() string {
	return net.JoinHostPort(s.BindAddress.String(), strconv.Itoa(s.BindPort))
}

func (s *SecureServingOptions) Validate() []error {
	return nil
}

func (s *SecureServingOptions) ApplyTo() error {
	return nil
}

func (s *SecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	if s == nil {
		return
	}
	fs.IPVar(&s.BindAddress, "bind-address", s.BindAddress, "The IP address on which to serve the secure server")
	fs.IntVar(&s.BindPort, "bind-port", s.BindPort, "The port on which to serve the secure server")
}
