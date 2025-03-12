"use client";
import {Spinner} from "@heroui/react";
export default function Loading() {
  return (
    <div className="flex justify-center items-center w-full h-[calc(100vh-4em)]">
      <Spinner classNames={{label: "text-foreground mt-4"}} label="加载中..." variant="wave" />
    </div>
  );
}     





