import type { Config } from "tailwindcss";
import { heroui } from "@heroui/react";

export default {
  content: [
    // "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    // "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./node_modules/@heroui/theme/dist/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  darkMode: "class",
  plugins: [
    heroui({
      defaultTheme: "light",
      defaultExtendTheme: "light",
    }),
  ],
} satisfies Config;
