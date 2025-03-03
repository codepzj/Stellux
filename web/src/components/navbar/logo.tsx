import lightLogo from "@/assets/logo/logo-light.png";
import darkLogo from "@/assets/logo/logo-dark.png";
import Image from "next/image";
import useThemeLoaded from "@/hooks/theme";

export default function Logo() {
  const isDark = useThemeLoaded();
  return isDark === undefined ? (
    <div className="h-8 w-8 animate-spin rounded-full border-b-2 border-t-2  border-gray-900 dark:border-gray-50"></div>
  ) : isDark ? (
    <Image src={darkLogo} alt="Stellux Logo" width={128} height={128} />
  ) : (
    <Image src={lightLogo} alt="Stellux Logo" width={128} height={128} />
  );
};
