"use client";
import { Skeleton } from "@heroui/react";

export default function App() {

  return (
    <div className="max-w-[800px] m-auto p-4 my-4 space-y-3">
      <div className="space-y-3">
        <Skeleton className="w-full rounded-lg">
          <div className="h-3 w-full rounded-lg bg-default-200" />
        </Skeleton>
        <Skeleton className="w-full rounded-lg">
          <div className="h-3 w-full rounded-lg bg-default-200" />
        </Skeleton>
        <Skeleton className="w-full rounded-lg">
          <div className="h-3 w-full rounded-lg bg-default-300" />
        </Skeleton>
        <Skeleton className="w-full rounded-lg">
          <div className="h-3 w-full rounded-lg bg-default-300" />
        </Skeleton>
      </div>
    </div>
  );
}
