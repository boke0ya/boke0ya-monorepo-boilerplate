package controllers

import (
	"net/http"

	. "app/internal/controllers/converters"
	. "app/internal/entities"
	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	Signup(token string, password string, screenName string, name string) (User, error)
	SignupEmailVerification(email string) error
	UpdateProfile(loginUser LoginUser, screenName string, name string) (User, error)
	CreateUpdateUserIconUrl(loginUser LoginUser) (string, error)
	UpdateEmail(loginUser LoginUser, token string) (User, error)
	UpdateEmailVerification(loginUser LoginUser, email string) error
	UpdatePassword(loginUser LoginUser, currentPassword string, password string) (User, error)
	Withdraw(loginUser LoginUser, password string) error
	Login(email string, password string) (string, error)
	Session(token string) (LoginUser, error)
}

type UsersController struct {
	Controller
	userUsecase       UserUsecase
	userViewConverter UserViewConverter
}

func NewUsersController(userUsecase UserUsecase, userViewConverter UserViewConverter) UsersController {
	return UsersController{
		userUsecase:       userUsecase,
		userViewConverter: userViewConverter,
	}
}

type EmailVerificationRequest struct {
	Email string `json:"email"`
}

func (uc UsersController) SignupEmailVerification(c *gin.Context) {
	var req EmailVerificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	email := req.Email
	err := uc.userUsecase.SignupEmailVerification(email)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, nil)
}

type UpdateEmailRequest struct {
	Email string `json:"email"`
}

func (uc UsersController) UpdateEmailVerification(c *gin.Context) {
	var req UpdateEmailRequest
	loginUser := uc.GetLoginUser(c)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	email := req.Email
	err := uc.userUsecase.UpdateEmailVerification(*loginUser, email)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, nil)
}

type SignupRequest struct {
	Token      string `json:"token"`
	Name       string `json:"name"`
	ScreenName string `json:"screenName"`
	Password   string `json:"password"`
}

type SignupResponse struct {
	UserView
	Token string `json:"token"`
}

func (uc UsersController) Signup(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	user, err := uc.userUsecase.Signup(req.Token, req.Password, req.ScreenName, req.Name)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	loginUser := LoginUser{User: user}
	token, err := loginUser.GetAuthorizationToken()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, SignupResponse{
		UserView: uc.userViewConverter.Convert(user),
		Token:    token,
	})
}

type UpdateEmailVerificationRequest struct {
	Token string `json:"token"`
}

func (uc UsersController) UpdateEmail(c *gin.Context) {
	var req UpdateEmailVerificationRequest
	loginUser := uc.GetLoginUser(c)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	token := req.Token
	user, err := uc.userUsecase.UpdateEmail(*loginUser, token)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, uc.userViewConverter.Convert(user))
}

type UpdatePasswordRequest struct {
	CurrentPassword string `json:"currentPassword"`
	Password        string `json:"password"`
}

func (uc UsersController) UpdatePassword(c *gin.Context) {
	var req UpdatePasswordRequest
	loginUser := uc.GetLoginUser(c)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	user, err := uc.userUsecase.UpdatePassword(*loginUser, req.CurrentPassword, req.Password)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, uc.userViewConverter.Convert(user))

}

type WithdrawRequest struct {
	Password string `json:"password"`
}

func (uc UsersController) Withdraw(c *gin.Context) {
	var req WithdrawRequest
	loginUser := uc.GetLoginUser(c)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	err := uc.userUsecase.Withdraw(*loginUser, req.Password)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, nil)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (uc UsersController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	email := req.Email
	password := req.Password
	token, err := uc.userUsecase.Login(email, password)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
	})
}

func (uc UsersController) Session(c *gin.Context) {
	loginUser := uc.GetLoginUser(c)
	c.JSON(http.StatusOK, uc.userViewConverter.Convert(loginUser.User))
}

type UpdateProfileRequest struct {
	ScreenName string `json:"screenName"`
	Name       string `json:"name"`
}

func (uc UsersController) UpdateProfile(c *gin.Context) {
	var req UpdateProfileRequest
	loginUser := uc.GetLoginUser(c)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	user, err := uc.userUsecase.UpdateProfile(*loginUser, req.ScreenName, req.Name)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, uc.userViewConverter.Convert(user))
}

type CreatePutUserIconUrlResponse struct {
	Url string `json:"url"`
}

func (uc UsersController) CreatePutUserIconUrl(c *gin.Context) {
	loginUser := uc.GetLoginUser(c)
	url, err := uc.userUsecase.CreateUpdateUserIconUrl(*loginUser)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, CreatePutUserIconUrlResponse{
		Url: url,
	})
}
