package main

import (
	"net/http"
	"os"

	. "app/internal/controllers"
	. "app/internal/controllers/converters"
	. "app/internal/infrastructures/gmail"
	. "app/internal/infrastructures/gorm"
	. "app/internal/infrastructures/gorm/repositories"
	. "app/internal/infrastructures/s3"
	. "app/internal/middlewares"
	. "app/internal/usecases"
	"github.com/gin-gonic/gin"
)

func main() {
	db := NewGormDatabase(
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
	)

	userRepository := NewGormUserRepository(db)
	emailVerificationRepository := NewGormEmailVerificationRepository(db)

	userIconBucketAdapter := NewS3Adapter(
		os.Getenv("BUCKET_PRIVATE_BASE_URL"),
		os.Getenv("BUCKET_PUBLIC_BASE_URL"),
		"beathub-user-icon",
		os.Getenv("BUCKET_ACCESS_KEY_ID"),
		os.Getenv("BUCKET_ACCESS_SECRET"),
	)
	mailAdapter := NewGMailAdapter(os.Getenv("GMAIL_EMAIL"), os.Getenv("GMAIL_PASSWORD"))

	userUsecase := NewUserUsecase(
		userIconBucketAdapter,
		mailAdapter,
		userRepository,
		emailVerificationRepository,
		os.Getenv("BASE_URL"),
		os.Getenv("JWT_SECRET"),
	)

	userViewConverter := NewUserViewConverter()

	usersController := NewUsersController(userUsecase, userViewConverter)

	r := gin.Default()
	v1 := r.Group("/v1")
	v1.Use(Error(), Authenticate(userUsecase))
	{
		//Endpoint for health check
		v1.GET("", func (c *gin.Context) {
			c.Status(http.StatusOK)
		})
		v1.POST("/login", usersController.Login)
		v1.GET("/session", Authorize(), usersController.Session)
		v1.POST("/email-verification/signup", usersController.SignupEmailVerification)
		v1.POST("/email-verification/update", Authorize(), usersController.UpdateEmailVerification)
		v1.POST("/signup", usersController.Signup)
		users := v1.Group("/users")
		users.Use(Authorize())
		{
			users.PUT("/email", usersController.UpdateEmail)
			users.PUT("/password", usersController.UpdatePassword)
			users.PUT("/profile", usersController.UpdateProfile)
			users.POST("/profile/icon", usersController.CreatePutUserIconUrl)
			users.DELETE("", usersController.Withdraw)
		}
	}
	r.Run()
}
