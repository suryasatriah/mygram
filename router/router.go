package router

import (
	"github.com/gin-gonic/gin"
	"github.com/suryasatriah/mygram/controller"
)

// func StartApp() *gin.Engine {
// 	router := gin.Default()
// 	userRouter := router.Group("/users")
// 	{
// 		userRouter.POST("/register", controller.UserRegister)
// 	}

//		return router
//	}
func PublicRoutes(g *gin.RouterGroup) {

	g.POST("/user/register", controller.UserRegister)
	g.POST("/login", controller.UserLogin)
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.POST("/socialmediaregister", controller.CreateSocialMedia)
	g.GET("/socialmedia", controller.GetSocialMedia)
}
