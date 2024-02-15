package app

import (
	"github.com/costa92/krm/cmd/krm/app/options"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start a krm server",
		Long:  "Start a krm server",
		RunE: func(cmd *cobra.Command, args []string) error {
			completedServerRunOption, err := Complete(s)
			if err != nil {
				return err
			}
			return Run(completedServerRunOption)
		},
	}
	return cmd
}
