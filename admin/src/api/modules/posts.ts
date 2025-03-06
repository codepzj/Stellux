import type {
  PostsReq,
  PostsUpdateStatusReq,
  PostsVO,
} from "@/api/interfaces/posts";
import type { PageResponse, Response } from "@/api/interfaces/resp";
import request from "@/utils/request";
export const createPost: (data: PostsReq) => Promise<Response<any>> = (
  data
) => {
  console.log("data", data);
  return request.post("/posts/create", data);
};

// 分页获取文章列表
export const getPostsByPage: (
  pageNo?: number,
  pageSize?: number
) => Promise<PageResponse<PostsVO>> = (page_no = 1, size = 10) => {
  return request.get("/posts/list", { params: { page_no, size } });
};

// 更新文章状态
export const updatePostStatus: (
  data: PostsUpdateStatusReq
) => Promise<PageResponse<any>> = (data) => {
  return request.put("/posts/update/status", data);
};

// 软删除文章
export const deletePostSoft: (id: string) => Promise<Response<any>> = (id) => {
  return request.delete(`/posts/soft-delete/${id}`);
};
