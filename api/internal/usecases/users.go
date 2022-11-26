package usecases

import (
	. "app/internal/adapters"
	. "app/internal/controllers"
	. "app/internal/entities"
	"app/internal/errors"
	"github.com/golang-jwt/jwt/v4"
)

// UserUsecase provides usecase of CRUD for User(eg. signup, signin, withdraw, settings)
type userUsecase struct {
	iconBucketAdapter           BucketAdapter
	mailAdapter                 MailAdapter
	userRepository              UserRepository
	emailVerificationRepository EmailVerificationRepository
	baseUrl                     string
	secret                      string
}

// NewUserUsecase creates UserUsecase
func NewUserUsecase(
	iconBucketAdapter BucketAdapter,
	mailAdapter MailAdapter,
	userRepository UserRepository,
	emailVerificationRepository EmailVerificationRepository,
	baseUrl string,
	secret string,
) UserUsecase {
	return userUsecase{
		iconBucketAdapter:           iconBucketAdapter,
		mailAdapter:                 mailAdapter,
		userRepository:              userRepository,
		emailVerificationRepository: emailVerificationRepository,
		baseUrl:                     baseUrl,
		secret:                      secret,
	}
}

func (uu userUsecase) Signup(emailToken string, password string, screenName string, name string) (User, error) {
	emailVerification, err := uu.emailVerificationRepository.FindByToken(emailToken)
	if err != nil {
		return User{}, err
	}
	user, err := NewUser(string(emailVerification.Email), password, screenName, name)
	if err != nil {
		return user, err
	}
	user, err = uu.userRepository.Create(user)
	uu.emailVerificationRepository.DeleteById(emailVerification.Id)
	return user, err
}

func (uu userUsecase) SignupEmailVerification(email string) error {
	emailVerification, err := NewEmailVerification(nil, email)
	if err != nil {
		return err
	}
	if _, err := uu.userRepository.FindByEmail(email); err == nil {
		return errors.New(errors.EmailAlreadyExistsError, nil)
	}
	uu.emailVerificationRepository.Create(emailVerification)
	return uu.mailAdapter.Send(
		email,
		"BeatHub 本人確認",
		"BeatHubのご登録ありがとうございます。\r\n"+
			"\r\n"+
			"下記URLより、本登録へお進み下さい\r\n"+
			uu.baseUrl+"/signup?token="+emailVerification.Token+"\r\n"+
			"\r\n"+
			"※このメールに心当たりが無い方は、メールを削除して頂いて構いません。\r\n"+
			"ご迷惑をおかけし申し訳ありません。",
	)
}

func (uu userUsecase) UpdateEmail(loginUser LoginUser, emailToken string) (User, error) {
	emailVerification, err := uu.emailVerificationRepository.FindByToken(emailToken)
	if err != nil {
		return loginUser.User, err
	}
	if err := loginUser.UpdateEmail(string(emailVerification.Email)); err != nil {
		return loginUser.User, err
	}
	user, err := uu.userRepository.Update(loginUser.User)
	if err != nil {
		return user, err
	}
	err = uu.emailVerificationRepository.DeleteById(emailVerification.Id)
	return user, err
}

func (uu userUsecase) UpdateEmailVerification(loginUser LoginUser, email string) error {
	emailVerification, err := NewEmailVerification(&loginUser.User, email)
	if err != nil {
		return err
	}
	if _, err := uu.userRepository.FindByEmail(email); err == nil {
		return errors.New(errors.EmailAlreadyExistsError, nil)
	}
	uu.emailVerificationRepository.Create(emailVerification)
	return uu.mailAdapter.Send(
		email,
		"BeatHub 本人確認",
		"メールアドレスの変更を受け付けました。\r\n"+
			"\r\n"+
			"下記URLより、変更を完了して下さい。\r\n"+
			uu.baseUrl+"/settings?token="+emailVerification.Token+"\r\n"+
			"\r\n"+
			"※このメールに心当たりが無い方は、メールを削除して頂いて構いません。\r\n"+
			"ご迷惑をおかけし申し訳ありません。",
	)
}

func (uu userUsecase) UpdatePassword(loginUser LoginUser, currentPassword string, password string) (User, error) {
	if err := loginUser.UpdatePassword(currentPassword, password); err != nil {
		return loginUser.User, err
	}
	return uu.userRepository.Update(loginUser.User)
}

func (uu userUsecase) UpdateProfile(loginUser LoginUser, screenName string, name string) (User, error) {
	if err := loginUser.UpdateProfile(screenName, name); err != nil {
		return loginUser.User, err
	}
	user, err := uu.userRepository.Update(loginUser.User)
	return user, err
}

func (uu userUsecase) CreateUpdateUserIconUrl(loginUser LoginUser) (string, error) {
	url, err := uu.iconBucketAdapter.CreatePutObjectUrl(loginUser.Id)
	return url, err
}

func (uu userUsecase) Withdraw(loginUser LoginUser, password string) error {
	if _, err := loginUser.Login(password); err != nil {
		return err
	}
	return uu.userRepository.DeleteById(loginUser.Id)
}

func (uu userUsecase) Login(email string, password string) (string, error) {
	user, err := uu.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}
	loginUser, err := user.Login(password)
	if err != nil {
		return "", err
	}
	return loginUser.GetAuthorizationToken()
}

func (uu userUsecase) Session(tokenString string) (LoginUser, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(uu.secret), nil
	})
	if err != nil {
		return LoginUser{}, errors.New(errors.InvalidAuthozationToken, err)
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		user, err := uu.userRepository.FindById(claims.Id)
		if err != nil {
			return LoginUser{}, errors.New(errors.InvalidAuthozationToken, err)
		}
		return LoginUser{
			User: user,
		}, nil
	}else{
		return LoginUser{}, errors.New(errors.InvalidAuthozationToken, nil)
	}
}
