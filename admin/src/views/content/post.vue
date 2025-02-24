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
    v-model="debouncedContent"
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
      <a-form-item label="内容" name="content">
        <a-textarea
          v-model:value="debouncedContent"
          placeholder="请输入内容"
        />
      </a-form-item>
      <a-form-item label="作者" name="author">
        <a-input v-model:value="configForm.author" placeholder="请输入作者" />
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
      <a-form-item label="是否置顶" name="isTop">
        <a-switch v-model:checked="configForm.isTop" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script lang="ts" setup>
import { ref, reactive, computed, watch} from "vue";
import { MdEditor } from "md-editor-v3";
import { useThemeStore } from "@/store/theme";
import "md-editor-v3/lib/style.css";
import { createPost } from "@/api/modules/posts";
import { message } from "ant-design-vue";
import type { Rule } from "ant-design-vue/es/form";
import type { PostsReq } from "@/api/interfaces/posts";

const themeStore = useThemeStore();
const theme = computed(() => themeStore.tailwindTheme);
const debouncedContent = ref("");
const open = ref<boolean>(false);
const confirmLoading = ref<boolean>(false);
const layout = ref<string>("vertical");
const configFormRef = ref();
const key = "submit-post";
const configForm = reactive({
  title: "",
  content: "",
  author: "codepzj",
  category: "",
  tags: [],
  cover: "",
  isTop: false,
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
    .then(async (res: PostsReq) => {
      message.loading({
        content: "文章发布中...",
        key,
      });
      try {
        const result = await createPost(res);
        open.value = false;
        confirmLoading.value = false;
        configFormRef.value.resetFields();
        debouncedContent.value = "";
        message.success({
          content: result.msg,
          key,
        });
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
  category: [{ required: true, message: "请选择分类" }],
  tags: [{ required: true, message: "请选择标签" }],
};
// 防抖函数
const debounce = (fn: Function, delay: number) => {
  let timeout: ReturnType<typeof setTimeout>;
  return (...args: any[]) => {
    clearTimeout(timeout);
    timeout = setTimeout(() => fn(...args), delay);
  };
};

// 监听输入并防抖更新 text
watch(
  debouncedContent,
  debounce((newVal: string) => {
    configForm.content = newVal;
  }, 100)
);
</script>
<style lang="scss">
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
</style>
