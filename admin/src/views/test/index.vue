<script setup lang="ts">
import Dashboard from "@uppy/dashboard";
import ImageEditor from "@uppy/image-editor";
import Uppy from "@uppy/core";
import zh_CN from "@uppy/locales/lib/zh_CN";
import XHR from "@uppy/xhr-upload";

import "@uppy/core/dist/style.css";
import "@uppy/dashboard/dist/style.css";
import "@uppy/image-editor/dist/style.min.css";

import { onMounted } from "vue";

onMounted(() => {
  new Uppy({
    locale: zh_CN,
  })
    .use(Dashboard, { inline: true, target: "#uppy-dashboard" })
    .use(ImageEditor)
    .use(XHR, { endpoint: "https://your-domain.com/upload" })
    .on("complete", result => {
      console.log("Upload result:", result);
    });
});
</script>

<template>
  <div id="uppy-dashboard"></div>
</template>
