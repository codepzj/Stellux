"use client";

import TocSimple from "@/components/toc/components/toc-simple";
import TocFancy from "@/components/toc/components/toc-fancy";
import useMdThemeStore from "@/store/md-theme";
import { useMounted } from "@/hooks/use-mounted";
import { TableOfContents } from "@/lib/toc";

export default function Toc({ toc }: { toc: TableOfContents }) {
  const { mdTheme } = useMdThemeStore();
  const mounted = useMounted();
  if (mounted) {
    if (mdTheme === "react-md") {
      return <TocSimple toc={toc} />;
    } else {
      return <TocFancy />;
    }
  }
}
