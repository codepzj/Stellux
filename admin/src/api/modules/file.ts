import request from "@/utils/request";
import type { PageResponse, Response } from "@/api/interfaces/resp";
import type { UploadFile } from "ant-design-vue/lib/upload/interface";
import type { IFile } from "@/api/interfaces/file";
export const uploadPicturesLocal: (data: {
  files: UploadFile[];
}) => Promise<Response<any>> = ({ files }: { files: UploadFile[] }) => {
  const formData = new FormData();
  files.forEach((file: any) => {
    formData.append("uids", file.uid);
    formData.append("file_names", file.name);
    formData.append("files", file);
  });
  return request.post("/picture/local/upload", formData);
};

export const getFilesByPage: (data: {
  page_no?: number;
  size?: number;
}) => Promise<PageResponse<IFile>> = ({
  page_no,
  size,
}: {
  page_no?: number;
  size?: number;
}) => {
  return request.get("/picture/list", { params: { page_no, size } });
};

export const deletePhotoByUid: (data: {
  uid: string;
}) => Promise<Response<any>> = ({ uid }: { uid: string }) => {
  return request.delete("/picture/local/delete", { params: { uid } });
};
