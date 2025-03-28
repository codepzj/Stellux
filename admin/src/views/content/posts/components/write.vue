<!-- 写文章展示组件 -->
<template>
  <div class="flex h-14 items-center justify-between">
    <a-input
      v-model:value="props.title"
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
    v-model="props.content"
    :theme="theme"
    previewTheme="github"
    @change="updateContent"
  />
</template>
<script lang="ts" setup>
import { computed } from "vue";
import { useThemeStore } from "@/store/theme";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";

const themeStore = useThemeStore();
const theme = computed(() => themeStore.tailwindTheme);

// 文章表单插槽
const props = defineProps<{
  mode: "create" | "update";
  title?: string;
  content?: string;
}>();

const emit = defineEmits(["update:title", "update:content"]);

// 共享title
const updateTitle = (e: InputEvent) => {
  emit("update:title", (e.target as HTMLInputElement).value);
};

// 共享content
const updateContent = (value: string) => {
  emit("update:content", value);
};
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
