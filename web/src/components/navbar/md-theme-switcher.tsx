"use client";

import * as React from "react";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Highlighter } from "lucide-react";
import useMdThemeStore from "@/store/md-theme";
export default function ModeToggle() {
  const { setMdTheme } = useMdThemeStore();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="secondary"
          size="icon"
          className="focus-visible:ring-0 focus-visible:border-none hover:bg-gray-100 dark:hover:bg-gray-800 cursor-pointer"
        >
          <Highlighter />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="start">
        <DropdownMenuLabel>主题</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem onClick={() => setMdTheme("react-md")}>
            react-md
          </DropdownMenuItem>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem onClick={() => setMdTheme("default")}>
            default
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => setMdTheme("github")}>
            github
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => setMdTheme("vuepress")}>
            vuepress
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => setMdTheme("mk-cute")}>
            mk-cute
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => setMdTheme("smart-blue")}>
            smart-blue
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => setMdTheme("cyanosis")}>
            cyanosis
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => setMdTheme("arknights")}>
            arknights
          </DropdownMenuItem>
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
