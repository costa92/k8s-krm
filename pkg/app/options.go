package app

import "github.com/costa92/krm/pkg/cli/cliflags"

// Path: pkg/server/options.go
// Compare this snippet from internal/api/options/options.go:
// package options

type CliOptions interface {
	Flags() (fs cliflags.NamedFlagSets)
	Validate() []error
}

type CompletedOptions interface {
	ApplyFlags() []error
}

type CompleteableOptions interface {
	Complete() error
}

type PrintableOptions interface {
	String() string
}
