package routers

import (
	"digimovie/src/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DigimovieRouter(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {c.JSON(http.StatusOK, responses.GenerateNormalResponse(true, http.StatusOK, "welcome to digimovie"))})
}