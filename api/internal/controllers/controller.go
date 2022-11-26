package controllers

import (
	. "app/internal/entities"
	"github.com/gin-gonic/gin"
)

type Controller struct { }

func (ctrl Controller) GetLoginUser(c *gin.Context) *LoginUser {
	user, ok := c.Get("loginUser")
	if ok {
		return user.(*LoginUser)
	}
	return nil
}
