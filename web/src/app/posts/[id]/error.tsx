"use client"
import { useEffect } from "react"

export default function Error({
     error,
     reset
}: {
     error: Error & { digest?: string },
     reset: () => void
}) {

     useEffect(() => {
         console.error(error)
     },[error])

    return (
        <div>
            文章不存在呢
        </div>
    );
}