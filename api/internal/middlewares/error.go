package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	. "github.com/boke0ya/beathub-api/internal/errors"
	"github.com/gin-gonic/gin"
)

func Error() func(c *gin.Context) {
    return func(c *gin.Context) {
        c.Next()
        err := c.Errors.ByType(gin.ErrorTypePublic).Last()
        if err != nil {
			var err_ ApiError
			if errors.As(err.Err, &err_) {
				var code int
				fmt.Println(err_.Error())
				if err_.Internal != nil {
					fmt.Println(err_.Internal.Error())
				}
				if err_.Code < 30000 {
					code = 400
				} else if err_.Code < 40000 {
					code = 403
				} else if err_.Code < 50000 {
					code = 404
				} else if err_.Code < 60000 {
					code = 401
				} else {
					code = 500
				}
				if os.Getenv("ENV") != "production" {
					var info string
					if err_.Internal != nil {
						info = err_.Internal.Error()
					}
					c.AbortWithStatusJSON(code, DebugErrorResponse{
						Code:    int(err_.Code),
						Message: err_.Error(),
						Info:    info,
					})
				} else {
					c.AbortWithStatusJSON(code, ErrorResponse{
						Code:    int(err_.Code),
						Message: err_.Error(),
					})
				}
			} else {
				fmt.Println(err.Error())
				if os.Getenv("ENV") != "production" {
					c.AbortWithStatusJSON(http.StatusInternalServerError, DebugErrorResponse{
						Code:    60000,
						Message: "Some error occured.",
						Info:    err.Error(),
					})
				} else {
					c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
						Code:    60000,
						Message: "Some error occured.",
					})
				}
			}
        }
    }
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DebugErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Info    string `json:"info"`
}
