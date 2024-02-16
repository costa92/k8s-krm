package server

import (
	"log"
	"net/http"
	"time"

	"github.com/costa92/krm/pkg/version"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"golang.org/x/sync/errgroup"
)

type GenericAPIServer struct {
	middlewares       []string
	SecureServingInfo *SecureServingInfo
	*gin.Engine
	healthz         bool
	enableMetrics   bool
	enableProfiling bool
	insecureServer  *http.Server
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

	// open prometheus metrics
	if s.enableMetrics {
		prometheus := ginprometheus.NewPrometheus("gin")
		prometheus.Use(s.Engine)
	}

	// open pprof http://ip:port/debug/pprof
	if s.enableProfiling {
		pprof.Register(s.Engine, "debug/pprof")
	}
}

// Run runs the server
// nolint: revive
func (s *GenericAPIServer) Run() error {
	s.insecureServer = &http.Server{
		Addr:              s.SecureServingInfo.Address,
		Handler:           s,
		ReadHeaderTimeout: 30 * time.Second,
	}
	var eg errgroup.Group
	eg.Go(func() error {
		if err := s.insecureServer.ListenAndServe(); err != nil {
			return err
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}
	return nil
}
