"use client";
import { Card, CardBody, Image, Code } from "@heroui/react";
import { useRouter } from "next/navigation";
import { FolderOpen, Tags } from "lucide-react";

export default function PostsCard({
  id,
  title,
  description,
  category,
  tags,
  cover,
}: {
  id: string;
  title: string;
  description: string;
  category: string;
  tags: string[];
  cover?: string;
}) {
  const router = useRouter();
  return (
    <Card
      shadow="sm"
      className="w-4/5 md:w-[750px] h-40 mt-4 mx-4 px-2 cursor-pointer dark:bg-gray-900 dark:border dark:border-gray-700 will-change-transform duration-200"
      isHoverable
      isPressable
      onPress={() => router.push(`/posts/${id}`)}
    >
      <CardBody className="p-4">
        <div className="flex h-full gap-4 items-center">
          <div className="flex-1 space-y-3">
            <h1 className="text-xl font-semibold text-gray-900 dark:text-white line-clamp-1">
              {title}
            </h1>
            <p className="text-sm text-gray-600 dark:text-gray-300 line-clamp-2">
              {description}
            </p>
            <div className="flex space-x-2 text-xs text-gray-500 dark:text-gray-400">
              <Code className="font-sans">
                <div className="flex items-center">
                  <FolderOpen size={16} className="mr-1" />
                  {category}
                </div>
              </Code>

              <div className="flex items-center">
                {tags.map((tag) => (
                  <Code className="font-sans mr-1" key={tag}>
                    <div className="flex items-center">
                      <Tags size={16} className="mr-1" />
                      <span>{tag}</span>
                    </div>
                  </Code>
                ))}
              </div>
            </div>
          </div>
          <div className="w-[100px] h-[100px] overflow-hidden rounded-xl">
            <Image
              src={cover || "/placeholder-image.jpg"}
              width={100}
              height={100}
              alt={title}
              className="object-cover w-full h-full hover:scale-110"
            />
          </div>
        </div>
      </CardBody>
    </Card>
  );
}
