<template>
  <a-upload
    v-if="photosWall.length > 0"
    v-model:file-list="photosWall"
    list-type="picture-card"
    @preview="handlePreview"
    @remove="handleRemove"
  />
  <a-empty v-else />

  <div class="flex flex-row justify-end">
    <a-pagination
      v-model:current="pagination.current"
      v-model:pageSize="pagination.size"
      :total="pagination.total"
      @change="handlePageChange"
    />
  </div>
  <a-modal
    :open="previewVisible"
    :title="previewTitle"
    :footer="null"
    @cancel="handleCancel"
  >
    <img alt="example" style="width: 100%" :src="previewImage" />
  </a-modal>
</template>

<script lang="ts" setup>
import { computed, reactive, ref, watch } from "vue";
import type { IFile } from "@/api/interfaces/file";
import { message, type UploadFile } from "ant-design-vue";
import { deletePhotoByUid } from "@/api/modules/file";

const props = defineProps<{
  current: number;
  total: number;
  size: number;
  list: IFile[];
}>();

const emit = defineEmits<{
  (e: "update:pagination", pagination: { current: number; size: number }): void;
  (e: "update:list", deleteItem: UploadFile): void;
}>();

const pagination = reactive({
  current: props.current,
  total: props.total,
  size: props.size,
});

const baseURL = import.meta.env.VITE_API_URL;

const previewVisible = ref(false);
const previewImage = ref("");
const previewTitle = ref("");

const photosWall = computed(() => {
  return props.list.map((item) => {
    const imageUrl = baseURL + item.url;
    const reg = /^https?:\/\/.*\.(jpg|jpeg|png|gif|bmp|webp|svg|ico|avif)$/i;
    if (reg.test(imageUrl)) {
      return {
        uid: item.uid,
        name: item.name,
        status: "done",
        url: imageUrl,
        thumbUrl: imageUrl,
      };
    }
    return {
      uid: item.id,
      name: "404",
      status: "error",
    };
  });
});

function getBase64(file: File) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = (error) => reject(error);
  });
}

// 预览图片
interface FileType {
  url?: string;
  preview?: string;
  originFileObj?: File;
  name?: string;
}
const handlePreview = async (file: UploadFile<FileType>) => {
  if (!file.url && !file.preview) {
    file.preview = (await getBase64(file.originFileObj as File)) as string;
  }
  previewImage.value = file.url || file.preview || "";
  previewVisible.value = true;
  previewTitle.value =
    file.name ||
    (file.url ? file.url.substring(file.url.lastIndexOf("/") + 1) : "");
};

// 取消预览
const handleCancel = () => {
  previewVisible.value = false;
  previewTitle.value = "";
};

// 删除图片
const handleRemove = async (file: UploadFile) => {
  try {
    const res = await deletePhotoByUid({ uid: file.uid });
    // 回调成功删除特定图片，避免重复获取图片列表
    emit("update:list", file);
    message.success(res.msg);
  } catch (error: any) {
    message.error(error);
  }
};

// 分页
const handlePageChange = (current: number, size: number) => {
  emit("update:pagination", { current, size });
};

// 监听父组件传入的参数
watch(props, () => {
  pagination.current = props.current;
  pagination.total = props.total;
  pagination.size = props.size;
});
</script>

<style lang="scss" scoped>
:deep(.ant-upload-list.ant-upload-list-picture-card) {
  margin: 0 auto;
  width: 100%;
}
:deep(.ant-upload-list-item-actions) {
  display: flex;
  justify-content: center;
  align-items: center;
  // 让内部a标签的文字垂直居中
  a {
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>
