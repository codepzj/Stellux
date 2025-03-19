<template>
  <a-config-provider
    :theme="{
      algorithm:
        themeStore.tailwindTheme === 'dark'
          ? theme.darkAlgorithm
          : theme.defaultAlgorithm,
      token: {
        fontSize: 15,
        wireframe: true,
        fontSizeSM: 13,
        borderRadius: 8,
      },
    }"
    :locale="locale"
  >
    <div class="h-screen dark:bg-gray-900">
      <!-- 全局居中和遮罩 -->
      <a-spin
        :spinning="spinning"
        :tip="tip"
        size="large"
        class="h-screen font-bold bg-white dark:bg-gray-900"
        ><router-view></router-view>
      </a-spin>
    </div>
  </a-config-provider>
</template>

<script setup lang="ts">
import zhCN from "ant-design-vue/es/locale/zh_CN";
import { useThemeStore } from "@/store/theme";
import { ref, onMounted, computed } from "vue";
import { theme } from "ant-design-vue";
import { message } from "ant-design-vue";
import { useSystemStore } from "./store/system";

message.config({
  duration: 1.5,
  maxCount: 1,
  rtl: true,
});

const locale = ref(zhCN);
const themeStore = useThemeStore();
const systemStore = useSystemStore();
const spinning = computed(() => systemStore.spinning);
const tip = computed(() => systemStore.tip);
onMounted(() => {
  themeStore.initTheme();
  console.log(themeStore.tailwindTheme);
});
</script>

<style lang="scss" scoped>
:deep(.ant-spin-nested-loading .ant-spin) {
  position: absolute;
  top: 0;
  inset-inline-start: 0;
  z-index: 999;
  display: block;
  width: 100%;
  height: 100vh;
  max-height: 100vh;
}
</style>
