package options

import (
	"net"
	"strconv"
)

type SecureServingOptions struct {
	// BindAddress is the IP address on which to listen for the --secure-port port.
	BindAddress net.IP
	// BindPort is the port on which to listen for secure HTTPS connections.
	BindPort int
}

func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindAddress: net.ParseIP("0.0.0.0"),
		BindPort:    8000,
	}
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
