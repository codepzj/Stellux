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

import { reactive, ref } from "vue";
import type { PostLabel, PostReq, PostVO } from "@/types/posts";
import type { Rule } from "ant-design-vue/es/form";
import { createPost } from "@/api/posts";
import { message } from "ant-design-vue";
import type { Response } from "@/types/response";
import { useRouter } from "vue-router";

const router = useRouter();
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
  is_publish: true,
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
    value: "技术",
  },
  {
    label: "生活",
    value: "生活",
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

const onSubmitPost = () => {
  modalRef.value.validate().then((res: PostReq) => {
    createPost(res).then((res: Response<PostVO>) => {
      message.success(res.msg);
      open.value = false;
      router.push("/content/list");
    });
  });
};
</script>
