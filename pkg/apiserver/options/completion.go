package options

// completedServerRunOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
type completedOptions struct {
	Options
}

// CompletedOptions is a wrapper around completedOptions that exposes the Options.
type CompletedOptions struct {
	*completedOptions
}
