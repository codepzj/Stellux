<template>
  <a-layout class="h-screen flex flex-col overflow-hidden">
    <a-layout-header
      :style="{
        height: '48px',
        background: themeStore.tailwindTheme === 'dark' ? '#000000' : '#ffffff',
        padding: '0 15px',
      }"
    >
      <app-header></app-header>
    </a-layout-header>
    <a-layout style="flex: 1; overflow: hidden;">
      <a-layout-sider
        :collapsed="width < 768"
        :style="{
          background: themeStore.tailwindTheme === 'dark' ? '#1d1d1d' : '#fff',
        }"
      >
        <Sidebar></Sidebar>
      </a-layout-sider>
      <a-layout style="margin: 0 0 0 15px;">
        <a-layout-content style="display: flex; flex-direction: column; height: 100%;">
          <tab
            style="
              position: sticky;
              top: 0;
              z-index: 10;
              background: inherit;
            "
          ></tab>
          <div style="flex: 1; overflow-y: auto;">
            <app-main></app-main>
          </div>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import Sidebar from "./components/SideBar/index.vue";
import Tab from "./components/Tab/index.vue";
import AppHeader from "./components/AppHeader.vue";
import AppMain from "./components/AppMain.vue";
import { useThemeStore } from "@/store/theme";
const themeStore = useThemeStore();
import { useWindowSize } from "@vueuse/core";

const { width } = useWindowSize();
</script>

<style lang="scss" scoped>
// Vertical scrollbar (side)
::-webkit-scrollbar {
  width: 8px; /* Vertical scrollbar width */
}

// Horizontal scrollbar (bottom)
::-webkit-scrollbar:horizontal {
  height: 5px; /* Horizontal scrollbar height */
  background: transparent;
}

// Track for both vertical and horizontal scrollbars
::-webkit-scrollbar-track {
  background: transparent; /* Transparent track */
}

// Thumb for both vertical and horizontal scrollbars
::-webkit-scrollbar-thumb {
  background: #aaa; /* Scrollbar color */
  border-radius: 3px;
}
</style>