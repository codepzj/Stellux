"use client";
import { Skeleton, Divider } from "@heroui/react";

export default function PostPageSkeleton() {
  return (
    <div className="max-w-[850px] m-auto p-6 my-4 space-y-6">
      {/* 标题 Skeleton */}
      <Skeleton className="w-1/4 h-8 rounded-lg bg-default-300" />

      {/* 分割线 */}
      <Divider className="my-4" />

      {/* 时间 Skeleton */}
      <Skeleton className="w-1/4 h-3 rounded bg-default-200" />

      {/* 文章内容 Skeleton，模拟多行文本 */}
      <div className="space-y-3">
        <Skeleton className="w-2/5 h-3 rounded bg-default-300" />
        <Skeleton className="w-3/5 h-3 rounded bg-default-300" />
        <Skeleton className="w-4/5 h-3 rounded bg-default-300" />
        <Skeleton className="w-3/5 h-3 rounded bg-default-300" />
      </div>
    </div>
  );
}
