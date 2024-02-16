package user

import "github.com/gin-gonic/gin"

func (s *User) List(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
