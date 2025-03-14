"use client";
import { Pagination } from "@heroui/react";
import { useRouter } from "next/navigation";
import qs from "qs";

export default function PaginationComponent({
  total,
  searchParams,
}: {
  total: number;
  searchParams: { page_no?: number; size?: number; keyword?: string };
  }) {
  const router = useRouter();
  // 获取服务端路径，按钮点击后，更新分页
  const handlePageChange = (
    number: number,
    searchParams: { page_no?: number; size?: number }
  ) => {
    searchParams.page_no = number;
    router.push(`/?${qs.stringify(searchParams)}`);
  };
  return (
    <Pagination
      onChange={(number) => handlePageChange(number, searchParams)}
      loop
      showControls
      initialPage={1}
      total={total}
    />
  );
}
