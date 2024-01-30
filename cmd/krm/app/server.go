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
			completedOption, err := Complete(s)
			if err != nil {
				return err
			}
			return Run(completedOption)
		},
	}
	return cmd
}

// completedServerRunOptions
type completedServerRunOptions struct {
	*options.CompletedOptions
}

func Complete(s *options.ServerRunOptions) (completedServerRunOptions, error) {
	var options completedServerRunOptions
	err := s.SecureServing.ApplyTo()
	if err != nil {
		return options, err
	}
	return options, nil
}

func Run(opts completedServerRunOptions) error {
	return nil
}
