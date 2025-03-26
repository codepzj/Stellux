<template>
  <a-layout class="h-screen overflow-hidden">
    <a-layout-header
      :style="{
        height: '48px',
        background: themeStore.tailwindTheme === 'dark' ? '#00151d' : '#f2f3f4',
        padding: '0 15px',
      }"
    >
      <app-header></app-header>
    </a-layout-header>
    <a-layout style="min-height: calc(100vh - 48px)">
      <a-layout-sider
        :collapsed="width<768"
        :style="{
          background: themeStore.tailwindTheme === 'dark' ? '#1d1d1d' : '#fff',
        }"
      >
        <Sidebar></Sidebar>
      </a-layout-sider>
      <a-layout style="padding: 0 15px">
        <a-layout-content>
          <tab></tab>
          <app-main></app-main>
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

<style lang="scss" scoped></style>
