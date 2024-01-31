package app

import (
	"github.com/costa92/krm/cmd/krm/app/options"
	"github.com/costa92/krm/pkg/apiserver"
)

type apiServer struct {
	opts             options.CompletedOptions
	genericAPIServer apiserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

func createAPIServer(opts completedServerRunOptions) *apiServer {
	gs := apiserver.NewConfig(opts.CompletedOpts.ServerRunOptions)
	return &apiServer{
		genericAPIServer: *gs,
		opts:             opts.CompletedOpts,
	}
}

func (s *apiServer) prepareRun() preparedAPIServer {
	return preparedAPIServer{
		apiServer: s,
	}
}

func (s *apiServer) Run() error {
	return s.genericAPIServer.Run(s.opts)
}

func Run(opts completedServerRunOptions) error {
	if err := createAPIServer(opts).prepareRun().Run(); err != nil {
		return err
	}
	return nil
}

type completedServerRunOptions struct {
	CompletedOpts options.CompletedOptions
}

// Complete completes all the required options.
func Complete(s *options.ServerRunOptions) (completedServerRunOptions, error) {
	var opts completedServerRunOptions
	completedOpts, err := s.Complete()
	if err != nil {
		return opts, err
	}
	opts.CompletedOpts = completedOpts
	return opts, err
}
