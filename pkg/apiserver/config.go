package apiserver

import (
	"github.com/costa92/krm/pkg/server/options"
	"github.com/gin-gonic/gin"
)

func NewConfig(opts *options.ServerRunOptions) *GenericAPIServer {
	res := &GenericAPIServer{
		healthz: opts.Healthz,
		Engine:  gin.Default(),
	}
	initGenericAPIServer(res)
	return res
}
