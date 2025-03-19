<template>
  <div class="w-full h-[calc(100vh-56px)] overflow-y-scroll">
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts"></script>
<style scoped>
/* 滚动条样式 */
::-webkit-scrollbar {
  width: 8px;
}
</style>
