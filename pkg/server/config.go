package server

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Mode          string
	Healthz       bool
	Middlewares   []string
	SecureServing *SecureServingInfo
	// RequestTimeout is the default timeout for requests to the server.
	RequestTimeout time.Duration // 请求时间

	// EnableProfiling specifies whether to enable profiling via web interface host:port/debug/pprof/
	EnableProfiling bool
	EnableMetrics   bool
}

type SecureServingInfo struct {
	Address string
}

// NewConfig returns a new Config with default values.
func NewConfig() *Config {
	return &Config{
		Mode:            gin.ReleaseMode,
		Healthz:         true,
		EnableProfiling: true,
		EnableMetrics:   true,
	}
}

// CompletedConfig is a completed server configuration.
type CompletedConfig struct {
	*Config
}

// Complete fills in any fields not set that are required to have valid data.
// ch: 完成填充任何未设置的字段，这些字段需要有效数据。
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

// New returns a new server from the config.
func (c CompletedConfig) New() (*GenericAPIServer, error) {
	gin.SetMode(c.Mode)
	s := &GenericAPIServer{
		middlewares:       c.Middlewares,
		SecureServingInfo: c.SecureServing,
		Engine:            gin.New(),
		healthz:           c.Healthz,
		enableMetrics:     c.EnableMetrics,
		enableProfiling:   c.EnableProfiling,
	}
	initGenericAPIServer(s)
	return s, nil
}
