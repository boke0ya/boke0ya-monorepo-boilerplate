package converters

import (
	"time"

	. "app/internal/entities"
	. "app/internal/adapters"
)

type UserViewConverter struct {
	userIconBucketAdapter BucketAdapter
}

type UserView struct {
	Id            string    `json:"id"`
	ScreenName    string    `json:"screenName"`
	Name          string    `json:"name"`
	IconUrl       string    `json:"iconUrl"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	LastLoginedAt time.Time `json:"lastLoginedAt"`
}

func NewUserViewConverter(userIconBucketAdapter BucketAdapter) UserViewConverter {
	return UserViewConverter{
		userIconBucketAdapter: userIconBucketAdapter,
	}
}

func (uv UserViewConverter) Convert(user User) UserView {
	return UserView{
		Id:            user.Id,
		ScreenName:    string(user.ScreenName),
		Name:          string(user.Name),
		IconUrl:       uv.userIconBucketAdapter.GetObjectUrl(user.Id),
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
