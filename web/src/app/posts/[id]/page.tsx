import request from "@/utils/request";
import StelluxMarkdown from "@/app/posts/components/markdown/md";

interface PostVO {
  id: string;
  title: string;
  content: string;
  created_at: string;
  updated_at: string;
}

export default async function PostPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;
  const post = await request.get<PostVO>(`/posts/${id}`);
  const { data } = post;
  return (
    <div className="w-[800px] m-auto">
      <h1 className="text-2xl font-bold">{data.title}</h1>
      <div className="flex justify-between">
        <span className="text-gray-500">{data.created_at}</span>
        <span className="text-gray-500">{data.updated_at}</span>
      </div>
      <StelluxMarkdown content={data.content} />
    </div>
  );
}
