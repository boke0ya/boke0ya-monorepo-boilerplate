package dao

import (
	"time"

	"github.com/boke0ya/beathub-api/internal/entities"
)

type EmailVerification struct {
	Id        string    `gorm:"id" gorm:"primaryKey"`
	Email     string    `gorm:"email"`
	Token     string    `gorm:"token" gorm:"uniqueIndex"`
	UserID    *string   `gorm:"user_id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (EmailVerification) TableName() string {
	return "email_verifications"
}

func ConvertToEmailVerificationEntity(emailVerification EmailVerification) entities.EmailVerification {
	return entities.EmailVerification{
		Id:        emailVerification.Id,
		Email:     entities.UserEmail(emailVerification.Email),
		Token:     emailVerification.Token,
		UserId:    emailVerification.UserID,
		CreatedAt: emailVerification.CreatedAt,
	}
}

func ConvertFromEmailVerificationEntity(emailVerification entities.EmailVerification) EmailVerification {
	return EmailVerification{
		Id:        emailVerification.Id,
		Email:     string(emailVerification.Email),
		Token:     emailVerification.Token,
		UserID:    emailVerification.UserId,
		CreatedAt: emailVerification.CreatedAt,
	}
}
