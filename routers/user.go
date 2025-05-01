package routers

import (
	"digimovie/src/handlers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	h := handlers.GetUserHelper()
	r.GET("/get/otp", h.GetOtp)
	r.GET("/validate/otp/new", h.ValidateOtpAndSignUp)
}