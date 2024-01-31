package apiserver

import (
	"github.com/costa92/krm/pkg/server/options"
	"github.com/gin-gonic/gin"
)

func NewConfig(opts *options.ServerRunOptions) *GenericAPIServer {
	gin.SetMode(opts.Mode)
	res := &GenericAPIServer{
		healthz: opts.Healthz,
		Engine:  gin.New(),
	}
	initGenericAPIServer(res)
	return res
}
