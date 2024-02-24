package api

import (
	"github.com/costa92/krm/internal/api/config"
	"github.com/costa92/krm/internal/api/options"
	"github.com/costa92/krm/pkg/app"
	"github.com/costa92/logger"
)

// NewApp creates a new app for the api service
// zh: 创建一个新的 api 服务实例
func NewApp(basename string) *app.App {
	opts := options.NewRunOptions()
	application := app.NewApp("api service for krm",
		basename,
		app.WithOptions(opts),
		app.WithRunFunc(run(opts)),
	)
	return application
}

// Path: internal/api/app.go
func run(opts *options.RunOptions) app.RunFunc {
	return func(basename string) error {
		logger.Init(opts.Log)
		defer logger.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
