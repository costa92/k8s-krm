package options

import "net"

type SecureServingOptions struct {
	BindPort    int
	BindAddress net.IP
}

func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindPort:    443,
		BindAddress: nil,
	}
}

func (s *SecureServingOptions) BindAddressFor() net.IP {
	if s.BindAddress != nil {
		return s.BindAddress
	}
	return net.ParseIP("0.0.0.0")
}
