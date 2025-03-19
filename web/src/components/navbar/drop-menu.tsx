"use client";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuGroup,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

import { Menu } from "lucide-react";
import { useRouter } from "next/navigation";

export default function DropMenu() {
  const router = useRouter();
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="secondary"
          size="icon"
          className="focus-visible:ring-0 focus-visible:border-none hover:bg-gray-100 dark:hover:bg-gray-800 cursor-pointer"
        >
          <Menu />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-16 cursor-pointer">
        <DropdownMenuGroup>
          <DropdownMenuItem onClick={() => router.push("/")}>
            🏠首页
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => router.push("/docs")}>
            📝文档
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => router.push("/notes")}>
            📒便签
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => router.push("/about")}>
            👋关于
          </DropdownMenuItem>
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
