<template>
  <div>
    <div class="w-full flex justify-between items-center my-4">
      <div class="flex gap-2" v-if="props.type === 'display'">
        <a-button type="primary" @click="handleSelectAll">全选</a-button>
        <a-button type="primary" @click="handleSelectNone">取消全选</a-button>
        <slot name="action1"></slot>
      </div>
      <div class="flex gap-2">
        <slot name="action2"></slot>
      </div>
    </div>

    <div ref="photoWallRef" v-if="photosWall.length">
      <a-upload
        v-model:file-list="photosWall"
        :list-type="props.mode"
        @preview="handlePreview"
      />
    </div>
    <a-empty v-else />

    <a-modal
      :open="previewVisible"
      :title="previewTitle"
      :footer="null"
      @cancel="handleCancel"
    >
      <img
        :alt="previewTitle"
        class="w-full max-w-[800px] max-h-[600px] mx-auto"
        :src="previewImage"
      />
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref, watch, onMounted, nextTick } from "vue";
import { type UploadFile } from "ant-design-vue";

const props = defineProps<{
  mode: "picture" | "picture-card";
  type: "display" | "select";
  list: { uid: string; name: string; url: string; id?: string }[];
}>();

// 更新父组件选中的图片
const emit = defineEmits<{
  (e: "selected-picture", photos: string[]): void;
}>();

const photoWallRef = ref<HTMLElement | null>(null);
const baseURL = import.meta.env.VITE_API_URL;
const previewVisible = ref(false);
const previewImage = ref("");
const previewTitle = ref("");
const selectedPhotos = ref<string[]>([]);

const photosWall = computed(() =>
  props.list.map(item => ({
    uid: item.uid,
    name: item.name,
    status: /\.(jpg|jpeg|png|gif|bmp|webp|svg|ico|avif)$/i.test(
      baseURL + item.url
    )
      ? "done"
      : "error",
    url: baseURL + item.url,
    thumbUrl: baseURL + item.url,
  }))
);

// 确保 DOM 渲染完成后添加选择框
onMounted(async () => {
  await nextTick();
  addSelectionControls();
});

watch(photosWall, async () => {
  await nextTick();
  addSelectionControls();
});

// 添加复选框或单选框
const addSelectionControls = () => {
  if (!photoWallRef.value) return;
  const items = Array.from(
    photoWallRef.value.querySelectorAll(".ant-upload-list-item")
  ) as HTMLElement[];

  items.forEach(item => {
    if (item.querySelector(".selection-control")) return;

    const input = document.createElement("input");
    input.className = "selection-control";
    input.style.cssText =
      "position: absolute; top: 2px; right: 2px; z-index: 10;";

    if (props.type === "display") {
      input.type = "checkbox";
      input.onchange = () => toggleMultipleSelection(item, input.checked);
    } else {
      input.type = "radio";
      input.name = "photo-selection";
      input.onchange = () => toggleSingleSelection(item);
    }
    item.appendChild(input);
  });
};

// 处理多选
const toggleMultipleSelection = (item: HTMLElement, checked: boolean) => {
  const url =
    item
      .querySelector(".ant-upload-list-item-thumbnail")
      ?.getAttribute("href") || "";
  item.classList.toggle("selected", checked);

  selectedPhotos.value = checked
    ? [...selectedPhotos.value, url]
    : selectedPhotos.value.filter(photo => photo !== url);
};

// 处理单选
const toggleSingleSelection = (item: HTMLElement) => {
  const url =
    item
      .querySelector(".ant-upload-list-item-thumbnail")
      ?.getAttribute("href") || "";

  // 清空所有选中的
  const items = Array.from(
    photoWallRef.value?.querySelectorAll(".ant-upload-list-item") || []
  ) as HTMLElement[];
  items.forEach(otherItem => {
    otherItem.classList.remove("selected");
    const input = otherItem.querySelector(
      ".selection-control"
    ) as HTMLInputElement;
    if (input) input.checked = false;
  });

  // 选中当前项
  item.classList.add("selected");
  const input = item.querySelector(".selection-control") as HTMLInputElement;
  if (input) input.checked = true;

  selectedPhotos.value = [url];
};

// 全选
const handleSelectAll = () => {
  if (props.type !== "display") return;
  const items = Array.from(
    photoWallRef.value?.querySelectorAll(".ant-upload-list-item") || []
  ) as HTMLElement[];

  selectedPhotos.value = items.map(item => {
    const checkbox = item.querySelector(
      ".selection-control"
    ) as HTMLInputElement;
    checkbox.checked = true;
    item.classList.add("selected");
    return (
      item
        .querySelector(".ant-upload-list-item-thumbnail")
        ?.getAttribute("href") || ""
    );
  });
};

// 取消全选
const handleSelectNone = () => {
  if (props.type !== "display") return;
  const items = Array.from(
    photoWallRef.value?.querySelectorAll(".ant-upload-list-item") || []
  ) as HTMLElement[];

  items.forEach(item => {
    const checkbox = item.querySelector(
      ".selection-control"
    ) as HTMLInputElement;
    checkbox.checked = false;
    item.classList.remove("selected");
  });
  selectedPhotos.value = [];
};

// 预览图片
const handlePreview = async (file: UploadFile) => {
  previewImage.value =
    file.url || (await getBase64(file.originFileObj as File)) || "";
  previewVisible.value = true;
  previewTitle.value = file.name || file.url?.split("/").pop() || "";
};

// 关闭预览
const handleCancel = () => {
  previewVisible.value = false;
  previewTitle.value = "";
};

// 监听选中照片
watch(
  selectedPhotos,
  () => {
    emit("selected-picture", selectedPhotos.value);
    console.log(selectedPhotos.value);
  },
  { deep: true }
);

// 获取 Base64 格式图片
const getBase64 = (file: File) =>
  new Promise<string>((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result as string);
    reader.onerror = reject;
  });
</script>

<style lang="scss" scoped>
:deep(.ant-upload-list-picture-card) {
  margin: 0 auto;
  width: 100%;
}

:deep(.ant-upload-list-item-actions) {
  display: flex;
  justify-content: center;
  align-items: center;
  a {
    margin: 0 4px;
  }
}

:deep(.ant-upload-list-picture-card .ant-upload-list-item::before),
:deep(.ant-upload-list-item-actions) {
  display: none;
}

:deep(.ant-upload-list-item.selected) {
  border: 1px dashed #1890ff;
  background-color: rgba(24, 144, 255, 0.1);
}
</style>
