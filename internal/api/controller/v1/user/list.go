package user

import (
	"github.com/costa92/logger"
	"github.com/gin-gonic/gin"
)

func (s *User) List(ctx *gin.Context) {
	logger.Errorw("user list error", "error", "user list error")
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
