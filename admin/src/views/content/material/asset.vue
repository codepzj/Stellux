<template>
  <div class="w-[98%] mx-auto">
    <UploadModal @update:list="getList" />
    <a-skeleton :loading="systemStore.loading" :paragraph="{ rows: 12 }">
      <PhotoWall
        :list="list"
        :current="pageNo"
        :size="size"
        :total="total"
        @update:pagination="handlePagePagination"
        @update:list="handleDeletePhoto"
      />
    </a-skeleton>
  </div>
</template>
<script setup lang="ts">
import { onMounted, ref } from "vue";
import PhotoWall from "./components/photo-wall.vue";
import UploadModal from "./components/upload-modal.vue";
import { getFilesByPage } from "@/api/modules/file";
import type { IFile } from "@/api/interfaces/file";
import { message, type UploadFile } from "ant-design-vue";
import { useSystemStore } from "@/store/system";

const systemStore = useSystemStore();
const pageNo = ref(1);
const size = ref(84);
const total = ref(0);
const list = ref<IFile[]>([]);
const getList = async () => {
  systemStore.setSkeleton(true);
  try {
    const res = await getFilesByPage({
      page_no: pageNo.value,
      size: size.value,
    });
    list.value = res.data.list;
    total.value = res.data.total_count;
  } catch (error: any) {
    message.error({
      content: error + "，图片列表获取失败",
    });
  } finally {
    systemStore.setSkeleton(false);
  }
};

// 分页
const handlePagePagination = (pagination: {
  current: number;
  size: number;
}) => {
  pageNo.value = pagination.current;
  size.value = pagination.size;
  getList();
};

const handleDeletePhoto = (file: UploadFile) => {
  list.value = list.value.filter((item) => item.uid !== file.uid);
};

onMounted(async () => {
  getList();
});
</script>
<style scoped lang="scss"></style>
