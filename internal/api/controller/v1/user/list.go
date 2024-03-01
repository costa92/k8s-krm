package user

import (
	"github.com/costa92/logger"
	"github.com/gin-gonic/gin"
)

func (s *User) List(ctx *gin.Context) {
	logger.L(ctx).Errorw("user list error", "err", "user list error")
	logger.Infow("user list", "user", "user list")
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
