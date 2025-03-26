<!-- 发布文章弹窗 -->
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
          :rows="10"
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
          ><template #maxTagPlaceholder> 。。。 </template></a-select
        >
      </a-form-item>
      <a-form-item label="封面" name="cover">
        <a-input v-model:value="props.form.cover" placeholder="请输入封面" />
      </a-form-item>
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

const formRef = ref<any>(null);

const props = defineProps<{
  open: boolean;
  form: {
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

const handleCancel = () => {
  formRef.value.resetFields();
  emit("update:open", false);
};

defineExpose({
  validate: () => {
    return formRef.value.validate();
  },
});
</script>
