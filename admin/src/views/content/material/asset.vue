<template>
  <div class="w-full mx-auto">
    <a-skeleton :loading="systemStore.loading" :paragraph="{ rows: 12 }">
      <PhotoWall :list="list" :mode="selectType" type="display">
        <template #action1>
          <a-button type="primary" danger @click="">删除</a-button>
        </template>
        <template #action2>
          <span>
            <a-select
              v-model:value="selectType"
              style="width: 6em"
              :options="selectTypeOptions"
              ><template #suffixIcon
                ><UnorderedListOutlined /></template
            ></a-select>
          </span>
          <UploadModal @update:list="getList" />
        </template>
      </PhotoWall>
    </a-skeleton>
  </div>
</template>
<script setup lang="ts">
import { onMounted, ref } from "vue";
import PhotoWall from "@/components/photo-wall.vue";
import UploadModal from "./components/upload-modal.vue";
import { getFilesByPage } from "@/api/file";
import type { IFile } from "@/types/file";
import { message } from "ant-design-vue";
import { useSystemStore } from "@/store/system";
import { UnorderedListOutlined } from "@ant-design/icons-vue";
const systemStore = useSystemStore();
const pageNo = ref(1);
const size = ref(84);
const total = ref(0);
const list = ref<IFile[]>([]);

const selectType = ref<"picture-card" | "picture">("picture-card");
const selectTypeOptions = ref<any[]>([
  {
    label: "图片墙",
    value: "picture-card",
  },
  {
    label: "图片卡",
    value: "picture",
  },
]);

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

onMounted(async () => {
  getList();
});
</script>
<style scoped lang="scss"></style>
