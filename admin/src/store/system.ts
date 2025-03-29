import { defineStore } from "pinia";
import { ref } from "vue";

export const useSystemStore = defineStore("system", () => {
  const spinning = ref(false);
  const tip = ref("加载中...");
  const loading = ref(false);

  // 设置页面加载状态
  const setSpinning = (value: boolean) => {
    spinning.value = value;
  };

  // 设置骨架屏加载状态
  const setSkeleton = (value: boolean) => {
    loading.value = value;
  };

  return {
    spinning,
    tip,
    setSpinning,
    loading,
    setSkeleton,
  };
});
