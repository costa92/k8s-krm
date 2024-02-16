package api

import "github.com/costa92/krm/internal/api/config"

func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}
	return server.PrepareRun().Run()
}
