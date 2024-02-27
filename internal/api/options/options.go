package options

import (
	"encoding/json"

	"github.com/costa92/krm/pkg/cli/cliflags"
	"github.com/costa92/krm/pkg/options"
	"github.com/costa92/krm/pkg/server"
	"github.com/costa92/logger"
)

type RunOptions struct {
	InsecureServing  *options.SecureServingOptions `json:"insecure" mapstructure:"insecure" yaml:"insecure"`
	ServerRunOptions *options.ServerRunOptions     `json:"server"  mapstructure:"server" yaml:"server"`
	Feature          *options.FeatureOptions       `json:"feature" mapstructure:"feature" yaml:"feature"`
	Log              *logger.Options               `json:"log" mapstructure:"log" yaml:"log"`
}

// NewRunOptions returns a new RunOptions
func NewRunOptions() *RunOptions {
	return &RunOptions{
		InsecureServing:  options.NewSecureServingOptions(),
		ServerRunOptions: options.NewServerRunOptions(),
		Feature:          options.NewFeatureOptions(),
		Log:              logger.NewOptions(),
	}
}

// Validate checks the options and returns a slice of errors if any is invalid
func (r *RunOptions) Validate() []error {
	var errs []error
	errs = append(errs, r.InsecureServing.Validate()...)
	errs = append(errs, r.ServerRunOptions.Validate()...)
	errs = append(errs, r.Feature.Validate()...)
	errs = append(errs, r.Log.Validate()...)
	return errs
}

func (r *RunOptions) Flags() (fss cliflags.NamedFlagSets) {
	r.InsecureServing.AddFlags(fss.FlagSet("insecure"))
	r.ServerRunOptions.AddFlags(fss.FlagSet("server"))
	r.Feature.AddFlags(fss.FlagSet("feature"))
	r.Log.AddFlags(fss.FlagSet("log"))
	return fss
}

func (r *RunOptions) ApplyTo(_ *server.Config) error {
	return nil
}

func (r *RunOptions) String() string {
	data, _ := json.Marshal(r)
	return string(data)
}

func (r *RunOptions) Complete() error {
	return r.InsecureServing.Complete()
}
