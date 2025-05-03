package handlers

import (
	"digimovie/src/config"
	"digimovie/src/database/models"
	"digimovie/src/dto"
	"digimovie/src/logging"
	"digimovie/src/responses"
	"digimovie/src/services"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var log = logging.NewLogger()

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
	sample := getOtp{}
	err := c.ShouldBindJSON(&sample)
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
	log.Info(logging.Otp, logging.Add, "new otp added to redis", nil)
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(true, http.StatusOK, map[string]string{"otp": otp, "expire_time": "2 min"}))
}

type UserService struct {
	base services.BaseService[models.User, dto.UserCreate, dto.UserUpdate, dto.UserRes]
}

func newUserService() *UserService {
	return &UserService{base: *services.NewBaseService[models.User, dto.UserCreate, dto.UserUpdate, dto.UserRes]()}
}

func (h *UserHelper) ValidateOtpAndSignUp(c *gin.Context) {
	user_service := newUserService()
	newUser := &dto.UserCreate{}
	err := c.ShouldBindJSON(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(false, http.StatusNotAcceptable, err))
		return
	}
	otp := c.Query("Otp")
	creater := c.Query("Userid")
	if len(creater) == 0 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, fmt.Errorf("userid creater didn't found")))
		return
	}
	creater2, err2 := strconv.Atoi(creater)
	if err2 != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, fmt.Errorf("userid creater is invalid")))
		return
	}
	newUser.CreatedBy = creater2
	fmt.Println(newUser.CreatedBy)
	service_otp := services.NewOtpService(config.GetConfig())
	err = service_otp.ValidateOtp(newUser.MobileNumber, otp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.GenerateResponseWithError(false, http.StatusBadRequest, err))
		return
	}
	password := newUser.Password
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, fmt.Errorf("error in bcrypt password")))
		return
	}
	newUser.Password = string(hashedpassword)
	res, err := user_service.base.Create(c, newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, responses.GenerateResponseWithError(false, http.StatusBadGateway, err))
		return
	}
	log.Info(logging.User, logging.New, "new user added", nil)
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(true, http.StatusOK, res))
}


func (h *UserHelper) EditInformation(c *gin.Context) {
	user_service := newUserService()
	user := dto.UserUpdate{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithValidationError(false, http.StatusNotAcceptable, err))
		return
	}
	user_id := c.Query("UserId")
	user_id2, err := strconv.Atoi(user_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, fmt.Errorf("id is invalid")))
		return
	}
	editor_id := c.Query("Editorid")
	if len(editor_id) == 0 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, fmt.Errorf("userid editor didn't found")))
		return
	}
	editor_id2, err2 := strconv.Atoi(editor_id)
	if err2 != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, fmt.Errorf("userid editor is invalid")))
		return
	}
	enable := user.Enabled
	res, err := user_service.base.Update(c, user_id2, &user, editor_id2, enable)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, err))
		return
	}
	log.Info(logging.User, logging.Edit, "user information edited", nil)
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(true, http.StatusOK, res))
}

func (h *UserHelper) Delete(c *gin.Context) {
	user_service := newUserService()
	user_id := c.Query("UserId")
	user_id2, err := strconv.Atoi(user_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, fmt.Errorf("UserId is invalid")))
		return
	}
	deleter_id := c.Query("DeleterId")
	if len(deleter_id) == 0 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, fmt.Errorf("DeleterId didn't found")))
		return
	}
	deleter_id2, err2 := strconv.Atoi(deleter_id)
	if err2 != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, fmt.Errorf("DeleterId is invalid")))
		return
	}
	err = user_service.base.Delete(c, user_id2, deleter_id2)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, responses.GenerateResponseWithError(false, http.StatusNotAcceptable, err))
		return
	}
	log.Info(logging.User, logging.Delete, "a user deleted", nil)
	c.JSON(http.StatusOK, responses.GenerateNormalResponse(true, http.StatusOK, "the user deleted"))
}