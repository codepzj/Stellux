import type { PostsReq, PostsVO } from "@/api/interfaces/posts";
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
) => Promise<PageResponse<PostsVO>> = (pageNo = 1, pageSize = 10) => {
  return request.get("/posts/list", { params: { pageNo, pageSize } });
};
