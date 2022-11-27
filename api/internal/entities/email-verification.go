package entities

import (
	"app/internal/errors"
	"time"

	"github.com/google/uuid"
)

type EmailVerificationRepository interface {
	Begin() EmailVerificationRepository
	Commit() error
	Rollback()
	Create(emailVerification EmailVerification) (EmailVerification, error)
	Update(emailVerification EmailVerification) (EmailVerification, error)
	DeleteById(emailVerificationId string) error
	FindById(emailVerificaitonId string) (EmailVerification, error)
	FindByToken(token string) (EmailVerification, error)
}

type EmailVerification struct {
	Id        string    `json:"id"`
	Email     UserEmail `json:"email"`
	Token     string    `json:"token"`
	UserId    *string   `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewEmailVerification(user *User, email string) (EmailVerification, error) {
	var userId *string = nil
	if user != nil {
		userId = &user.Id
	}
	emailVerification := EmailVerification{
		Id:     uuid.NewString(),
		Email:  UserEmail(email),
		Token:  uuid.NewString(),
		UserId: userId,
	}
	if err := emailVerification.Email.Validate(); err != nil {
		return emailVerification, err
	}
	return emailVerification, nil
}

func (ev EmailVerification) Signup(password string, screenName string, name string) (User, error) {
	if ev.UserId != nil {
		return User{}, errors.New(errors.EmailVerificationNotForSignup, nil)
	}
	return NewUser(
		string(ev.Email),
		password,
		screenName,
		name,
	)
}
