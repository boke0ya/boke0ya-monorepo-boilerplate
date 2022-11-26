package converters

import (
	"time"

	. "app/internal/entities"
)

type UserViewConverter struct {
}

type UserView struct {
	Id            string    `json:"id"`
	ScreenName    string    `json:"screenName"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	LastLoginedAt time.Time `json:"lastLoginedAt"`
}

func NewUserViewConverter() UserViewConverter {
	return UserViewConverter{}
}

func (uv UserViewConverter) Convert(user User) UserView {
	return UserView{
		Id:            user.Id,
		ScreenName:    string(user.ScreenName),
		Name:          string(user.Name),
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		LastLoginedAt: user.LastLoginedAt,
	}
}

func (uv UserViewConverter) ConvertAll(users []User) []UserView {
	var userViews []UserView
	for _, user := range users {
		userViews = append(userViews, uv.Convert(user))
	}
	return userViews
}
