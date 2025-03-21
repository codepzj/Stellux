// @ts-nocheck
"use client"

import * as React from "react"

import { TableOfContents } from "@/lib/toc"
import { cn } from "@/lib/utils"
import { useMounted } from "@/hooks/use-mounted"

interface TocProps {
  toc: TableOfContents
}

export function DashboardTableOfContents({ toc }: TocProps) {
  const itemIds = React.useMemo(
    () =>
      toc.items
        ? toc.items
          .flatMap((item) => [item.url, item?.items?.map((item) => item.url)])
          .flat()
          .filter(Boolean)
          .map((id) => id?.split("#")[1])
        : [],
    [toc]
  )
  const activeHeading = useActiveItem(itemIds)
  const mounted = useMounted()

  if (!toc?.items?.length) {
    return null
  }

  return (
    <nav className="space-y-4 rounded-lg border p-4">
      <p className="text-lg font-semibold">目录</p>
      <Tree tree={toc} activeItem={activeHeading} />
    </nav>
  )
}

function useActiveItem(itemIds: string[]) {
  const [activeId, setActiveId] = React.useState(null)

  React.useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            setActiveId(entry.target.id)
          }
        })
      },
      { rootMargin: `0% 0% -80% 0%` }
    )

    itemIds?.forEach((id) => {
      const element = document.getElementById(id)
      if (element) {
        observer.observe(element)
      }
    })

    return () => {
      itemIds?.forEach((id) => {
        const element = document.getElementById(id)
        if (element) {
          observer.unobserve(element)
        }
      })
    }
  }, [itemIds])

  return activeId
}

interface TreeProps {
  tree: TableOfContents
  level?: number
  activeItem?: string
}

function Tree({ tree, level = 1, activeItem }: TreeProps) {
  return tree?.items?.length && level < 3 ? (
    <ul className={cn("m-0 list-none space-y-2", { "pl-2 border-l border-muted": level !== 1 })}>
      {tree.items.map((item, index) => {
        const isActiveParent = item.items?.some((subItem) => `#${activeItem}` === subItem.url)
        return (
          <li key={index} className="relative group pl-2">
            <a
              href={item.url}
              className={cn(
                "block transition-colors duration-200",
                "hover:text-primary", // 悬停时高亮
                item.url === `#${activeItem}`
                  ? "font-semibold text-primary" // 选中时高亮
                  : "text-muted-foreground"
              )}
            >
              {item.title}
            </a>
            {item.items?.length ? (
              <ul className={cn("pl-4 border-l border-muted", { "block": isActiveParent, "hidden group-hover:block": !isActiveParent })}>
                <Tree tree={item} level={level + 1} activeItem={activeItem} />
              </ul>
            ) : null}
          </li>
        )
      })}
    </ul>
  ) : null
}