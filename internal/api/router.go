package api

import (
	"github.com/costa92/krm/internal/api/controller/v1/user"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installControllerV1(g)
}

func installMiddleware(g *gin.Engine) {
}

func installControllerV1(g *gin.Engine) *gin.Engine {
	v1 := g.Group("/v1")
	{
		userrv := v1.Group("/user")
		{
			userController := user.NewUser()
			userrv.GET("/list", userController.List)
		}
	}
	return g
}
