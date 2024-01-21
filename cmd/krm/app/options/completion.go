package options

type completedOptions struct {
}

type CompletedOptions struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedOptions
}

func (opts *ServerRunOptions) Complete() (CompletedOptions, error) {
	if opts == nil {
		return CompletedOptions{
			completedOptions: &completedOptions{},
		}, nil
	}
	completed := completedOptions{}
	return CompletedOptions{
		completedOptions: &completed,
	}, nil
}
