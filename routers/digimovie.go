package routers

import "github.com/gin-gonic/gin"

func DigimovieRouter(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {c.JSON(200, "hello")})
}