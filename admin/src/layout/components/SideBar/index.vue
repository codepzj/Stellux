<template>
  <div>
    <a-menu
      v-model:selectedKeys="state.selectedKeys"
      mode="inline"
      :open-keys="state.openKeys"
      :items="items"
      @openChange="onOpenChange"
      @select="onSelect"
      :style="{ height: '100%', borderRight: 0 }"
    ></a-menu>
  </div>
</template>
<script lang="ts" setup>
import { VueElement, h, reactive } from "vue";
import { useRouter } from "vue-router";
import {
  HomeOutlined,
  UserOutlined,
  AppstoreOutlined,
  InteractionOutlined,
  LineChartOutlined,
  BellOutlined,
  SettingOutlined,
} from "@ant-design/icons-vue";
import type { ItemType } from "ant-design-vue";

function getItem(
  label: VueElement | string,
  key: string,
  icon?: any,
  children?: ItemType[],
  type?: "group"
): ItemType {
  return {
    key: key.charAt(0).toUpperCase() + key.slice(1), // 将key改为大驼峰
    icon,
    children,
    label,
    type,
  } as ItemType;
}

const items: ItemType[] = reactive([
  getItem("仪表盘", "Dashboard", () => h(HomeOutlined)),
  getItem("用户管理", "User", () => h(UserOutlined), [
    getItem("用户列表", "UserList", null),
    getItem("用户角色", "UserRole", null),
  ]),
  getItem("内容管理", "Content", () => h(AppstoreOutlined), [
    getItem("发布文章", "PublishArticle", null),
    getItem("文章列表", "ArticleList", null),
    getItem("草稿箱", "Draft", null),
    getItem("素材库", "MaterialLibrary", null),
    getItem("回收站", "RecycleBin", null),
  ]),
  getItem("互动管理", "Interaction", () => h(InteractionOutlined), [
    getItem("评论管理", "Comment", null),
    getItem("点赞管理", "Like", null),
    getItem("私信管理", "Message", null),
  ]),
  getItem("数据分析", "Data", () => h(LineChartOutlined), [
    getItem("文章统计", "ArticleStatistics", null),
    getItem("用户统计", "UserStatistics", null),
    getItem("互动统计", "InteractionStatistics", null),
  ]),
  getItem("通知中心", "Notice", () => h(BellOutlined), [
    getItem("系统通知", "SystemNotifications", null),
    getItem("用户通知", "UserNotifications", null),
  ]),
  getItem("系统设置", "System", () => h(SettingOutlined), [
    getItem("系统设置", "SystemSettings", null),
  ]),
]);

const state = reactive({
  rootSubmenuKeys: [
    "Dashboard",
    "User",
    "Content",
    "Interaction",
    "Data",
    "Notice",
    "System",
  ],
  openKeys: ["User", "Content"],
  selectedKeys: [],
});

const router = useRouter();

const onSelect = ({ key }: { key: string }) => {
  router.push({ name: key });
};

const onOpenChange = (openKeys: string[]) => {
  const latestOpenKey = openKeys.find(
    key => state.openKeys.indexOf(key) === -1
  );
  if (
    typeof latestOpenKey === "string" &&
    state.rootSubmenuKeys.indexOf(latestOpenKey) === -1
  ) {
    state.openKeys = openKeys;
  } else {
    state.openKeys = latestOpenKey ? [latestOpenKey] : [];
  }
};
</script>
