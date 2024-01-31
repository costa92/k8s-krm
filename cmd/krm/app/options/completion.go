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
	if opts == nil {
		return CompletedOptions{
			completedOptions: &completedOptions{
				ServerRunOptions: options.NewServerRunOptions(),
			},
		}, nil
	}
	completed := completedOptions{
		ServerRunOptions: options.NewServerRunOptions(),
		SecureServingOptions: &options.SecureServingOptions{
			BindAddress: opts.SecureServing.BindAddress,
			BindPort:    opts.SecureServing.BindPort,
		},
	}
	return CompletedOptions{
		completedOptions: &completed,
	}, nil
}
