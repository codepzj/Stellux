import request from "@/utils/request";
import type { PageResponse, Response } from "@/types/response";
import type { UploadFile } from "ant-design-vue/lib/upload/interface";
import type { IFile } from "@/types/file";

const ALLOWED_IMAGE_TYPES = [
  "image/jpeg",
  "image/png",
  "image/gif",
  "image/bmp",
  "image/webp",
] as const;

const isValidImageType = (type: string): boolean => {
  return ALLOWED_IMAGE_TYPES.includes(type as any);
};

const cropTo16By9 = (file: File): Promise<Blob> => {
  return new Promise((resolve, reject) => {
    const img = new Image();
    img.onload = () => {
      const aspectRatio = 16 / 9;
      const canvas = document.createElement("canvas");
      const ctx = canvas.getContext("2d");

      let targetWidth = img.width;
      let targetHeight = img.height;

      if (targetWidth / targetHeight > aspectRatio) {
        targetWidth = targetHeight * aspectRatio;
      } else {
        targetHeight = targetWidth / aspectRatio;
      }

      canvas.width = targetWidth;
      canvas.height = targetHeight;

      if (ctx) {
        ctx.clearRect(0, 0, targetWidth, targetHeight); // 清除画布，保持透明背景

        const scale = Math.min(
          targetWidth / img.width,
          targetHeight / img.height
        );
        const drawWidth = img.width * scale;
        const drawHeight = img.height * scale;
        const offsetX = (targetWidth - drawWidth) / 2;
        const offsetY = (targetHeight - drawHeight) / 2;

        ctx.drawImage(img, offsetX, offsetY, drawWidth, drawHeight);
      }

      canvas.toBlob(
        blob =>
          blob ? resolve(blob) : reject(new Error("Canvas 转换 Blob 失败")),
        "image/png", // 使用 PNG 以支持透明背景
        1
      );
    };

    img.onerror = () => reject(new Error("图片加载失败"));
    img.src = URL.createObjectURL(file);
  });
};

export const uploadPicturesLocal = async ({
  files,
}: {
  files: UploadFile[];
}): Promise<Response<any>> => {
  const formData = new FormData();

  for (const file of files) {
    try {
      if (!file.type || !isValidImageType(file.type)) {
        throw new Error(
          `不支持的文件类型: ${file.type || "未知类型"}。仅支持: ${ALLOWED_IMAGE_TYPES.join(", ")}`
        );
      }

      const croppedBlob = await cropTo16By9(file as unknown as File);
      formData.append("uids", file.uid);
      formData.append("file_names", file.name);
      formData.append("files", croppedBlob, file.name);
    } catch (error) {
      console.error("文件处理失败:", file.name, error);
      throw error;
    }
  }

  return request.post("/picture/local/upload", formData);
};

export const getFilesByPage = ({
  page_no,
  size,
}: {
  page_no?: number;
  size?: number;
}): Promise<PageResponse<IFile>> => {
  return request.get("/picture/list", { params: { page_no, size } });
};

export const deletePhotoByUid = ({
  uid,
}: {
  uid: string;
}): Promise<Response<any>> => {
  return request.delete("/picture/local/delete", { params: { uid } });
};
