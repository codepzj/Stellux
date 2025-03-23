"use client";
import { useRouter } from "next/navigation";
import { useTheme } from "next-themes";
import Image from "next/image";
import LightLogo from "@/assets/logo/logo-light.png";
import DarkLogo from "@/assets/logo/logo-dark.png";
import { useMounted } from "@/hooks/use-mounted";

export default function Logo() {
  const router = useRouter();
  const { resolvedTheme } = useTheme();
  const mounted = useMounted();
  return (
    mounted && (
      <Image
        src={resolvedTheme === "dark" ? DarkLogo : LightLogo}
        alt="logo"
        width={128}
        height={128}
        className="mr-4 flex items-center cursor-pointer lg:mr-6 hover:scale-95"
        onClick={() => router.push("/")}
      />
    )
  );
}
