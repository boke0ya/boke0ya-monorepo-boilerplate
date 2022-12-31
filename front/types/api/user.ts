export interface User {
  id: string;
  name: string;
  screenName: string;
  iconUrl: string;
  createdAt: string;
  updatedAt: string;
}

export interface LoginRequest {
  id: string;
  password: string;
}

export interface LoginResponse {
  token: string;
}

export interface SignupEmailVerificationRequest {
  email: string;
}

export interface SignupRequest {
  token: string;
  name: string;
  screenName: string;
  password: string;
}

export interface SignupResponse extends User {
  token: string;
}

export interface UpdateEmailRequest {
  email: string;
}

export interface UpdateEmailResponse {}

export interface UpdatePasswordRequest {
  currentPassword: string;
  password: string;
}

export interface UpdatePasswordResponse {}

export interface UpdateProfileRequest {
  name: string;
  screenName: string;
}

export interface UpdateProfileResponse extends User {}
