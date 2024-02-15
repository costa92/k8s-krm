package apiserver

import (
	options2 "github.com/costa92/krm/pkg/apiserver/options"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/costa92/krm/pkg/version"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type GenericAPIServer struct {
	middlewares []string

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
	// open pprof http://ip:port/debug/pprof
	if s.enableProfiling {
		pprof.Register(s.Engine, "debug/pprof")
	}
}

// Run runs the server
func (s *GenericAPIServer) Run(opts options2.CompletedOptions) error {
	addr := opts.SecureServing.Address()
	s.insecureServer = &http.Server{
		Addr:    addr,
		Handler: s,
	}

	var eg errgroup.Group
	eg.Go(func() error {
		if err := http.ListenAndServe(addr, s); err != nil {
			return err
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}
	return nil
}
