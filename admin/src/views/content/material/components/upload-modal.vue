<template>
  <div class="w-full flex flex-row justify-end my-2 pr-8">
    <a-button @click="showModal"> 上传附件 </a-button>
  </div>
  <a-modal
    v-model:open="open"
    title="附件上传"
    :ok-text="fileList?.length === 0 ? '上传' : '确认上传'"
    :cancel-text="fileList?.length === 0 ? '关闭' : '取消'"
    @cancel="handleModalCancel"
    @ok="handleModalOk"
    width="100%"
    wrap-class-name="full-modal"
  >
    <div class="flex flex-col items-center">
      <a-tabs v-model:activeKey="activeKey">
        <a-tab-pane key="1" tab="文件上传"></a-tab-pane>
        <a-tab-pane key="2" tab="文件夹上传"></a-tab-pane>
      </a-tabs>
      <a-upload-dragger
        v-model:file-list="fileList"
        :multiple="true"
        :beforeUpload="handleBeforeUpload"
        list-type="picture"
        :directory="activeKey === '2'"
        class="w-3/5 h-[280px] inline-block"
        @remove="handleRemove"
      >
        <div class="flex flex-col items-center justify-center h-full">
          <p class="ant-upload-drag-icon">
            <inbox-outlined></inbox-outlined>
          </p>
          <p class="ant-upload-text hidden md:block">
            点击或拖拽{{ activeKey === "1" ? "文件" : "文件夹" }}到此区域上传
          </p>
          <p class="ant-upload-hint hidden md:block">支持单个或批量上传</p>
        </div>
      </a-upload-dragger>
    </div>
  </a-modal>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { deletePhotoByUid, uploadPicturesLocal } from "@/api/modules/file";
import { message, type UploadFile } from "ant-design-vue";
import { InboxOutlined } from "@ant-design/icons-vue";
import type { UploadProps } from "ant-design-vue/lib/upload/interface";

// 父组件重新获取图片列表
const emit = defineEmits(["update:list"]);
const open = ref<boolean>(false);
const key = "upload-pictures-local";
const fileList = ref<UploadProps["fileList"]>([]);
const activeKey = ref("1");

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
      });
    } catch (error: any) {
      message.error({
        content: error + "，图片上传失败",
        key,
      });
    }
  }
  return false;
};

const showModal = () => {
  open.value = true;
};

const handleModalCancel = () => {
  // 如果文件列表为空，则关闭弹窗
  if (fileList.value?.length === 0) {
    open.value = false;
    fileList.value = [];
    emit("update:list");
  }
  // 如果文件列表不为空，则提示用户确认上传
  else {
    // TODO: 回滚上传的文件
    message.success("取消上传成功");
    open.value = false;
    fileList.value = [];
  }
};

const handleModalOk = () => {
  if (fileList.value?.length === 0) {
    message.error("请上传文件");
    return;
  }
  // 清空文件列表
  fileList.value = [];
  message.success("确认上传成功");
  emit("update:list");
};

// 删除已上传的特定文件
const handleRemove = async (file: UploadFile) => {
  try {
    const res = await deletePhotoByUid({ uid: file.uid });
    message.success(res.msg);
    emit("update:list");
  } catch (error: any) {
    message.error(error);
    return false;
  }
};
</script>

<style lang="scss">
.full-modal {
  .ant-modal {
    max-width: 100%;
    top: 0;
    padding-bottom: 0;
    margin: 0;
  }
  .ant-modal-content {
    display: flex;
    flex-direction: column;
    height: calc(100vh);
  }
  .ant-modal-body {
    flex: 1;
    overflow-y: auto;
  }
}

::-webkit-scrollbar {
  width: 8px;
}
</style>
