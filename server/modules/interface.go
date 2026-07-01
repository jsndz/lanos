package modules

import "github.com/gin-gonic/gin"

type Module interface {
	Name() string
	RegisterRoutes(router *gin.RouterGroup)
}
