package server

import (
	"time"
)

type Config struct {
	SecureServing *SecureServingInfo

	// RequestTimeout is the default timeout for requests to the server.
	RequestTimeout time.Duration // 请求时间
}

type SecureServingInfo struct {
	Addr string
	Port int
}

func NewConfig() *Config {
	return &Config{
		SecureServing: &SecureServingInfo{
			Addr: "",
		},
	}
}
