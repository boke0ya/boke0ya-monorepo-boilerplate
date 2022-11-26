/*
Error

このパッケージでは、全体で扱うエラーを定義します。

Rules

10000番代 : リクエストされた値の違反によるエラー
- リクエストされた値がバリデーションに失敗した時に発生する。
- 400 Bad Request

20000番代 : ユーザーの行動規律違反によるエラー
- ユーザーが想定外の挙動をした場合に発生する。
- 400 Bad Request

30000番代 : 権限超越によるエラー
- ユーザーが認められた権限を越えた操作をしようとした時に発生する。
- 403 Forbidden

40000番代 : リソースが存在しない場合に発生するエラー
- リクエストされた値、もしくはリクエストに関連して呼び出されるべきリソースが存在しない場合に発生する。
- 404 Not Found

50000番代 : 認証に失敗した場合に発生するエラー
- 認証処理の過程で失敗した場合に発生する。
- 401 Unauthorized

60000番代 : その他サーバーの不具合によるエラー
- サーバー側のなにかしらの不具合により、処理が実行できなかった場合に発生する。
- 500 Internal Server Error

------

0000番代 : 全般エラー
0100番代 : Userエンティティ関連のエラー
*/

package errors

import "fmt"

type ErrorCode uint

const (
	OK                               ErrorCode = 0
	InvalidEmailFormatError                    = 10102
	InvalidPasswordFormatError                 = 10103
	InvalidScreenNameFormatError               = 10104
	InvalidNameFormatError                     = 10105
	EmailAlreadyExistsError                    = 20100
	UserNotFoundError                          = 40100
	EmailVerificationNotFoundError             = 40102
	AuthorizationRequired                      = 50100
	WrongPassword                              = 50101
	PermissionDenied                           = 50102
	InvalidAuthozationToken                    = 50103
	FailedToPersistBucketObject                = 60000
	FailedToSendEmailError                     = 60001
	FailedToPersistUser                        = 60100
	FailedToPersistEmailVerification           = 60101
)

type ApiError struct {
	Code     ErrorCode
	Internal error
}

func (err ApiError) Error() string {
	return fmt.Sprintf("%s", err.Message())
}

func (err ApiError) Message() string {
	switch err.Code {
	case OK:
		return "OK"
	case InvalidEmailFormatError:
		return "Invalid email format"
	case InvalidPasswordFormatError:
		return "Invalid password format"
	case InvalidScreenNameFormatError:
		return "Invalid screen name format"
	case InvalidNameFormatError:
		return "Invalid name format"
	case EmailAlreadyExistsError:
		return "Email already exists"
	case UserNotFoundError:
		return "User not found"
	case AuthorizationRequired:
		return "Authorization required"
	case WrongPassword:
		return "Wrong password"
	case PermissionDenied:
		return "Permission denied"
	case InvalidAuthozationToken:
		return "Invalid authorization token"
	case FailedToSendEmailError:
		return "Failed to send email"
	case FailedToPersistUser:
		return "Failed to persist user"
	case FailedToPersistBucketObject:
		return "Failed to persist bucket object"
	}
	return ""
}

func New(code ErrorCode, internal error) ApiError {
	return ApiError{
		Code:     code,
		Internal: internal,
	}
}
