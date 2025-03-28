import {
  createWebHistory,
  createRouter,
  type RouteRecordRaw,
} from "vue-router";

import Layout from "@/layout/index.vue";
import Dashboard from "@/views/dashboard/index.vue";
import UserList from "@/views/user/list.vue";
import Login from "@/views/auth/login.vue";
import Test from "@/views/test/index.vue";
import { errorView } from "@/router/modules/errorView";
import UserRole from "@/views/user/role.vue";
import CreateArticle from "@/views/content/posts/create.vue";
import EditArticle from "@/views/content/posts/edit.vue";
import ArticleList from "@/views/content/list/index.vue";
import Draft from "@/views/content/draft.vue";
import MaterialLibrary from "@/views/content/material/asset.vue";
import RecycleBin from "@/views/content/bin.vue";
import Comment from "@/views/interaction/comment.vue";
import Like from "@/views/interaction/like.vue";
import ArticleStatistics from "@/views/analyse/post.vue";
import UserStatistics from "@/views/analyse/user.vue";
import InteractionStatistics from "@/views/analyse/interaction.vue";
import SystemNotifications from "@/views/notification/system.vue";
import UserNotifications from "@/views/notification/user.vue";
import SystemSettings from "@/views/system/index.vue";

export const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: Layout,
    children: [
      {
        path: "",
        name: "Dashboard",
        component: Dashboard,
        meta: { title: "仪表盘", icon: "HomeOutlined" },
      },
      {
        path: "/user",
        name: "User",
        meta: { title: "用户管理", icon: "UserOutlined" },
        children: [
          {
            path: "list",
            component: UserList,
            name: "UserList",
            meta: { title: "用户列表", icon: "UnorderedListOutlined" },
          },
          {
            path: "role",
            component: UserRole,
            name: "UserRole",
            meta: { title: "用户角色", icon: "UsergroupAddOutlined" },
          },
        ],
      },
      {
        path: "/content",
        name: "Content",
        meta: { title: "内容管理", icon: "AppstoreOutlined" },
        children: [
          {
            path: "create",
            component: CreateArticle,
            name: "CreateArticle",
            meta: { title: "发布文章", icon: "FileTextOutlined" },
          },
          {
            path: "edit/:id",
            component: EditArticle,
            name: "EditArticle",
            meta: { title: "编辑文章", hidden: true },
          },
          {
            path: "list",
            component: ArticleList,
            name: "ArticleList",
            meta: { title: "文章列表", icon: "FileDoneOutlined" },
          },
          {
            path: "draft",
            component: Draft,
            name: "Draft",
            meta: { title: "草稿箱", hidden: true },
          },
          {
            path: "assets",
            component: MaterialLibrary,
            name: "MaterialLibrary",
            meta: { title: "素材库", icon: "FileImageOutlined" },
          },
          {
            path: "bin",
            component: RecycleBin,
            name: "RecycleBin",
            meta: { title: "回收站", hidden: true },
          },
        ],
      },
      {
        path: "/interaction",
        name: "Interaction",
        meta: { title: "互动管理", icon: "InteractionOutlined" },
        children: [
          {
            path: "comment",
            component: Comment,
            name: "Comment",
            meta: { title: "评论管理", icon: "MessageOutlined" },
          },
          {
            path: "like",
            component: Like,
            name: "Like",
            meta: { title: "点赞管理", icon: "LikeOutlined" },
          },
        ],
      },
      {
        path: "/analyse",
        name: "Analyse",
        meta: { title: "数据分析", icon: "LineChartOutlined" },
        children: [
          {
            path: "posts",
            component: ArticleStatistics,
            name: "ArticleStatistics",
            meta: { title: "文章统计", icon: "FileTextOutlined" },
          },
          {
            path: "users",
            component: UserStatistics,
            name: "UserStatistics",
            meta: { title: "用户统计", icon: "UserOutlined" },
          },
          {
            path: "interaction",
            component: InteractionStatistics,
            name: "InteractionStatistics",
            meta: { title: "互动统计", icon: "InteractionOutlined" },
          },
        ],
      },
      {
        path: "/inform",
        name: "Inform",
        meta: { title: "通知中心", icon: "BellOutlined" },
        children: [
          {
            path: "system",
            component: SystemNotifications,
            name: "SystemNotifications",
            meta: { title: "系统通知", icon: "SettingOutlined" },
          },
          {
            path: "user",
            component: UserNotifications,
            name: "UserNotifications",
            meta: { title: "用户通知", icon: "UserOutlined" },
          },
        ],
      },
      {
        path: "/system",
        name: "System",
        meta: { title: "系统设置", icon: "SettingOutlined" },
        children: [
          {
            path: "",
            component: SystemSettings,
            name: "SystemSettings",
            meta: { title: "系统设置", icon: "SettingOutlined" },
          },
        ],
      },
      {
        path: "/test",
        component: Test,
        name: "Test",
        meta: { title: "测试", icon: "ExperimentOutlined" },
      },
    ],
  },
  {
    path: "/login",
    component: Login,
    name: "Login",
  },

  {
    path: "/error",
    redirect: { name: "404" },
    children: errorView,
  },
  {
    path: "/:pathMatch(.*)",
    redirect: { name: "404" },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
