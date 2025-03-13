import request from "@/utils/request";
import StelluxMarkdown from "@/components/markdown/md";
import { PostVO } from "@/types/posts";
import { Divider } from "@heroui/divider";
import dayjs from "dayjs";

export default async function PostPage({
  params,
}: {
  params: Promise<{ id: string }>;
  }) {
  const { id } = await params;
  const post = await request.get<PostVO>(`/posts/${id}`);
  const { data } = post;
  return (
    <div className="max-w-[850px] m-auto p-6 my-4">
      <h1 className="text-4xl font-bold mt-4">{data.title}</h1>
      <Divider className="my-4" />
      <div className="flex justify-between">
        <span className="text-sm text-gray-500 dark:text-slate-400">
          {dayjs(data.created_at).format("YYYY-MM-DD")}
        </span>
      </div>

      <StelluxMarkdown content={data.content} />
    </div>
  );
}
