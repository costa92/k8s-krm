package options

import (
	"encoding/json"

	"github.com/costa92/krm/pkg/cli/cliflags"
	"github.com/costa92/krm/pkg/options"
	"github.com/costa92/krm/pkg/server"
)

type RunOptions struct {
	InsecureServing *options.SecureServingOptions `json:"insecure" mapstructure:"insecure" yaml:"insecure"`
}

// NewRunOptions returns a new RunOptions
func NewRunOptions() *RunOptions {
	return &RunOptions{
		InsecureServing: options.NewSecureServingOptions(),
	}
}

// Validate checks the options and returns a slice of errors if any is invalid
func (r *RunOptions) Validate() []error {
	var errs []error
	errs = append(errs, r.InsecureServing.Validate()...)
	return errs
}

func (r *RunOptions) Flags() (fss cliflags.NamedFlagSets) {
	r.InsecureServing.AddFlags(fss.FlagSet("insecure"))
	return fss
}

func (r *RunOptions) ApplyTo(cfg *server.Config) error {
	return nil
}

func (r *RunOptions) String() string {
	data, _ := json.Marshal(r)
	return string(data)
}
