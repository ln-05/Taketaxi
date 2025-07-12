// API响应类型
export interface ApiResponse<T = any> {
  code: number;
  data: T;
  msg: string;
}

// 短信相关类型
export interface SendSmsRequest {
  phone: string;
  source: string;
}

export interface SendSmsResponse {
  code: number;
  msg: string;
}

// 用户相关类型
export interface UserInfo {
  id: string;
  phone: string;
  nickname?: string;
  avatar?: string;
  createTime?: string;
  updateTime?: string;
}

// 微信相关类型
export interface WechatUserInfo {
  openid: string;
  nickname: string;
  avatar: string;
  userId: string;
}

export interface WechatLoginStatus {
  status: 'waiting' | 'scanned' | 'success' | 'expired';
  user_info?: WechatUserInfo;
  access_token?: string;
}

export interface WechatCallbackRequest {
  code: string;
  state: string;
}

// 登录相关类型
export interface PhoneLoginRequest {
  phone: string;
  code: string;
}

export interface WechatLoginRequest {
  openid: string;
  nickname: string;
  avatar: string;
}

export interface LoginResponse {
  token: string;
  userInfo: UserInfo;
}

// 统一的业务响应类型
export interface BusinessResponse<T = any> {
  code: number;
  data: T;
  msg: string;
} 