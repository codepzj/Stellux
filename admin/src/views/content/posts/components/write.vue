<template>
  <div class="flex h-14 items-center justify-between">
    <a-input
      v-model:value="processedTitle"
      addon-before="标题"
      show-count
      :maxlength="50"
      class="md:w-2/5 w-[300px]"
      @change="updateTitle"
    />
    <slot name="action" />
  </div>
  <MdEditor
    class="h-[calc(100vh-180px)]"
    v-model="processedContent"
    :theme="theme"
    previewTheme="github"
    @change="handleContentChange"
  />
</template>

<script lang="ts" setup>
import { computed, ref, watch } from "vue";
import { useThemeStore } from "@/store/theme";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";

const themeStore = useThemeStore();
const theme = computed(() => themeStore.tailwindTheme);

const props = defineProps<{
  mode: "create" | "update";
  title?: string;
  content?: string;
}>();

const emit = defineEmits(["update:title", "update:content"]);

// 使用 ref 管理处理后的值
const processedTitle = ref(props.title?.trim() || "");
const processedContent = ref(props.content?.trim() || "");

// 处理 content 并提取一级标题
const processContent = (content: string) => {
  const trimmedContent = (content || "").trim();
  let newTitle = processedTitle.value;
  let newContent = trimmedContent;

  // 检查是否有内容
  if (trimmedContent) {
    const lines = trimmedContent.split("\n");
    const firstLine = lines[0].trim();

    // 如果第一行是一级标题，则提取标题并移除该行
    if (firstLine.startsWith("# ")) {
      newTitle = firstLine.substring(2).trim();
      newContent = lines.slice(1).join("\n").trim();
    }
  }

  return { newTitle, newContent };
};

// 处理 Markdown 编辑器内容变化
const handleContentChange = (value: string) => {
  const { newTitle, newContent } = processContent(value);
  processedTitle.value = newTitle;
  processedContent.value = newContent;
  emit("update:title", newTitle);
  emit("update:content", newContent);
};

// 处理标题输入框变化
const updateTitle = (e: InputEvent) => {
  const newTitle = (e.target as HTMLInputElement).value.trim();
  processedTitle.value = newTitle;
  emit("update:title", newTitle);
};

// 监听 props 变化以同步外部更新
watch(
  () => props.content,
  newContent => {
    if (newContent?.trim() !== processedContent.value) {
      const { newTitle, newContent: processed } = processContent(
        newContent || ""
      );
      processedTitle.value = newTitle;
      processedContent.value = processed;
    }
  }
);

watch(
  () => props.title,
  newTitle => {
    if (newTitle?.trim() !== processedTitle.value) {
      processedTitle.value = newTitle?.trim() || "";
    }
  }
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

.cm-line {
  font-size: 17px;
  margin-top: 10px;
  font-weight: 500;
  font-family: "consolas";
}
</style>
