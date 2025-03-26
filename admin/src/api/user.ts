import type { LoginForm, LoginVO, UserListVO } from "@/types/user";
import type { Response } from "@/types/response";
import request from "@/utils/request";

export const userLogin: (
  data: LoginForm
) => Promise<Response<LoginVO>> = data => {
  return request.post("/user/login", data);
};

export const getUserList: () => Promise<Response<UserListVO>> = () => {
  return request.post("/user/list");
};
