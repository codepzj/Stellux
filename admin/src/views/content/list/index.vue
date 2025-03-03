<template>
  <div>文章列表</div>
</template>
<script lang="ts" setup>
import { onMounted, reactive } from "vue";
import type { PostsVO } from "@/api/interfaces/posts";
import { getPostsByPage } from "@/api/modules/posts";

const posts = reactive({
  page_no: 1,
  size: 10,
  total_count: 0,
  total_page: 1,
  list: [] as PostsVO[],
});

onMounted(async () => {
  try {
    const res = await getPostsByPage(1, 10);
    posts.list = res.data.list;
    posts.total_count = res.data.total_count;
    posts.total_page = res.data.total_page;
    posts.page_no = res.data.page_no;
    posts.size = res.data.size;
    console.log(posts);
  } catch (error) {
    console.error(error);
  }
});
</script>
