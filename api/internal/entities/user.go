package entities

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"regexp"
	"time"

	"github.com/boke0ya/beathub-api/internal/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(User) (User, error)
	Update(User) (User, error)
	DeleteById(id string) error
	FindById(id string) (User, error)
	FindByEmail(email string) (User, error)
}

type User struct {
	Id            string         `json:"id"`
	Email         UserEmail      `json:"-"`
	Password      UserPassword   `json:"-"`
	ScreenName    UserScreenName `json:"screenName"`
	Name          UserName       `json:"name"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	LastLoginedAt time.Time      `json:"lastLoginedAt"`
}

type LoginUser struct {
	User
}

type UserEmail string

func (un UserEmail) Validate() error {
	if ok, _ := regexp.MatchString(`^(.*)@(.*)$`, string(un)); ok {
		return nil
	} else {
		return errors.New(errors.InvalidEmailFormatError, nil)
	}
}

type UserPassword string

func NewUserPassword(password string) UserPassword {
	secret := os.Getenv("HAJIKKO_ONLINE_SECRET")
	hashedPasswordBytes := sha256.Sum256([]byte(password + secret))
	hashedPassword := hex.EncodeToString(hashedPasswordBytes[:])
	return UserPassword(hashedPassword)
}
func (un UserPassword) Validate() error {
	ok, _ := regexp.MatchString(`^\S+$`, string(un))
	if ok {
		return nil
	} else {
		return errors.New(errors.InvalidPasswordFormatError, nil)
	}
}
func (un UserPassword) Check(password string) bool {
	secret := os.Getenv("HAJIKKO_ONLINE_SECRET")
	hashedPasswordBytes := sha256.Sum256([]byte(password + secret))
	hashedPassword := hex.EncodeToString(hashedPasswordBytes[:])
	return hashedPassword == string(un)
}

type UserScreenName string

func (us UserScreenName) Validate() error {
	if ok, _ := regexp.MatchString(`^\S+$`, string(us)); ok {
		return nil
	} else {
		return errors.New(errors.InvalidScreenNameFormatError, nil)
	}
}

type UserName string

func (un UserName) Validate() error {
	if ok, _ := regexp.MatchString(`^(.+)$`, string(un)); ok {
		return nil
	} else {
		return errors.New(errors.InvalidNameFormatError, nil)
	}
}

func NewUser(email string, password string, screenName string, name string) (User, error) {
	user := User{
		Id:         uuid.NewString(),
		Email:      UserEmail(email),
		Password:   NewUserPassword(password),
		ScreenName: UserScreenName(screenName),
		Name:       UserName(name),
	}
	if err := user.Email.Validate(); err != nil {
		return user, err
	}
	if err := user.Password.Validate(); err != nil {
		return user, err
	}
	if err := user.ScreenName.Validate(); err != nil {
		return user, err
	}
	if err := user.Name.Validate(); err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) UpdateEmail(email string) error {
	u.Email = UserEmail(email)
	return u.Email.Validate()
}

func (u *User) UpdatePassword(currentPassword, password string) error {
	if u.Password != NewUserPassword(currentPassword) {
		return errors.New(errors.WrongPassword, nil)
	}
	u.Password = NewUserPassword(password)
	return u.Password.Validate()
}

func (u *User) UpdateProfile(screenName string, name string) error {
	u.ScreenName = UserScreenName(screenName)
	if err := u.ScreenName.Validate(); err != nil {
		return err
	}
	u.Name = UserName(name)
	return u.Name.Validate()
}

// Login returns JWT and update LastLoginedAt
func (u *User) Login(password string) (LoginUser, error) {
	if !u.Password.Check(password) {
		return LoginUser{}, errors.New(errors.WrongPassword, nil)
	}
	u.LastLoginedAt = time.Now()
	return LoginUser{
		User:           *u,
	}, nil
}

func (lu *LoginUser) GetAuthorizationToken() (string, error) {
	claims := jwt.StandardClaims{
		Id:        lu.Id,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 14).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("HAJIKKO_ONLINE_SECRET")))
	return tokenString, nil
}
