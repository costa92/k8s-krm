package options

import (
	"github.com/costa92/krm/pkg/server"
	"github.com/spf13/pflag"
)

type FeatureOptions struct {
	EnableProfiling bool `json:"profiling"      mapstructure:"profiling"`
	EnableMetrics   bool `json:"enable-metrics" mapstructure:"enable-metrics"`
}

func NewFeatureOptions() *FeatureOptions {
	defaults := server.NewConfig()
	return &FeatureOptions{
		EnableMetrics:   defaults.EnableMetrics,
		EnableProfiling: defaults.EnableProfiling,
	}
}

func (o *FeatureOptions) ApplyTo(c *server.Config) error {
	c.EnableProfiling = o.EnableProfiling
	c.EnableMetrics = o.EnableMetrics
	return nil
}

func (o *FeatureOptions) Validate() []error {
	return []error{}
}

func (o *FeatureOptions) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&o.EnableProfiling, "enable-profiling", o.EnableProfiling, "Enable profiling via web interface host:port/debug/pprof/")
	fs.BoolVar(&o.EnableMetrics, "enable-metrics", o.EnableMetrics, "Enable metrics")
}

func (o *FeatureOptions) Complete() error {
	return nil
}
