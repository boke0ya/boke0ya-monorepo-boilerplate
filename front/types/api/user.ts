export interface User {
  id: string;
  name: string;
  screenName: string;
  iconUrl: string;
  createdAt: string;
  updatedAt: string;
}

export interface LoginRequest {
  email?: string;
  screenName?: string;
  password: string;
}

export interface LoginResponse {
  token: string;
}

export interface SignupEmailVerificationRequest {
  email: string;
}
