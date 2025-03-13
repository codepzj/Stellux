"use client";
import { Pagination } from "@heroui/react";

export default function PaginationComponent({ total }: { total: number }) {
  return <Pagination initialPage={1} total={total} />;
}
