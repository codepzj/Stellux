<template>
  <a-table :columns="columns" :data-source="userList">
    <template #headerCell="{ column }">
      <template v-if="column.key === 'name'">
        <span> Name </span>
      </template>
    </template>

    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'username'">
        <a>
          {{ record.username }}
        </a>
      </template>
      <template v-else-if="column.key === 'role_id'">
        <span>
          <a-tag class="text-sm" :color="toRole(record.role_id).color">
            {{ toRole(record.role_id).role }}
          </a-tag>
        </span>
      </template>
      <template v-else-if="column.key === 'created_at'">
        {{ formatTime(record.created_at) }}
      </template>
      <template v-else-if="column.key === 'updated_at'">
        {{ formatTime(record.updated_at) }}
      </template>
      <template v-else-if="column.key === 'action'">
        <span>
          <a>Invite 一 {{ record.name }}</a>
          <a-divider type="vertical" />
          <a>Delete</a>
          <a-divider type="vertical" />
          <a> More actions </a>
        </span>
      </template>
    </template>
  </a-table>
</template>
<script lang="ts" setup>
import { ref, onMounted } from "vue";
import type { UserListVO } from "@/api/interfaces/user";
import { getUserList } from "@/api/modules/user";
import { message } from "ant-design-vue";
import { formatTime } from "@/utils/time";
const userList = ref<UserListVO>([]);

onMounted(async () => {
  try {
    const res = await getUserList();
    userList.value = res.data;
  } catch (error: any) {
    message.error(error);
  }
});
const columns = [
  {
    title: "用户名",
    dataIndex: "username",
    key: "username",
  },
  {
    title: "角色",
    dataIndex: "role_id",
    key: "role_id",
  },
  {
    title: "创建时间",
    dataIndex: "created_at",
    key: "created_at",
  },
  {
    title: "更新时间",
    key: "updated_at",
    dataIndex: "updated_at",
  },
  {
    title: "操作",
    key: "action",
  },
];

function toRole(role_id: number) {
  switch (role_id) {
    case 0:
      return {
        role: "管理员",
        color: "red", // 管理员使用红色
      };
    case 1:
      return {
        role: "普通用户",
        color: "orange", // 普通用户使用橙色
      };
    case 2:
      return {
        role: "游客",
        color: "green", // 游客使用蓝色
      };
    default:
      return {
        role: "未知",
        color: "gray", // 未知角色使用灰色
      };
  }
}
</script>
