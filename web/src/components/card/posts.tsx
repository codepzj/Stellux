"use client";
import {
  Card,
  CardContent,
  CardDescription,
  CardTitle,
} from "@/components/ui/card";
import Image from "next/image";
import { Separator } from "@/components/ui/separator";
import { Badge } from "@/components/ui/badge";
import { FolderOpen, Tag } from "lucide-react";

import { useRouter } from "next/navigation";

import { timeAgo } from "@/utils/time";
import { IPostCard } from "@/types/posts";

export default function PostsCard({ post }: { post: IPostCard }) {
  const router = useRouter();
  return (
    <Card
      className="w-full py-4 cursor-pointer transition-all duration-150 active:scale-[0.99] hover:shadow-lg dark:hover:shadow-sky-400/40 active:bg-gray-100 dark:active:bg-gray-800"
      onClick={() => {
        router.push(`/posts/${post.id}`);
      }}
    >
      <CardContent className="flex flex-row-reverse items-center gap-4">
        <div className="w-[108px] h-[84px] relative hidden md:block">
          <Image
            src={post.cover || ""}
            alt={post.title}
            fill
            className="border-1 border-gray-200 rounded-sm"
          />
        </div>
        <div className="flex-1 flex flex-col justify-between">
          <CardTitle className="text-xl font-bold">{post.title}</CardTitle>
          <CardDescription className="line-clamp-1 my-1">
            {post.description || post.content}
          </CardDescription>
          <div className="w-full flex justify-between items-center text-xs text-gray-500">
            <div className="flex h-5 items-center space-x-2.5">
              <span>{post.author}</span>
              <Separator orientation="vertical" />
              <span>{timeAgo(post.created_at)}</span>
            </div>
            <div className="flex items-center space-x-2 py-1">
              <Badge variant="secondary">
                <FolderOpen />
                {post.category}
              </Badge>
              {post.tags?.map(tag => (
                <Badge variant="secondary" key={tag}>
                  <Tag />
                  {tag}
                </Badge>
              ))}
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}