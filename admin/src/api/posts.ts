import type {
  PostReq,
  PostUpdateStatusReq,
  PostVO,
} from "@/types/posts";
import type { PageReq, PageResponse, Response } from "@/types/response";
import request from "@/utils/request";

// 创建文章
export const createPost: (data: PostReq) => Promise<Response<any>> = data => {
  return request.post("/admin-api/posts/create", data);
};

// 更新文章
export const updatePost: (data: PostReq) => Promise<Response<any>> = data => {
  return request.put("/admin-api/posts/update", data);
};

// 根据id获取文章
export const getPostById: (id: string) => Promise<Response<PostVO>> = id => {
  return request.get(`/posts/${id}`);
};

// 分页获取文章列表
export const getPostsByPage: ({
  page_no,
  size,
}: PageReq) => Promise<PageResponse<PostVO>> = ({
  page_no = 1,
  size = 10,
  keyword,
  field,
  order,
}) => {
  return request.post("/admin-api/posts/list", {
    page_no,
    size,
    keyword,
    field,
    order,
  });
};

// 更新文章状态
export const updatePostStatus: (
  data: PostUpdateStatusReq
) => Promise<PageResponse<any>> = data => {
  return request.patch("/admin-api/posts/status", data);
};

// 软删除文章
export const deletePostSoft: (id: string) => Promise<Response<any>> = id => {
  return request.patch(`/admin-api/posts/delete/${id}`);
};
