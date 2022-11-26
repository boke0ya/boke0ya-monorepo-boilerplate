package controllers

import (
	. "github.com/boke0ya/beathub-api/internal/entities"
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
