package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/costa92/logger"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-isatty"
)

// defaultLogFormatter is the default log format function Logger middleware uses.
var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}
	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency -= param.Latency % time.Second
	}
	return fmt.Sprintf("%s%3d%s - [%s] \"%v %s%s%s %s\" %s",
		// param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.ClientIP,
		param.Latency,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

func Logger() gin.HandlerFunc {
	f := gin.LoggerConfig{
		Formatter: GetDefaultLogFormatterWithRequestID(),
		Output:    nil,
		SkipPaths: nil,
	}
	return LoggerWithConfig(f)
}

func LoggerWithConfig(conf gin.LoggerConfig) gin.HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}
	out := conf.Output
	if out == nil {
		out = gin.DefaultWriter
	}

	notlogged := conf.SkipPaths

	isTerm := true
	if w, ok := out.(*os.File); !ok || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd())) {
		isTerm = false
	}
	if isTerm {
		gin.ForceConsoleColor()
	}

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// Process request
		c.Next()
		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := gin.LogFormatterParams{
				Request: c.Request,
				Keys:    c.Keys,
			}
			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)
			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = c.Writer.Size()
			if raw != "" {
				path = path + "?" + raw
			}
			param.Path = path
			logger.L(c).Infof("logger middleware: %s", formatter(param))
		}
	}
}
