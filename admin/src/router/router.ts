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
import PublishArticle from "@/views/content/posts/create.vue";
import ArticleList from "@/views/content/list/index.vue";
import Draft from "@/views/content/draft.vue";
import MaterialLibrary from "@/views/content/material/asset.vue";
import RecycleBin from "@/views/content/bin.vue";
import Comment from "@/views/interaction/comment.vue";
import Like from "@/views/interaction/like.vue";
import Message from "@/views/interaction/private.vue";
import ArticleStatistics from "@/views/analyse/post.vue";
import UserStatistics from "@/views/analyse/user.vue";
import InteractionStatistics from "@/views/analyse/interaction.vue";
import SystemNotifications from "@/views/notification/system.vue";
import UserNotifications from "@/views/notification/user.vue";
import SystemSettings from "@/views/system/index.vue";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: Layout,
    children: [
      {
        path: "/",
        component: Dashboard,
        name: "Dashboard",
        meta: { title: "仪表盘" },
      },
      {
        path: "/user/list",
        component: UserList,
        name: "UserList",
        meta: { title: "用户列表" },
      },
      {
        path: "/user/role",
        component: UserRole,
        name: "UserRole",
        meta: { title: "用户角色" },
      },
      {
        path: "/content/posts",
        component: PublishArticle,
        name: "PublishArticle",
        meta: { title: "发布文章" },
      },
      {
        path: "/content/list",
        component: ArticleList,
        name: "ArticleList",
        meta: { title: "文章列表" },
      },
      {
        path: "/content/draft",
        component: Draft,
        name: "Draft",
        meta: { title: "草稿箱" },
      },
      {
        path: "/content/assets",
        component: MaterialLibrary,
        name: "MaterialLibrary",
        meta: { title: "素材库" },
      },
      {
        path: "/content/bin",
        component: RecycleBin,
        name: "RecycleBin",
        meta: { title: "回收站" },
      },
      {
        path: "/interaction/comment",
        component: Comment,
        name: "Comment",
        meta: { title: "评论管理" },
      },
      {
        path: "/interaction/like",
        component: Like,
        name: "Like",
        meta: { title: "点赞管理" },
      },
      {
        path: "/interaction/message",
        component: Message,
        name: "Message",
        meta: { title: "私信管理" },
      },
      {
        path: "/analyse/posts",
        component: ArticleStatistics,
        name: "ArticleStatistics",
        meta: { title: "文章统计" },
      },
      {
        path: "/analyse/users",
        component: UserStatistics,
        name: "UserStatistics",
        meta: { title: "用户统计" },
      },
      {
        path: "/analyse/interaction",
        component: InteractionStatistics,
        name: "InteractionStatistics",
        meta: { title: "互动统计" },
      },
      {
        path: "/inform/system",
        component: SystemNotifications,
        name: "SystemNotifications",
        meta: { title: "系统通知" },
      },
      {
        path: "/inform/user",
        component: UserNotifications,
        name: "UserNotifications",
        meta: { title: "用户通知" },
      },
      {
        path: "/system/",
        component: SystemSettings,
        name: "SystemSettings",
        meta: { title: "系统设置" },
      },
    ],
  },
  {
    path: "/login",
    component: Login,
    name: "Login",
  },
  {
    path: "/test",
    component: Test,
    name: "Test",
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
