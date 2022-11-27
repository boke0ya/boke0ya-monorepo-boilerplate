package dao

import (
	"time"

	"app/internal/entities"
)

type User struct {
	Id            string    `gorm:"id" gorm:"primaryKey"`
	ScreenName    string    `gorm:"screenName"`
	Name          string    `gorm:"name"`
	Email         string    `gorm:"email"`
	Password      string    `gorm:"password"`
	CreatedAt     time.Time `gorm:"created_at"`
	UpdatedAt     time.Time `gorm:"updated_at"`
	LastLoginedAt time.Time `gorm:"last_logined_at"`
}

func (User) TableName() string {
	return "users"
}

func ConvertToUserEntity(user User) entities.User {
	return entities.User{
		Id:            user.Id,
		ScreenName:    entities.UserScreenName(user.ScreenName),
		Name:          entities.UserName(user.Name),
		Email:         entities.UserEmail(user.Email),
		Password:      entities.UserPassword(user.Password),
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		LastLoginedAt: user.LastLoginedAt,
	}
}

func ConvertFromUserEntity(user entities.User) User {
	return User{
		Id:            user.Id,
		ScreenName:    string(user.ScreenName),
		Name:          string(user.Name),
		Email:         string(user.Email),
		Password:      string(user.Password),
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		LastLoginedAt: user.LastLoginedAt,
	}
}
