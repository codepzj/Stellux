<template>
  <a-button type="primary" @click="handleClick"> 选择图片 </a-button>
  <div v-if="isSelecting">
    <a-button type="primary" @click="handleSelectAll"> 全选 </a-button>
    <a-button type="primary" @click="handleSelectNone"> 取消全选 </a-button>
  </div>
  <div ref="photoWallRef" v-if="photosWall.length > 0">
    <a-upload
      v-model:file-list="photosWall"
      list-type="picture-card"
      @preview="handlePreview"
      @remove="handleRemove"
    />
  </div>
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
import type { IFile } from "@/types/file";
import { message, type UploadFile } from "ant-design-vue";
import { deletePhotoByUid } from "@/api/file";

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

const photoWallRef = ref<any>(null);
const pagination = reactive({
  current: props.current,
  total: props.total,
  size: props.size,
});
const baseURL = import.meta.env.VITE_API_URL;
const previewVisible = ref(false);
const previewImage = ref("");
const previewTitle = ref("");
const selectedPhotos = ref<string[]>([]);
const isSelecting = ref(false); // 新增状态：是否处于选择模式
const photosWall = computed(() => {
  return props.list.map(item => {
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

// 进入选择模式并添加复选框
const handleClick = () => {
  isSelecting.value = true; // 进入选择模式，显示全选/取消全选按钮
  const items = photoWallRef.value.querySelectorAll(".ant-upload-list-item");
  items.forEach((item: any) => {
    // 避免重复添加复选框
    if (!item.querySelector(".ant-checkbox")) {
      const checkbox = document.createElement("input");
      checkbox.type = "checkbox";
      checkbox.classList.add("ant-checkbox");
      checkbox.style.cssText = "position: absolute; top: 2px; right: 2px;";
      checkbox.addEventListener("change", () => {
        const url = item.querySelector(".ant-upload-list-item-thumbnail")?.href;
        if (checkbox.checked) {
          item.classList.add("selected");
          if (url && !selectedPhotos.value.includes(url)) {
            selectedPhotos.value.push(url);
          }
        } else {
          item.classList.remove("selected");
          selectedPhotos.value = selectedPhotos.value.filter(
            photo => photo !== url
          );
        }
        console.log(selectedPhotos.value);
      });
      item.appendChild(checkbox);
    }
  });
};

// 全选
const handleSelectAll = () => {
  const items = photoWallRef.value.querySelectorAll(".ant-upload-list-item");
  selectedPhotos.value = [];
  items.forEach((item: any) => {
    const checkbox = item.querySelector(".ant-checkbox") as HTMLInputElement;
    if (checkbox) {
      checkbox.checked = true;
      item.classList.add("selected");
      const url = item.querySelector(".ant-upload-list-item-thumbnail")?.href;
      if (url && !selectedPhotos.value.includes(url)) {
        selectedPhotos.value.push(url);
      }
    }
  });
};

// 取消全选
const handleSelectNone = () => {
  const items = photoWallRef.value.querySelectorAll(".ant-upload-list-item");
  items.forEach((item: HTMLElement) => {
    const checkbox = item.querySelector(".ant-checkbox") as HTMLInputElement;
    if (checkbox) {
      checkbox.checked = false;
      item.classList.remove("selected");
    }
  });
  selectedPhotos.value = [];
};

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

// 监听selectedPhotos
watch(selectedPhotos, () => {
  console.log(selectedPhotos.value);
});

function getBase64(file: File) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
}
</script>

<style lang="scss" scoped>
// 保持原有样式不变
:deep(.ant-upload-list.ant-upload-list-picture-card) {
  margin: 0 auto;
  width: 100%;
}
:deep(.ant-upload-list-item-actions) {
  display: flex;
  justify-content: center;
  align-items: center;
  a {
    display: flex;
    justify-content: center;
    align-items: center;
  }
}

:deep(.ant-upload-list-picture-card .ant-upload-list-item::before) {
  border-radius: 0 25px;
}

:deep(
  .ant-upload-wrapper.ant-upload-picture-card-wrapper
    .ant-upload-list-item.ant-upload-list-item-done.selected
) {
  border: 1px dashed #1890ff;
}

</style>
