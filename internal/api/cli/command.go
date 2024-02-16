package cli

import (
	"github.com/costa92/krm/pkg/cli"
	"github.com/spf13/cobra"
)

func NewCommand(name string, buildCommand func(string) error) *cobra.Command {
	cmd := cli.BuildCommand(name, buildCommand)
	cmd.Short = "api service for krm"
	cmd.Long = "api service for krm"
	return cmd
}
