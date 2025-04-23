package server

import (
	"digimovie/src/config"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *config.Config) {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
}