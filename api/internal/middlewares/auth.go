package middlewares

import (
	"strings"

	. "github.com/boke0ya/beathub-api/internal/controllers"
	"github.com/boke0ya/beathub-api/internal/errors"
	"github.com/gin-gonic/gin"
)

//Authenticate returns authentication middleware
func Authenticate(userUsecase UserUsecase) gin.HandlerFunc {
    return func(c *gin.Context){
        authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
        seps := strings.Split(authHeader, " ")
        if len(seps) < 2 {
            c.Next()
            return
        }
        tokenString := seps[1]
        user, err := userUsecase.Session(tokenString)
        if err != nil {
            c.Next()
            return
        }
        c.Set("loginUser", &user)
        c.Next()
    }
}

//Authenticate returns authorize middleware for endpoint
func Authorize() gin.HandlerFunc {
    return func(c *gin.Context){
        loginUser, ok := c.Get("loginUser")
        if loginUser == nil || !ok {
            c.Error(errors.New(errors.AuthorizationRequired, nil)).SetType(gin.ErrorTypePublic)
            c.Abort()
            return
        }
        c.Next()
    }
}

