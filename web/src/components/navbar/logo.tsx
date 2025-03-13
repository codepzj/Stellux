import lightLogo from "@/assets/logo/logo-light.png";
import darkLogo from "@/assets/logo/logo-dark.png";
import Image from "next/image";
import useThemeLoaded from "@/hooks/theme";
import { useRouter } from "next/navigation";
import { Skeleton } from "@heroui/react";

export default function Logo() {
  const isDark = useThemeLoaded();
  const router = useRouter();
  return isDark === undefined ? (
    <div className="max-w-[150px] w-full flex items-center transition-all ease-in-out duration-300">
      <div className="w-[86px] flex items-center justify-center">
        <Skeleton className="flex rounded-full w-9 h-9" />
      </div>
      <div className="w-full flex flex-col gap-2">
        <Skeleton className="h-2 w-3/5 rounded-lg" />
        <Skeleton className="h-2 w-4/5 rounded-lg" />
      </div>
    </div>
  ) : (
    <Image
      src={isDark ? darkLogo : lightLogo}
      alt="Stellux Logo"
      width={128}
      height={128}
      style={{ width: "auto", height: "auto" }}
      className="cursor-pointer"
      onClick={() => router.push("/")}
    />
  );
}
