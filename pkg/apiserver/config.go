package apiserver

import (
	apiserveropts "github.com/costa92/krm/pkg/apiserver/options"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Options apiserveropts.CompletedOptions
}

//func NewConfig(opts *options.ServerRunOptions) *GenericAPIServer {
//	gin.SetMode(opts.Mode)
//	res := &GenericAPIServer{
//		healthz:         opts.Healthz,
//		Engine:          gin.New(),
//		enableMetrics:   opts.EnableMetrics,
//		enableProfiling: opts.EnableProfiling,
//	}
//	initGenericAPIServer(res)
//	return res
//}

func NewConfig(opts apiserveropts.CompletedOptions) (*Config, *GenericAPIServer) {
	cfg := &Config{
		Options: opts,
	}
	genericServerRunOptions := opts.GenericServerRunOptions
	gin.SetMode(genericServerRunOptions.Mode)
	gs := &GenericAPIServer{
		healthz:         genericServerRunOptions.Healthz,
		Engine:          gin.New(),
		enableMetrics:   genericServerRunOptions.EnableMetrics,
		enableProfiling: genericServerRunOptions.EnableProfiling,
	}
	return cfg, gs
}

func (s *Config) Complete() error {
	return nil
}
