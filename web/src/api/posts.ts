import { Page, PageVO } from "@/types/page";
import { PostVO } from "@/types/posts";
import request from "@/utils/request";

// 获取首页文章列表
export const getPostsList = (page: Page) => {
  return request.get<PageVO<PostVO>>("/posts/list", page);
};
