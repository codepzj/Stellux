import { ThemeProvider } from "@/components/theme-provider";

import NavBar from "@/components/navbar";

// 引入全局样式
import "@/styles/global.scss";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html suppressHydrationWarning lang="zh-CN">
      <head>
        <link
          rel="stylesheet"
          href="https://cdn.jsdelivr.net/npm/cn-fontsource-lxgw-wen-kai-gb-screen/font.css"
        />
      </head>
      <body style={{ fontFamily: "LXGW WenKai GB Screen" }}>
        {/* 亮暗模式容器 */}
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <div className="flex flex-col h-screen">
            <NavBar />
            <div className="flex-1 mt-12">{children}</div>
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
