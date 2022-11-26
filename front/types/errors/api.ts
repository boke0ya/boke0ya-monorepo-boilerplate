export enum ApiErrorCode {
  InvalidEmailFormatError = 10102,
  InvalidPasswordFormatError = 10103,
  InvalidScreenNameFormatError = 10104,
  InvalidNameFormatError = 10105,
  EmailAlreadyExistsError = 20100,
  UserNotFoundError = 40100,
  EmailVerificationNotFoundError = 40102,
  AuthorizationRequired = 50100,
  WrongPassword = 50101,
  PermissionDenied = 50102,
  InvalidAuthozationToken = 50103,
  FailedToPersistBucketObject = 60000,
  FailedToSendEmailError = 60001,
  FailedToPersistUser = 60100,
  FailedToPersistEmailVerification = 60101,
}

export class ApiError extends Error {
  code: ApiErrorCode;
  message: string;
  info: string;
  constructor(props: {
    code: ApiErrorCode,
    message: string,
    info?: string
  }) {
    super()
    this.code = props.code
    this.message = props.message
    this.info = props.info
  }
  getMessage(): string {
    switch(this.code){
      case ApiErrorCode.InvalidEmailFormatError:
        return 'メールアドレスの形式が間違っています'
      case ApiErrorCode.InvalidPasswordFormatError:
        return 'パスワードの形式が間違っています'
      case ApiErrorCode.InvalidScreenNameFormatError:
        return 'IDの形式が間違っています'
      case ApiErrorCode.InvalidNameFormatError:
        return '名前の形式が間違っています'
      case ApiErrorCode.EmailAlreadyExistsError:
        return 'メールアドレスが既に登録されています'
      case ApiErrorCode.UserNotFoundError:
        return 'ユーザーが見つかりませんでした'
      case ApiErrorCode.EmailVerificationNotFoundError:
        return '本人確認が実施されていません'
      case ApiErrorCode.AuthorizationRequired:
        return 'ログインが必要です'
      case ApiErrorCode.WrongPassword:
        return 'パスワードが間違っています'
      case ApiErrorCode.PermissionDenied:
        return '権限がありません'
      case ApiErrorCode.InvalidAuthozationToken:
        return '不正なトークンが検出されました'
      default:
        return '処理に失敗しました'
    }
  }
}
