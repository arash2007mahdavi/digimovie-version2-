package handlers

import "github.com/gin-gonic/gin"

type UserHelper struct{}

func GetUserHelper() *UserHelper {
	return &UserHelper{}
}

type NewUserDto struct {
	
}

func (h *UserHelper) NewUser(c *gin.Context) {

}