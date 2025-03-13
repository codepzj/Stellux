import PostsCard from "@/components/card/posts";
import PaginationComponent from "@/components/pagination";
import { getPostsList } from "@/api/posts";

export default async function App() {
  const res = await getPostsList({ page_no: 1, size: 10 });
  const PostsList = res.data.list;
  return (
    <div className="w-4/5 md:max-w-[800px] mx-auto display flex flex-col items-center">
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
      <div className="w-full flex justify-end items-center my-4">
        <PaginationComponent total={res.data.total_page} />
      </div>
    </div>
  );
}
