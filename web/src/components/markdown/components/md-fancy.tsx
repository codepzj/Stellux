"use client";
import { useState } from "react";
import { MdPreview } from "md-editor-rt";
import "md-editor-rt/lib/preview.css";
import { useTheme } from "next-themes";
import { Themes } from "@/types/theme";
import useMdThemeStore from "@/store/md-theme";

export default function MdFancy({ content }: { content: string }) {
  const [id] = useState("preview-only");
  const { resolvedTheme } = useTheme();
  const { mdTheme } = useMdThemeStore();
  return (
    <>
      <MdPreview
        id={id}
        value={content}
        theme={resolvedTheme as Themes}
        previewTheme={mdTheme}
      />
    </>
  );
}
