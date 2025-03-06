<template>
  <div class="flex h-14 items-center justify-between">
    <a-input
      v-model:value="configForm.title"
      addon-before="标题"
      show-count
      :maxlength="20"
      class="w-2/5"
    />
    <div class="flex items-center gap-2">
      <a-button type="dashed">保存</a-button>
      <a-button type="dashed">设置</a-button>
      <a-button type="primary" @click="handleCreatePost">发布</a-button>
    </div>
  </div>
  <MdEditor
    class="h-[calc(100vh-180px)]"
    :style="{ fontFamily: 'LXGW WenKai GB Screen' }"
    v-model="configForm.content"
    :theme="theme"
    previewTheme="vuepress"
  />
  <a-modal
    v-model:open="open"
    title="文章配置信息"
    :confirm-loading="confirmLoading"
    okText="确定"
    cancelText="取消"
    @ok="handleOk"
    style="top: 20px"
    width="800px"
  >
    <a-form
      :layout="layout"
      :model="configForm"
      :rules="rules"
      ref="configFormRef"
    >
      <a-form-item label="标题" name="title">
        <a-input v-model:value="configForm.title" placeholder="请输入标题" />
      </a-form-item>
      <a-form-item label="作者" name="author">
        <a-input v-model:value="configForm.author" placeholder="请输入作者" />
      </a-form-item>
      <a-form-item label="内容" name="content">
        <a-textarea
          v-model:value="configForm.content"
          placeholder="请输入内容"
        />
      </a-form-item>
      <a-form-item label="简介" name="description">
        <a-input
          v-model:value="configForm.description"
          placeholder="请输入简介"
        />
      </a-form-item>
      <a-form-item label="分类" name="category">
        <a-select
          v-model:value="configForm.category"
          style="width: 100%"
          placeholder="请选择分类"
          :options="options.category"
        ></a-select>
      </a-form-item>
      <a-form-item label="标签" name="tags">
        <a-select
          v-model:value="configForm.tags"
          mode="multiple"
          style="width: 100%"
          placeholder="请选择标签"
          :options="options.tags"
        ></a-select>
      </a-form-item>
      <a-form-item label="封面" name="cover">
        <a-input v-model:value="configForm.cover" placeholder="请输入封面" />
      </a-form-item>
      <a-form-item label="操作">
        <a-space>
          <span>置顶</span>
          <a-switch v-model:checked="configForm.isTop" />
          <a-divider type="vertical" />
          <span>发布</span>
          <a-switch v-model:checked="configForm.isPublish" />
        </a-space>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script lang="ts" setup>
import { ref, reactive, computed } from "vue";
import { MdEditor } from "md-editor-v3";
import { useThemeStore } from "@/store/theme";
import "md-editor-v3/lib/style.css";
import { createPost } from "@/api/modules/posts";
import { message } from "ant-design-vue";
import type { Rule } from "ant-design-vue/es/form";
import { useRouter } from "vue-router";

const router = useRouter();
const themeStore = useThemeStore();
const theme = computed(() => themeStore.tailwindTheme);
const open = ref<boolean>(false);
const confirmLoading = ref<boolean>(false);
const layout = ref<string>("vertical");
const configFormRef = ref();
const key = "submit-post";
const configForm = reactive({
  title: "",
  content: "",
  author: "codepzj",
  description: "",
  category: "",
  tags: [],
  cover: "",
  isTop: false,
  isPublish: true,
});
const options = reactive({
  category: [
    { value: "技术", label: "技术" },
    { value: "生活", label: "生活" },
    { value: "学习", label: "学习" },
    { value: "其他", label: "其他" },
  ],
  tags: [
    { value: "golang", label: "Golang" },
    { value: "javascript", label: "JavaScript" },
    { value: "typescript", label: "TypeScript" },
    { value: "python", label: "Python" },
    { value: "java", label: "Java" },
    { value: "csharp", label: "C#" },
  ],
});
const handleCreatePost = async () => {
  open.value = true;
};

const handleOk = async () => {
  configFormRef.value
    .validate()
    .then(async () => {
      message.loading({
        content: "文章发布中...",
        key,
      });
      try {
        const result = await createPost(configForm);
        open.value = false;
        confirmLoading.value = false;
        configFormRef.value.resetFields();
        message.success({
          content: result.msg,
          key,
        });
        router.push({ name: "ArticleList" });
      } catch (error: any) {
        message.error({
          content: error,
          key,
        });
      } finally {
        confirmLoading.value = false;
      }
    })
    .catch(() => {
      message.error({
        content: "文章发布失败",
        key,
      });
    });
};

const rules: Record<string, Rule[]> = {
  title: [{ required: true, message: "请输入标题" }],
  content: [{ required: true, message: "请输入内容" }],
  author: [{ required: true, message: "请输入作者" }],
  description: [{ required: true, message: "请输入简介" }],
  category: [{ required: true, message: "请选择分类" }],
  tags: [{ required: true, message: "请选择标签" }],
};
</script>
<style lang="scss">
@import url("https://cdn.jsdelivr.net/npm/cn-fontsource-lxgw-wen-kai-gb-screen@1.0.6/font.min.css");

.md-editor-toolbar-wrapper {
  height: 40px;
  .md-editor-toolbar {
    height: 100%;
    .md-editor-toolbar-left,
    .md-editor-toolbar-right {
      height: 100%;
      .md-editor-toolbar-item {
        height: 100%;
        svg {
          width: 32px;
          height: 32px;
        }
      }
    }
  }
}

.cm-line {
  font-size: 17px;
  margin-top: 10px;
  font-weight: 500;
  font-family: "LXGW WenKai GB Screen";
}
</style>
