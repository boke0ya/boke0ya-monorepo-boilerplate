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
	emailVerificationTx := uu.emailVerificationRepository.Begin()
	userTx := uu.userRepository.Begin()
	defer func() {
		emailVerificationTx.Rollback()
		userTx.Rollback()
	}()
	emailVerification, err := emailVerificationTx.FindByToken(emailToken)
	if err != nil {
		return User{}, err
	}
	user, err := emailVerification.Signup(password, screenName, name)
	if err != nil {
		return user, err
	}
	user, err = userTx.Create(user)
	emailVerificationTx.DeleteById(emailVerification.Id)
	emailVerificationTx.Commit()
	userTx.Commit()
	return user, err
}

func (uu userUsecase) SignupEmailVerification(email string) error {
	emailVerificationTx := uu.emailVerificationRepository.Begin()
	userTx := uu.userRepository.Begin()
	defer func() {
		emailVerificationTx.Rollback()
		userTx.Rollback()
	}()
	emailVerification, err := NewEmailVerification(nil, email)
	if err != nil {
		return err
	}
	if _, err := userTx.FindByEmail(email); err == nil {
		return errors.New(errors.EmailAlreadyExistsError, nil)
	}
	if _, err := emailVerificationTx.Create(emailVerification); err != nil {
		return err
	}
	err = uu.mailAdapter.Send(
		email,
		"<サービス名> 本人確認",
		"<サービス名>のご登録ありがとうございます。\r\n"+
			"\r\n"+
			"下記URLより、本登録へお進み下さい\r\n"+
			uu.baseUrl+"/signup?token="+emailVerification.Token+"\r\n"+
			"\r\n"+
			"※このメールに心当たりが無い方は、メールを削除して頂いて構いません。\r\n"+
			"ご迷惑をおかけし申し訳ありません。",
	)
	if err != nil {
		return err
	}
	emailVerificationTx.Commit()
	userTx.Commit()
	return nil
}

func (uu userUsecase) UpdateEmail(loginUser LoginUser, emailToken string) (User, error) {
	emailVerificationTx := uu.emailVerificationRepository.Begin()
	userTx := uu.userRepository.Begin()
	defer func() {
		emailVerificationTx.Rollback()
		userTx.Rollback()
	}()
	emailVerification, err := emailVerificationTx.FindByToken(emailToken)
	if err != nil {
		return User{}, err
	}
	user, err := userTx.FindById(loginUser.Id)
	if err != nil {
		return user, err
	}
	if err := user.UpdateEmail(string(emailVerification.Email)); err != nil {
		return user, err
	}
	user, err = userTx.Update(user)
	if err != nil {
		return user, err
	}
	if err := emailVerificationTx.DeleteById(emailVerification.Id); err != nil {
		return user, err
	}
	emailVerificationTx.Commit()
	userTx.Commit()
	return user, err
}

func (uu userUsecase) UpdateEmailVerification(loginUser LoginUser, email string) error {
	emailVerificationTx := uu.emailVerificationRepository.Begin()
	userTx := uu.userRepository.Begin()
	defer func() {
		emailVerificationTx.Rollback()
		userTx.Rollback()
	}()
	emailVerification, err := NewEmailVerification(&loginUser.User, email)
	if err != nil {
		return err
	}
	if _, err := uu.userRepository.FindByEmail(email); err == nil {
		return errors.New(errors.EmailAlreadyExistsError, nil)
	}
	uu.emailVerificationRepository.Create(emailVerification)
	err = uu.mailAdapter.Send(
		email,
		"<サービス名> 本人確認",
		"メールアドレスの変更を受け付けました。\r\n"+
			"\r\n"+
			"下記URLより、変更を完了して下さい。\r\n"+
			uu.baseUrl+"/settings?token="+emailVerification.Token+"\r\n"+
			"\r\n"+
			"※このメールに心当たりが無い方は、メールを削除して頂いて構いません。\r\n"+
			"ご迷惑をおかけし申し訳ありません。",
	)
	if err != nil {
		return err
	}
	userTx.Commit()
	emailVerificationTx.Commit()
	return nil
}

func (uu userUsecase) UpdatePassword(loginUser LoginUser, currentPassword string, password string) (User, error) {
	tx := uu.userRepository.Begin()
	defer tx.Rollback()
	user, err := tx.FindById(loginUser.Id)
	if err != nil {
		return user, err
	}
	if err := user.UpdatePassword(currentPassword, password); err != nil {
		return user, err
	}
	if _, err := tx.Update(user); err != nil {
		return user, err
	}
	tx.Commit()
	return user, nil
}

func (uu userUsecase) UpdateProfile(loginUser LoginUser, screenName string, name string) (User, error) {
	tx := uu.userRepository.Begin()
	defer tx.Rollback()
	user, err := tx.FindById(loginUser.Id)
	if err != nil {
		return user, err
	}
	if err := user.UpdateProfile(screenName, name); err != nil {
		return user, err
	}
	user, err = tx.Update(user)
	tx.Commit()
	return user, err
}

func (uu userUsecase) CreateUpdateUserIconUrl(loginUser LoginUser) (string, error) {
	url, err := uu.iconBucketAdapter.CreatePutObjectUrl(loginUser.Id)
	return url, err
}

func (uu userUsecase) Withdraw(loginUser LoginUser, password string) error {
	tx := uu.userRepository.Begin()
	user, err := tx.FindById(loginUser.Id)
	if err != nil {
		return err
	}
	if _, err := user.Login(password); err != nil {
		return err
	}
	if err := tx.DeleteById(user.Id); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (uu userUsecase) LoginByEmail(email string, password string) (string, error) {
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

func (uu userUsecase) LoginByScreenName(screenName string, password string) (string, error) {
	user, err := uu.userRepository.FindByScreenName(screenName)
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
