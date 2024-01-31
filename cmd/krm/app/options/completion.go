package options

import "github.com/costa92/krm/pkg/server/options"

type completedOptions struct {
	ServerRunOptions     *options.ServerRunOptions
	SecureServingOptions *options.SecureServingOptions
}

type CompletedOptions struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedOptions
}

func (opts *ServerRunOptions) Complete() (CompletedOptions, error) {
	completed := completedOptions{
		ServerRunOptions:     options.NewServerRunOptions(),
		SecureServingOptions: options.NewSecureServingOptions(),
	}
	return CompletedOptions{
		completedOptions: &completed,
	}, nil
}
