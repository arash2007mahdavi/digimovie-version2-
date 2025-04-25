package server

import (
	"digimovie/src/config"
	"digimovie/src/logging"
	"digimovie/src/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)
var log = logging.NewLogger()

func InitServer(cfg *config.Config) {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())

	digimovie := engine.Group("/digimovie")
	{
		welcome := digimovie.Group("/welcome")
		routers.DigimovieRouter(welcome)
	}

	engine.Run(fmt.Sprintf(":%v", cfg.Server.Port))
}