<template>
  <div class="flex justify-center">
    <a-upload-dragger
      v-model:file-list="fileList"
      :multiple="true"
      :beforeUpload="handleBeforeUpload"
      list-type="picture"
      class="w-3/5 h-[280px] inline-block"
    >
      <div class="flex flex-col items-center justify-center h-full">
        <p class="ant-upload-drag-icon">
          <inbox-outlined></inbox-outlined>
        </p>
        <p class="ant-upload-text hidden md:block">
          点击或拖拽文件到此区域上传
        </p>
        <p class="ant-upload-hint hidden md:block">支持单个或批量上传</p>
      </div>
    </a-upload-dragger>
  </div>
</template>

<script lang="ts" setup>
import { InboxOutlined } from "@ant-design/icons-vue";
import { ref } from "vue";
import type { UploadProps } from "ant-design-vue/lib/upload/interface";
import { uploadPicturesLocal } from "@/api/modules/upload";
import { message } from "ant-design-vue";

const fileList = ref<UploadProps["fileList"]>([]);

const key = "upload-pictures-local";
const handleBeforeUpload: UploadProps["beforeUpload"] = async (
  file,
  FileList
) => {
  fileList.value = [...(fileList.value || []), file];
  if (fileList.value.length === FileList.length) {
    message.loading({
      content: "图片上传中...",
      key,
    });
    try {
      const res = await uploadPicturesLocal({
        files: fileList.value,
      });
      message.success({
        content: res.msg,
        key,
        duration: 1.5,
      });
      setTimeout(() => {
        fileList.value = [];
      }, 1500);
    } catch (error: any) {
      message.error({
        content: error + "，图片上传失败",
        key,
      });
    }
  }

  return false;
};
</script>
<style scoped>
.ant-upload-list-item {
  transition: opacity 0.5s ease-in-out;
}
</style>
