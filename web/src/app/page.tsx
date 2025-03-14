import PostsCard from "@/components/card/posts";
import PaginationComponent from "@/components/pagination";
import { getPostsList } from "@/api/posts";

import { stringToNumber } from "@/utils/convert";
import Image from "next/image";
import EmptySvg from "@/assets/svg/empty.svg";
export default async function App({
  searchParams,
}: {
  searchParams?: Promise<{ page_no?: string; size?: string; keyword?: string }>;
  }) {
  // 获取搜索参数
  const { page_no, size, keyword } = (await searchParams) || {};
  // 将搜索参数转换为数字
  const page_no_number = stringToNumber(page_no || "1");
  const size_number = stringToNumber(size || "10");
  // 获取文章列表
  const res = await getPostsList({
    page_no: page_no_number,
    size: size_number,
    keyword,
  });
  // 获取文章列表
  const PostsList = res.data.list;
  return res.data.total_count > 0 ? (
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
      <div className="w-full flex justify-center md:justify-end items-center my-4">
        <PaginationComponent
          total={res.data.total_page}
          searchParams={{ page_no: page_no_number, size: size_number }}
        />
      </div>
    </div>
  ) : (
    <div className="absolute w-screen h-[calc(100vh-6em)] z-[999]">
      <div className="w-full h-[73vh] mt-20">
        <div className="w-full h-full flex justify-center items-center flex-col">
          <Image
            src={EmptySvg}
            alt="404"
            className="w-3/5 xl:w-[28rem] lg:w-[24rem] md:w-[20rem]"
          />

          <div className="xl:w-[32rem] lg:w-[26rem] md:w-[20rem] text-center mx-4">
            <h2 className="text-2xl sm:text-xl font-bold my-4">
              未搜索到“{keyword}”相关内容
            </h2>
          </div>
        </div>
      </div>
    </div>
  );
}
