package options

import (
	genericoptions "github.com/costa92/krm/pkg/server/options"
	"github.com/spf13/pflag"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions
	SecureServing           *genericoptions.SecureServingOptions
}

// NewOptions returns a new Options object with default parameters
func NewOptions() *Options {
	s := Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
	}
	return &s
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (s *Options) Complete() (CompletedOptions, error) {
	var completed completedOptions
	completed.Options = *s
	return CompletedOptions{&completed}, nil
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet
func (s *Options) AddFlags(fs *pflag.FlagSet) {
	s.GenericServerRunOptions.AddFlags(fs)
	s.SecureServing.AddFlags(fs)
}

// Validate checks that the options are valid
func (s *Options) Validate() []error {
	var errs []error
	errs = append(errs, s.GenericServerRunOptions.Validate()...)
	errs = append(errs, s.SecureServing.Validate()...)
	return errs
}
