package server

import (
	"digimovie/src/config"
	"digimovie/src/logging"
	"digimovie/src/routers"
	"digimovie/src/validations"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)
var log = logging.NewLogger()

func InitServer(cfg *config.Config) {
	engine := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobileNumber", validations.ValidateMobileNumber, true)
		val.RegisterValidation("password", validations.ValidatePassword, true)
	}

	engine.Use(gin.Recovery(), gin.Logger())

	digimovie := engine.Group("/digimovie")
	{
		welcome := digimovie.Group("/welcome")
		routers.DigimovieRouter(welcome)
		user := digimovie.Group("/user")
		routers.UserRouter(user)
	}


	engine.Run(fmt.Sprintf(":%v", cfg.Server.Port))
}