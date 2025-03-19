import Image from "next/image";
import EmptySvg from "@/assets/svg/empty.svg";

export default function Docs() {
  return (
    <div className="w-full h-[calc(100vh-10em)] flex flex-col justify-center items-center">
      <Image src={EmptySvg} alt="404" className="w-3/5 md:w-[20em]" />
      <div className="text-xl md:text-2xl font-bold">功能开发中...</div>
    </div>
  );
}
