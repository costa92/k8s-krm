package apiserver

import (
	"github.com/costa92/krm/pkg/server/options"
	"github.com/gin-gonic/gin"
)

func NewConfig(opts *options.ServerRunOptions) *GenericAPIServer {
	gin.SetMode(opts.Mode)
	res := &GenericAPIServer{
		healthz:         opts.Healthz,
		Engine:          gin.New(),
		enableMetrics:   opts.EnableMetrics,
		enableProfiling: opts.EnableProfiling,
	}
	initGenericAPIServer(res)
	return res
}
