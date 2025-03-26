<template>
  <Write
    mode="create"
    v-model:title="form.title"
    v-model:content="form.content"
  >
    <template #action>
      <div class="flex gap-2">
        <a-button type="primary" @click="open = true">发布</a-button>
      </div>
    </template>
  </Write>
  <Modal
    v-model:open="open"
    :form="form"
    ref="modalRef"
    :rules="formRules"
    :category="category"
    :tags="tags"
  >
    <template #action>
        <a-button type="text" danger @click="open = false">取消</a-button>
        <a-button type="text" @click="onSubmitPost">发布</a-button>
    </template>
  </Modal>
</template>
<script lang="ts" setup>
import Write from "./components/write.vue";
import Modal from "./components/modal.vue";

import { onMounted, reactive, ref } from "vue";
import type { PostLabel, PostReq } from "@/types/posts";
import type { Rule } from "ant-design-vue/es/form";

const open = ref(false); // 弹窗状态
const form = reactive<PostReq>({
  title: "",
  content: "",
  author: "",
  description: "",
  category: "",
  tags: [],
  cover: "",
  is_top: false,
  is_publish: false,
});

const modalRef = ref<any>(null);

// 标签校验规则
const checkTags = async (_rule: Rule, value: Array<string>) => {
  if (value.length > 2) {
    return Promise.reject(new Error("标签不能多于2个"));
  }
  return Promise.resolve();
};

const formRules = ref({
  title: [{ required: true, message: "请输入标题" }],
  content: [{ required: true, message: "请输入内容" }],
  author: [{ required: true, message: "请输入作者" }],
  tags: [{ validator: checkTags, trigger: "change" }],
});

const category = ref<PostLabel[]>([
  {
    label: "技术",
    value: "technology",
  },
  {
    label: "生活",
    value: "life",
  },
]);

const tags = ref<PostLabel[]>([
  {
    label: "Vue",
    value: "vue",
  },
  {
    label: "React",
    value: "react",
  },
  {
    label: "JavaScript",
    value: "javascript",
  },
]);
onMounted(() => {
  console.log(modalRef.value);
});

const onSubmitPost = () => {
  modalRef.value.validate().then((res: any) => {
    console.log(res);
  });
};
</script>
