package app

import (
	"github.com/costa92/krm/cmd/krm/app/options"
	"github.com/spf13/cobra"
)

func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start a krm server",
		Long:  "Start a krm server",
		RunE: func(cmd *cobra.Command, args []string) error {
			completedOption, err := s.Complete()
			if err != nil {
				return err
			}
			return Run(completedOption)
		},
	}
	return cmd
}

func Run(opts options.CompletedOptions) error {
	return nil
}
