package auth

import "github.com/gin-gonic/gin"

func AuthRouter(router *gin.RouterGroup, userHandler *UserHandler, middleware *Middleware) {

	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

}
