<!-- 发布文章弹窗 ui控件 -->
<template>
  <a-modal
    v-model:open="props.open"
    title="文章配置信息"
    centered
    width="800px"
    :footer="null"
    @cancel="handleCancel"
  >
    <a-form :model="props.form" ref="formRef" :rules="props.rules">
      <a-form-item has-feedback label="标题" name="title">
        <a-input
          v-model:value="props.form.title"
          placeholder="请输入标题"
          :maxlength="50"
        />
      </a-form-item>
      <a-form-item has-feedback label="作者" name="author">
        <a-input v-model:value="props.form.author" placeholder="请输入作者" />
      </a-form-item>
      <a-form-item label="内容" name="content">
        <a-textarea
          v-model:value="props.form.content"
          placeholder="请输入内容"
          :rows="5"
        />
      </a-form-item>
      <a-form-item label="简介" name="description">
        <a-input
          v-model:value="props.form.description"
          placeholder="请输入简介"
        />
      </a-form-item>
      <a-form-item label="分类" name="category">
        <a-select
          v-model:value="props.form.category"
          style="width: 100%"
          placeholder="请选择分类"
          :options="props.category"
        ></a-select>
      </a-form-item>
      <a-form-item label="标签" name="tags">
        <a-select
          has-feedback
          v-model:value="props.form.tags"
          mode="multiple"
          style="width: 100%"
          placeholder="请选择标签"
          :options="props.tags"
          :max-tag-count="2"
          ><template #maxTagPlaceholder="omittedValues">
            <span style="color: red">+ {{ omittedValues.length }} ...</span>
          </template></a-select
        >
      </a-form-item>
      <a-form-item label="封面" name="cover">
        <img
          v-if="props.form.cover"
          :src="props.form.cover"
          alt="封面"
          class="w-[160px] h-[90px] rounded-md cursor-pointer"
          @click="getList"
        />
        <a-button v-else type="primary" @click="getList"
          >选择封面</a-button
        >
        <a-modal
          style="top: 10px"
          bodyStyle="height: 600px; margin: 0 auto; overflow-y: auto;"
          v-model:visible="pictureModalOpen"
          title="选择封面"
          width="610px"
          :footer="null"
        >
          <PhotoWall
            :list="pictureList"
            mode="picture-card"
            type="select"
            @selected-picture="handleSelectPicture"
          ></PhotoWall>
          <a-divider />
          <div class="flex justify-end">
            <a-button type="primary" @click="handleSelectPictureConfirm"
              >确定</a-button
            >
          </div>
        </a-modal>
      </a-form-item>
      <div class="flex justify-start items-center gap-4">
        <a-form-item label="是否置顶" name="is_top">
          <a-switch v-model:checked="props.form.is_top" />
        </a-form-item>
        <a-divider type="vertical" />
        <a-form-item label="是否发布" name="is_publish">
          <a-switch v-model:checked="props.form.is_publish" />
        </a-form-item>
      </div>
    </a-form>
    <a-divider />
    <div class="flex justify-end">
      <slot name="action" />
    </div>
  </a-modal>
</template>
<script lang="ts" setup>
import type { PostLabel } from "@/types/posts";
import { ref } from "vue";
import PhotoWall from "@/components/photo-wall.vue";
import { getFilesByPage } from "@/api/file";
import { message } from "ant-design-vue";

const formRef = ref<any>(null);
const pictureList = ref<any[]>([]);
const selectedPicture = ref<string>("");

const props = defineProps<{
  open: boolean;
  form: {
    id?: string;
    title: string;
    author: string;
    content: string;
    description: string;
    category: string;
    tags: string[];
    cover: string;
    is_top: boolean;
    is_publish: boolean;
  };
  rules: Record<string, any>;
  category: PostLabel[];
  tags: PostLabel[];
}>();

const emit = defineEmits<{
  (e: "update:open", value: boolean): void;
}>();

const pictureModalOpen = ref(false);

const getList = async () => {
  try {
    const res = await getFilesByPage({
      page_no: 1,
      size: 20,
    });
    pictureList.value = res.data.list;
    pictureModalOpen.value = true;
  } catch (error: any) {
    message.error({
      content: error + "，图片列表获取失败",
    });
  }
};

const handleCancel = () => {
  formRef.value.resetFields();
  emit("update:open", false);
};

// 选择图片
const handleSelectPicture = (photos: string[]) => {
  selectedPicture.value = photos[0];
};

// 确认选择图片
const handleSelectPictureConfirm = () => {
  pictureModalOpen.value = false;
  props.form.cover = selectedPicture.value;
};

defineExpose({
  validate: () => {
    return formRef.value.validate();
  },
});
</script>
