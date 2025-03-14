"use client";
import { useState } from "react";
import {
  Navbar,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
  Input,
  Button,
} from "@heroui/react";
import { SearchIconProps } from "@/types/icons";
import Logo from "./logo";
import Link from "next/link";
import Switcher from "@/components/navbar/switcher";
import { useRouter } from "next/navigation";

export const SearchIcon = ({
  size = 24,
  strokeWidth = 1.5,
  width,
  height,
  ...props
}: SearchIconProps) => {
  return (
    <svg
      aria-hidden="true"
      fill="none"
      focusable="false"
      height={height || size}
      role="presentation"
      viewBox="0 0 24 24"
      width={width || size}
      {...props}
    >
      <path
        d="M11.5 21C16.7467 21 21 16.7467 21 11.5C21 6.25329 16.7467 2 11.5 2C6.25329 2 2 6.25329 2 11.5C2 16.7467 6.25329 21 11.5 21Z"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={strokeWidth}
      />
      <path
        d="M22 22L20 20"
        stroke="currentColor"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={strokeWidth}
      />
    </svg>
  );
};

export default function NavBar() {
  const [isExpanded, setIsExpanded] = useState(false);
  const [searchValue, setSearchValue] = useState("");
  const router = useRouter();
  const handleSearch = () => {
    router.push(`/?keyword=${searchValue}`);
    setIsExpanded(false); // 搜索时，收缩搜索框
    setSearchValue(""); // 搜索时，清空搜索框
  };
  return (
    <Navbar
      className="w-full z-50
        bg-white/20 dark:bg-gray-900/20 
        hover:bg-white dark:hover:bg-gray-800 
        backdrop-blur-lg
        transition-all duration-300
        border-gray-300 dark:border-gray-700"
      isBordered
      maxWidth="2xl"
    >
      <NavbarBrand className="max-w-[150px] flex items-center justify-center">
        <Logo />
      </NavbarBrand>
      <NavbarContent className="hidden sm:flex gap-6" justify="start">
        <NavbarItem>
          <Link
            href="/"
            className="
              text-gray-700 dark:text-gray-200
              hover:text-blue-600 dark:hover:text-blue-400
              transition-colors duration-200
              font-medium"
          >
            首页
          </Link>
        </NavbarItem>
      </NavbarContent>
      <NavbarContent justify="end">
        <NavbarItem>
          <Input
            classNames={{
              base: `transition-all duration-300 ${
                isExpanded ? "w-64" : "w-40"
              }`,
              mainWrapper: "h-full",
              input: "text-sm",
              inputWrapper:
                "h-full font-normal text-gray-900 dark:text-gray-100 bg-gray-200 dark:bg-gray-700", // 修改颜色
            }}
            placeholder="搜索文章..."
            size="sm"
            radius="sm"
            startContent={<SearchIcon size={16} width={16} height={16} />}
            endContent={
              isExpanded && (
                <Button
                  color="primary"
                  radius="sm"
                  className="h-7 text-white text-sm absolute right-0"
                  onPress={handleSearch} // 按钮点击触发搜索
                >
                  搜索
                </Button>
              )
            }
            value={searchValue}
            onChange={(e) => setSearchValue(e.target.value)}
            onFocus={() => setIsExpanded(true)}
            onBlur={() => {
              if (!searchValue.trim()) {
                setIsExpanded(false);
              }
            }}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                handleSearch();
              }
            }}
            type="search"
          />
        </NavbarItem>
        <NavbarItem>
          <Switcher />
        </NavbarItem>
      </NavbarContent>
    </Navbar>
  );
}
