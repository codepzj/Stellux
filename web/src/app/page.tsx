import PostsCard from "@/components/card/posts";

import { getPostsList } from "@/api/posts";

export default async function App() {
  const res = await getPostsList({ page_no: 1, size: 10 });
  const PostsList = res.data.list;
  return (
    <div className="display flex flex-col items-center">
      {PostsList.map((posts) => (
        <PostsCard
          id={posts.id}
          title={posts.title}
          description={posts.description}
          category={posts.category}
          tags={posts.tags}
          cover={posts.cover}
          key={posts.id}
        ></PostsCard>
      ))}
    </div>
  );
}
