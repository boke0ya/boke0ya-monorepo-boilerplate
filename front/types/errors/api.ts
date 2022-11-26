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
}
