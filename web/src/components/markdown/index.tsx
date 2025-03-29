"use client";
import MdSimple from "@/components/markdown/components/md-simple";
import MdFancy from "@/components/markdown/components/md-fancy";
import useMdThemeStore from "@/store/md-theme";
import { useMounted } from "@/hooks/use-mounted";

export default function Md({ content }: { content: string }) {
  const { mdTheme } = useMdThemeStore();
  const mounted = useMounted();
  if (mounted) {
    if (mdTheme === "react-md") {
      return <MdSimple content={content} />;
    } else {
      return <MdFancy content={content} />;
    }
  }
}
