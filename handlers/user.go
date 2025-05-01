package handlers

import (
	"digimovie/src/config"
	"digimovie/src/database/models"
	"digimovie/src/dto"
	"digimovie/src/logging"
	"digimovie/src/responses"
	"digimovie/src/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHelper struct{
	log logging.Logger
}

func GetUserHelper() *UserHelper {
	return &UserHelper{log: logging.NewLogger()}
}

type getOtp struct {
	MobileNumber string `json:"mobileNumber" binding:"mobileNumber"`
}

func (h *UserHelper) GetOtp(c *gin.Context) {
	sample := &getOtp{}
	err := c.ShouldBindJSON(sample)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(false, http.StatusNotAcceptable, err))
		return
	}
	service_otp := services.NewOtpService(config.GetConfig())
	otp := services.MakeOtp()
	err = service_otp.SetOtp(sample.MobileNumber, otp, 2)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.GenerateResponseWithError(false, http.StatusBadRequest, err))
		return
	}
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(true, http.StatusOK, map[string]string{"otp": otp, "expire_time": "2 min"}))
}

type UserService struct {
	base services.BaseService[models.User, dto.UserCreate, dto.UserUpdate, dto.UserRes]
}

func newUserService() *UserService {
	return &UserService{base: *services.NewBaseService[models.User, dto.UserCreate, dto.UserUpdate, dto.UserRes]()}
}

func (h *UserHelper) NewUser(c *gin.Context) {
	user_service := newUserService()
	newUser := &dto.UserCreate{}
	err := c.ShouldBindJSON(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(false, http.StatusNotAcceptable, err))
		return
	}
	creater := c.Query("Userid")
	if len(creater) == 0 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, fmt.Errorf("userid creater didn't found")))
		return
	}
	c.AddParam("Userid", creater)
	res, err := user_service.base.Create(c, newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, responses.GenerateResponseWithError(false, http.StatusBadGateway, err))
		return
	}
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(true, http.StatusOK, res))
}