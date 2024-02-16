package api

import (
	"github.com/costa92/krm/internal/api/config"
	genericserver "github.com/costa92/krm/pkg/server"
)

type apiServer struct {
	genericAPIServer *genericserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

// Run runs the api server

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}
	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	// api service 实例化
	server := &apiServer{
		genericAPIServer: genericServer,
	}

	return server, nil
}

func (a *apiServer) PrepareRun() preparedAPIServer {
	initRouter(a.genericAPIServer.Engine)
	return preparedAPIServer{a}
}

func (a *apiServer) Run() error {
	return a.genericAPIServer.Run()
}

// Path: internal/api/config/config.go
func buildGenericConfig(cfg *config.Config) (genericConfig *genericserver.Config, lastErr error) {
	genericConfig = genericserver.NewConfig()
	// Apply the insecure serving options to the generic server
	// 注意 是 cfg 的值赋值到 genericConfig
	if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.ServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.Feature.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return genericConfig, nil
}
