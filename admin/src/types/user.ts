export interface User {
  id: string;
  username: string;
  role_id: number;
}

export interface LoginForm extends LoginReq {
  remember: boolean;
}

export interface LoginReq {
  username: string;
  password: string;
}

export interface LoginVO {
  token: string; // 令牌
  user: User;
}

export interface UserManage {
  id: string;
  created_at: string;
  updated_at: string;
  username: string;
  password: string;
  role_id: number;
}

export type UserListVO = UserManage[];
