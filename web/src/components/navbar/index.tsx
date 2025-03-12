"use client";
import {
  Navbar,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
  Input,
} from "@heroui/react";
import { SearchIconProps } from "@/types/icons";
import Logo from "./logo";
import ThemeSwitcher from "./switcher";
import Link from "next/link";

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
  return (
    <Navbar
      className="
        bg-white/80 dark:bg-gray-900/80 
        hover:bg-white dark:hover:bg-gray-800 
        backdrop-blur-md 
        transition-all duration-300
        border-gray-300 dark:border-gray-700"
      isBordered
      maxWidth="2xl"
    >
      <NavbarBrand>
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
        <Input
          classNames={{
            base: "max-w-full sm:max-w-[15rem] h-10",
            mainWrapper: "h-full",
            input: "text-md",
            inputWrapper:
              "h-full font-normal text-default-500 bg-default-400/20 dark:bg-default-500/20",
          }}
          placeholder="搜索文章..."
          size="md"
          startContent={<SearchIcon size={18} width={18} height={18} />}
          type="search"
        />
        <NavbarItem>
          <ThemeSwitcher />
        </NavbarItem>
      </NavbarContent>
    </Navbar>
  );
}
