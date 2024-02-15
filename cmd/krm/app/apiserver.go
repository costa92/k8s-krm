package app

import (
	appoptions "github.com/costa92/krm/cmd/krm/app/options"
	"github.com/costa92/krm/pkg/apiserver"
	"github.com/costa92/krm/pkg/apiserver/options"
	controlplaneapiserver "github.com/costa92/krm/pkg/apiserver/options"
)

type apiServer struct {
	opts             controlplaneapiserver.CompletedOptions
	genericAPIServer apiserver.GenericAPIServer
	config           *apiserver.Config
}

type preparedAPIServer struct {
	*apiServer
}

func createAPIServer(opts completedServerRunOptions) *apiServer {
	cfg, gs := apiserver.NewConfig(opts.CompletedOpts)
	return &apiServer{
		genericAPIServer: *gs,
		opts:             opts.CompletedOpts,
		config:           cfg,
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

func Run(s completedServerRunOptions) error {
	if err := createAPIServer(s).prepareRun().Run(); err != nil {
		return err
	}
	return nil
}

type completedServerRunOptions struct {
	CompletedOpts options.CompletedOptions
}

// Complete completes all the required options.
func Complete(s *appoptions.ServerRunOptions) (completedServerRunOptions, error) {
	var opts completedServerRunOptions
	completedOpts, err := s.Complete()
	if err != nil {
		return opts, err
	}
	opts.CompletedOpts = completedOpts
	return opts, err
}
