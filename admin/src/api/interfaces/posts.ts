export interface PostsReq {
  title: string;
  content: string;
  author: string;
  description: string;
  category: string;
  tags: string[];
  cover: string;
  isTop: boolean;
  isPublish: boolean;
}

export interface PostsVO {
  id: string;
  created_at: string;
  updated_at: string;
  title: string;
  content: string;
  author: string;
  description: string;
  category: string;
  tags: string[];
  cover: string;
  is_top: boolean;
  is_publish: boolean;
}

export interface PostsUpdateStatusReq {
  id: string;
  is_publish: boolean;
}
