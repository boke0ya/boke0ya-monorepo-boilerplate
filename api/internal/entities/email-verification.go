package entities

import (
	"time"

	"github.com/google/uuid"
)

type EmailVerificationRepository interface {
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
	var userId string
	if user != nil {
		userId = user.Id
	}
	emailVerification := EmailVerification{
		Id:     uuid.NewString(),
		Email:  UserEmail(email),
		Token:  uuid.NewString(),
		UserId: &userId,
	}
	if err := emailVerification.Email.Validate(); err != nil {
		return emailVerification, err
	}
	return emailVerification, nil
}
