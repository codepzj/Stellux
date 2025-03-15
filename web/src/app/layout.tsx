import { Providers } from "./providers";
import NavBar from "@/components/navbar";

// 引入全局样式
import "@/styles/main.scss";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html suppressHydrationWarning lang="zh-CN">
      <body>
        <Providers themeProps={{ attribute: "class", defaultTheme: "light" }}>
          <NavBar />
          {children}
        </Providers>
      </body>
    </html>
  );
}
