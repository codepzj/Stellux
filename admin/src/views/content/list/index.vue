<template>
  <a-table
    :columns="columns"
    :data-source="posts.list"
    :scroll="{ x: 1200 }"
    :expand-column-width="100"
  >
    <template #headerCell="{ column }">
      <template v-if="column.key === 'title'"> 标题 </template>
    </template>

    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'title'">
        <a>
          {{ record.title }}
        </a>
      </template>
      <template v-else-if="column.key === 'author'">
        <span>
          {{ record.author }}
        </span>
      </template>
      <template v-else-if="column.key === 'created_at'">
        <span>
          {{ formatTime(record.created_at) }}
        </span>
      </template>
      <template v-else-if="column.key === 'updated_at'">
        <span>
          {{ formatTime(record.updated_at) }}
        </span>
      </template>
      <template v-else-if="column.key === 'category'">
        <a-tag class="text-sm">{{ record.category }}</a-tag>
      </template>
      <template v-else-if="column.key === 'tags'">
        <a-tag class="text-sm" v-for="tag in record.tags" :key="tag">
          {{ tag }}
        </a-tag>
      </template>
      <template v-else-if="column.key === 'cover'">
        <a-image
          :width="50"
          :height="50"
          :src="record.cover"
          fallback="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMIAAADDCAYAAADQvc6UAAABRWlDQ1BJQ0MgUHJvZmlsZQAAKJFjYGASSSwoyGFhYGDIzSspCnJ3UoiIjFJgf8LAwSDCIMogwMCcmFxc4BgQ4ANUwgCjUcG3awyMIPqyLsis7PPOq3QdDFcvjV3jOD1boQVTPQrgSkktTgbSf4A4LbmgqISBgTEFyFYuLykAsTuAbJEioKOA7DkgdjqEvQHEToKwj4DVhAQ5A9k3gGyB5IxEoBmML4BsnSQk8XQkNtReEOBxcfXxUQg1Mjc0dyHgXNJBSWpFCYh2zi+oLMpMzyhRcASGUqqCZ16yno6CkYGRAQMDKMwhqj/fAIcloxgHQqxAjIHBEugw5sUIsSQpBobtQPdLciLEVJYzMPBHMDBsayhILEqEO4DxG0txmrERhM29nYGBddr//5/DGRjYNRkY/l7////39v///y4Dmn+LgeHANwDrkl1AuO+pmgAAADhlWElmTU0AKgAAAAgAAYdpAAQAAAABAAAAGgAAAAAAAqACAAQAAAABAAAAwqADAAQAAAABAAAAwwAAAAD9b/HnAAAHlklEQVR4Ae3dP3PTWBSGcbGzM6GCKqlIBRV0dHRJFarQ0eUT8LH4BnRU0NHR0UEFVdIlFRV7TzRksomPY8uykTk/zewQfKw/9znv4yvJynLv4uLiV2dBoDiBf4qP3/ARuCRABEFAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghggQAQZQKAnYEaQBAQaASKIAQJEkAEEegJmBElAoBEgghgg0Aj8i0JO4OzsrPv69Wv+hi2qPHr0qNvf39+iI97soRIh4f3z58/u7du3SXX7Xt7Z2enevHmzfQe+oSN2apSAPj09TSrb+XKI/f379+08+A0cNRE2ANkupk+ACNPvkSPcAAEibACyXUyfABGm3yNHuAECRNgAZLuYPgEirKlHu7u7XdyytGwHAd8jjNyng4OD7vnz51dbPT8/7z58+NB9+/bt6jU/TI+AGWHEnrx48eJ/EsSmHzx40L18+fLyzxF3ZVMjEyDCiEDjMYZZS5wiPXnyZFbJaxMhQIQRGzHvWR7XCyOCXsOmiDAi1HmPMMQjDpbpEiDCiL358eNHurW/5SnWdIBbXiDCiA38/Pnzrce2YyZ4//59F3ePLNMl4PbpiL2J0L979+7yDtHDhw8vtzzvdGnEXdvUigSIsCLAWavHp/+qM0BcXMd/q25n1vF57TYBp0a3mUzilePj4+7k5KSLb6gt6ydAhPUzXnoPR0dHl79WGTNCfBnn1uvSCJdegQhLI1vvCk+fPu2ePXt2tZOYEV6/fn31dz+shwAR1sP1cqvLntbEN9MxA9xcYjsxS1jWR4AIa2Ibzx0tc44fYX/16lV6NDFLXH+YL32jwiACRBiEbf5KcXoTIsQSpzXx4N28Ja4BQoK7rgXiydbHjx/P25TaQAJEGAguWy0+2Q8PD6/Ki4R8EVl+bzBOnZY95fq9rj9zAkTI2SxdidBHqG9+skdw43borCXO/ZcJdraPWdv22uIEiLA4q7nvvCug8WTqzQveOH26fodo7g6uFe/a17W3+nFBAkRYENRdb1vkkz1CH9cPsVy/jrhr27PqMYvENYNlHAIesRiBYwRy0V+8iXP8+/fvX11Mr7L7ECueb/r48eMqm7FuI2BGWDEG8cm+7G3NEOfmdcTQw4h9/55lhm7DekRYKQPZF2ArbXTAyu4kDYB2YxUzwg0gi/41ztHnfQG26HbGel/crVrm7tNY+/1btkOEAZ2M05r4FB7r9GbAIdxaZYrHdOsgJ/wCEQY0J74TmOKnbxxT9n3FgGGWWsVdowHtjt9Nnvf7yQM2aZU/TIAIAxrw6dOnAWtZZcoEnBpNuTuObWMEiLAx1HY0ZQJEmHJ3HNvGCBBhY6jtaMoEiJB0Z29vL6ls58vxPcO8/zfrdo5qvKO+d3Fx8Wu8zf1dW4p/cPzLly/dtv9Ts/EbcvGAHhHyfBIhZ6NSiIBTo0LNNtScABFyNiqFCBChULMNNSdAhJyNSiECRCjUbEPNCRAhZ6NSiAARCjXbUHMCRMjZqBQiQIRCzTbUnAARcjYqhQgQoVCzDTUnQIScjUohAkQo1GxDzQkQIWejUogAEQo121BzAkTI2agUIkCEQs021JwAEXI2KoUIEKFQsw01J0CEnI1KIQJEKNRsQ80JECFno1KIABEKNdtQcwJEyNmoFCJAhELNNtScABFyNiqFCBChULMNNSdAhJyNSiECRCjUbEPNCRAhZ6NSiAARCjXbUHMCRMjZqBQiQIRCzTbUnAARcjYqhQgQoVCzDTUnQIScjUohAkQo1GxDzQkQIWejUogAEQo121BzAkTI2agUIkCEQs021JwAEXI2KoUIEKFQsw01J0CEnI1KIQJEKNRsQ80JECFno1KIABEKNdtQcwJEyNmoFCJAhELNNtScABFyNiqFCBChULMNNSdAhJyNSiECRCjUbEPNCRAhZ6NSiAARCjXbUHMCRMjZqBQiQIRCzTbUnAARcjYqhQgQoVCzDTUnQIScjUohAkQo1GxDzQkQIWejUogAEQo121BzAkTI2agUIkCEQs021JwAEXI2KoUIEKFQsw01J0CEnI1KIQJEKNRsQ80JECFno1KIABEKNdtQcwJEyNmoFCJAhELNNtScABFyNiqFCBChULMNNSdAhJyNSiEC/wGgKKC4YMA4TAAAAABJRU5ErkJggg=="
        />
      </template>
      <template v-else-if="column.key === 'is_publish'">
        <span>
          {{ record.is_publish ? "已发布" : "未发布" }}
        </span>
      </template>
      <template v-else-if="column.key === 'action'">
        <span>
          <a-button
            size="small"
            @click="handleChange(record.is_publish, record.id)"
            >{{ record.is_publish ? "下架" : "发布" }}</a-button
          >
          <a-divider type="vertical" />
          <a-button size="small">编辑</a-button>
          <a-divider type="vertical" />
          <a-button size="small" danger @click="handleDelete(record.id)"
            >删除</a-button
          >
        </span>
      </template>
    </template>
  </a-table>
</template>
<script lang="ts" setup>
import { onMounted, reactive } from "vue";
import type { PostVO } from "@/api/interfaces/posts";
import {
  deletePostSoft,
  getPostsByPage,
  updatePostStatus,
} from "@/api/modules/posts";

import { formatTime } from "@/utils/time";
import { message } from "ant-design-vue";

const posts = reactive({
  page_no: 1,
  size: 10,
  total_count: 0,
  total_page: 1,
  list: [] as PostVO[],
});

onMounted(async () => {
  const res = await getPostsByPage(1, 10);
  posts.list = res.data.list;
});

const columns = [
  {
    name: "标题",
    width: 200,
    dataIndex: "title",
    key: "title",
    fixed: "left",
  },
  {
    title: "作者",
    width: 100,
    dataIndex: "author",
    key: "author",
    fixed: "left",
  },
  {
    title: "创建时间",
    width: 200,
    dataIndex: "created_at",
    key: "created_at",
  },
  {
    title: "更新时间",
    width: 200,
    dataIndex: "updated_at",
    key: "updated_at",
  },
  {
    title: "分类",
    width: 100,
    key: "category",
    dataIndex: "category",
  },
  {
    title: "标签",
    width: 200,
    key: "tags",
    dataIndex: "tags",
  },
  {
    title: "封面",
    width: 100,
    key: "cover",
    dataIndex: "cover",
  },
  {
    title: "状态",
    width: 100,
    key: "is_publish",
    dataIndex: "is_publish",
  },
  {
    title: "操作",
    width: 250,
    key: "action",
    dataIndex: "action",
    fixed: "right",
  },
];

const handleChange = async (is_publish: boolean, id: string) => {
  try {
    const res = await updatePostStatus({
      id: id,
      is_publish: !is_publish,
    });
    if (res.code === 200) {
      message.success(is_publish ? "下架成功" : "发布成功");
      posts.list = posts.list.map((item) => {
        if (item.id === id) {
          item.is_publish = !is_publish;
        }
        return item;
      });
    } else {
      message.error(res.msg);
    }
  } catch (err: any) {
    message.error(err);
  }
};

const handleDelete = async (id: string) => {
  try {
    const res = await deletePostSoft(id);
    message.success(res.msg);
    posts.list = posts.list.filter((item) => item.id !== id);
  } catch (err: any) {
    message.error(err);
  }
};
</script>
