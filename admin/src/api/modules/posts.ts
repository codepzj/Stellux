import type {
  PostReq,
  PostUpdateStatusReq,
  PostVO,
} from "@/api/interfaces/posts";
import type { PageResponse, Response } from "@/api/interfaces/resp";
import request from "@/utils/request";

export const createPost: (data: PostReq) => Promise<Response<any>> = (data) => {
  console.log("data", data);
  return request.post("/admin-api/posts/create", data);
};

// 分页获取文章列表
export const getPostsByPage: (
  pageNo?: number,
  pageSize?: number
) => Promise<PageResponse<PostVO>> = (page_no = 1, size = 10) => {
  return request.get("/posts/list", { params: { page_no, size } });
};

// 更新文章状态
export const updatePostStatus: (
  data: PostUpdateStatusReq
) => Promise<PageResponse<any>> = (data) => {
  return request.patch("/admin-api/posts/status", data);
};

// 软删除文章
export const deletePostSoft: (id: string) => Promise<Response<any>> = (id) => {
  return request.delete(`/admin-api/posts/delete/${id}`);
};
