// @/store/sidebar.ts
import { defineStore } from "pinia";
import { ref } from "vue";

export const useSidebarStore = defineStore(
  "sidebar",
  () => {
    // State
    const selectedKeys = ref<string[]>([]);
    const openKeys = ref<string[]>([]);

    // Actions
    function setSelectedKeys(keys: string[]) {
      selectedKeys.value = keys;
    }

    function setOpenKeys(keys: string[]) {
      openKeys.value = keys;
    }

    return {
      selectedKeys,
      openKeys,
      setSelectedKeys,
      setOpenKeys,
    };
  },
  {
    persist: true,
  }
);
