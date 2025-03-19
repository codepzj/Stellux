import { defineStore } from "pinia";
import { ref } from "vue";
import router from "@/router";

// 通过守卫路由传递key，动态增加tab栏
export const useTabStore = defineStore(
  "tab",
  () => {
    const panes = ref<{ title: string; key: string; closable?: boolean }[]>([
      { title: "仪表盘", key: "Dashboard", closable: false },
    ]);

    const activeKey = ref(panes.value[0].key);

    // 接受route传递的key和title
    const addTab = ({ key, title }: { key: string; title: string }) => {
      const existingTab = panes.value.find(tabItem => tabItem.key === key);
      if (!existingTab) {
        panes.value.push({ title, key, closable: true });
      }
      activeKey.value = key;
    };

    // 删除tab
    const removeTab = (targetKey: string) => {
      panes.value = panes.value.filter(item => item.key !== targetKey);
      if (activeKey.value === targetKey) {
        activeKey.value = panes.value[panes.value.length - 1].key;
        router.push({ name: activeKey.value });
      }
    };

    const editTab = (targetKey: string) => {
      removeTab(targetKey);
    };

    const changeTab = (key: string) => {
      activeKey.value = key;
      router.push({ name: key });
    };

    return {
      panes,
      activeKey,
      addTab, // 添加tab
      removeTab, // 删除tab
      editTab, // 编辑tab
      changeTab, // 切换tab
    };
  },
  {
    persist: true,
  }
);
