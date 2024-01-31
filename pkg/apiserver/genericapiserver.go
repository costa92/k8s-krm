package apiserver

import (
	"github.com/gin-gonic/gin"
)

type GenericAPIServer struct {
	middlewares []string

	*gin.Engine

	healthz       bool
	enableMetrics bool
}

func NewGenericAPIServer() *GenericAPIServer {
	return &GenericAPIServer{
		Engine: gin.New(),
	}
}

func (s *GenericAPIServer) Run(addr string) error {
	return s.Engine.Run(addr)
}
