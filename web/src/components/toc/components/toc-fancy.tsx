"use client";
import { useState } from "react";
import { MdCatalog } from "md-editor-rt";
import "md-editor-rt/lib/preview.css";

export default function TocFancy() {
  const [id] = useState("preview-only");
  const [scrollElement] = useState(document.documentElement);

  return (
    <>
      <MdCatalog editorId={id} scrollElement={scrollElement} />
    </>
  );
}
