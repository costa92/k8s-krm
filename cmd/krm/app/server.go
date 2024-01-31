package app

import (
	"github.com/costa92/krm/cmd/krm/app/options"
	"github.com/costa92/krm/pkg/apiserver"
	"github.com/spf13/cobra"
)

func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
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

type completedServerRunOptions struct {
	CompletedOpts options.CompletedOptions
	Gs            *apiserver.GenericAPIServer
}

func Complete(s *options.ServerRunOptions) (completedServerRunOptions, error) {
	var opts completedServerRunOptions
	completedOpts, err := s.Complete()
	if err != nil {
		return opts, err
	}
	return completedServerRunOptions{
		CompletedOpts: completedOpts,
		Gs:            apiserver.NewGenericAPIServer(),
	}, err
}

func Run(opts completedServerRunOptions) error {
	opts.Gs.Run(":8000")
	return nil
}
