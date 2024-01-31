package apiserver

import (
	"github.com/costa92/krm/cmd/krm/app/options"
	"github.com/costa92/krm/pkg/version"
	"github.com/gin-gonic/gin"
	"log"
)

type GenericAPIServer struct {
	middlewares []string

	*gin.Engine

	healthz       bool
	enableMetrics bool
}

func initGenericAPIServer(s *GenericAPIServer) {
	s.Setup()
	s.InstallAPIs()
}

func (s *GenericAPIServer) Setup() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

func (s *GenericAPIServer) InstallAPIs() {
	if s.healthz {
		s.Engine.GET("/healthz", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}
	s.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": version.GetVersion(),
		})
	})
}

func (s *GenericAPIServer) Run(opts options.CompletedOptions) error {
	addr := opts.SecureServingOptions.Address()
	return s.Engine.Run(addr)
}
