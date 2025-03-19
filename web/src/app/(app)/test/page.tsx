import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

import { Menu } from "lucide-react";
import Link from "next/link";

export default function DropdownMenuDemo() {
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
        <DropdownMenuItem>
          <Link href="/">🏠首页</Link>
        </DropdownMenuItem>
        <DropdownMenuItem>
          <Link href="/docs">📝文档</Link>
        </DropdownMenuItem>
        <DropdownMenuItem>
          <Link href="/notes">📒便签</Link>
        </DropdownMenuItem>
        <DropdownMenuItem>
          <Link href="/about">👋关于</Link>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
