export interface PostsReq {
  title: string;
  content: string;
  author: string;
  description: string;
  category: string;
  tags: string[];
  cover: string;
  isTop: boolean;
}
