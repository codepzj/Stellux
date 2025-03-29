<template>
  <a-modal
    :open="props.open"
    @update:open="emit('update:open', $event)"
    title="文章配置信息"
    centered
    width="800px"
    :footer="null"
    @cancel="handleCancel"
  >
    <a-form :model="localForm" ref="formRef" :rules="props.rules">
      <a-form-item has-feedback label="标题" name="title">
        <a-input
          v-model:value="localForm.title"
          placeholder="请输入标题"
          :maxlength="50"
          @change="updateForm"
        />
      </a-form-item>
      <a-form-item has-feedback label="作者" name="author">
        <a-input
          v-model:value="localForm.author"
          placeholder="请输入作者"
          @change="updateForm"
        />
      </a-form-item>
      <a-form-item label="内容" name="content">
        <a-textarea
          v-model:value="localForm.content"
          placeholder="请输入内容"
          :rows="5"
          @change="updateForm"
        />
      </a-form-item>
      <a-form-item label="简介" name="description">
        <a-input
          v-model:value="localForm.description"
          placeholder="请输入简介"
          @change="updateForm"
        />
      </a-form-item>
      <a-form-item label="分类" name="category">
        <a-select
          v-model:value="localForm.category"
          style="width: 100%"
          placeholder="请选择分类"
          :options="props.category"
          @change="updateForm"
        ></a-select>
      </a-form-item>
      <a-form-item label="标签" name="tags">
        <a-select
          has-feedback
          v-model:value="localForm.tags"
          mode="multiple"
          style="width: 100%"
          placeholder="请选择标签"
          :options="props.tags"
          :max-tag-count="2"
          @change="updateForm"
        >
          <template #maxTagPlaceholder="omittedValues">
            <span style="color: red">+ {{ omittedValues.length }} ...</span>
          </template>
        </a-select>
      </a-form-item>
      <a-form-item label="封面" name="cover">
        <img
          v-if="localForm.cover"
          :src="localForm.cover"
          alt="封面"
          class="w-[160px] h-[90px] rounded-md cursor-pointer"
          @click="getList"
        />
        <a-button v-else type="primary" @click="getList">选择封面</a-button>
        <a-modal
          style="top: 10px"
          :open="pictureModalOpen"
          @update:open="pictureModalOpen = $event"
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
          <a-switch v-model:checked="localForm.is_top" @change="updateForm" />
        </a-form-item>
        <a-divider type="vertical" />
        <a-form-item label="是否发布" name="is_publish">
          <a-switch
            v-model:checked="localForm.is_publish"
            @change="updateForm"
          />
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
import { ref, watch } from "vue";
import PhotoWall from "@/components/photo-wall.vue";
import { getFilesByPage } from "@/api/file";
import { message } from "ant-design-vue";

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
  category: any[];
  tags: any[];
}>();

const emit = defineEmits<{
  (e: "update:open", value: boolean): void;
  (e: "update:form", value: any): void;
}>();

const formRef = ref<any>(null);
const pictureList = ref<any[]>([]);
const selectedPicture = ref<string>("");
const pictureModalOpen = ref(false);

// 复制 props.form，避免直接修改 props
const localForm = ref({ ...props.form });

// 监听 props.form 变化，保持数据同步
watch(
  () => props.form,
  newVal => {
    localForm.value = { ...newVal };
  },
  { deep: true, immediate: true }
);

// 更新表单数据
const updateForm = () => {
  emit("update:form", { ...localForm.value });
};

// 取消弹窗
const handleCancel = () => {
  formRef.value?.resetFields();
  emit("update:open", false);
};

// 获取图片列表
const getList = async () => {
  try {
    const res = await getFilesByPage({ page_no: 1, size: 20 });
    pictureList.value = res.data.list;
    pictureModalOpen.value = true;
  } catch (error: any) {
    message.error(`图片列表获取失败: ${error}`);
  }
};

// 选择图片
const handleSelectPicture = (photos: string[]) => {
  selectedPicture.value = photos[0];
};

// 确认选择图片
const handleSelectPictureConfirm = () => {
  pictureModalOpen.value = false;
  localForm.value.cover = selectedPicture.value;
  updateForm();
};

defineExpose({
  validate: () => {
    return formRef.value?.validate();
  },
});
</script>
