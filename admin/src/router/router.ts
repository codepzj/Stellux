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
import PublishArticle from "@/views/content/posts/index.vue";
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
      { path: "/", component: Dashboard, name: "Dashboard" },
      {
        path: "/user/list",
        component: UserList,
        name: "UserList",
      },
      {
        path: "/user/role",
        component: UserRole,
        name: "UserRole",
      },
      {
        path: "/content/posts",
        component: PublishArticle,
        name: "PublishArticle",
      },
      {
        path: "/content/list",
        component: ArticleList,
        name: "ArticleList",
      },
      {
        path: "/content/draft",
        component: Draft,
        name: "Draft",
      },
      {
        path: "/content/assets",
        component: MaterialLibrary,
        name: "MaterialLibrary",
      },
      {
        path: "/content/bin",
        component: RecycleBin,
        name: "RecycleBin",
      },
      {
        path: "/interaction/comment",
        component: Comment,
        name: "Comment",
      },
      {
        path: "/interaction/like",
        component: Like,
        name: "Like",
      },
      {
        path: "/interaction/message",
        component: Message,
        name: "Message",
      },
      {
        path: "/analyse/posts",
        component: ArticleStatistics,
        name: "ArticleStatistics",
      },
      {
        path: "/analyse/users",
        component: UserStatistics,
        name: "UserStatistics",
      },
      {
        path: "/analyse/interaction",
        component: InteractionStatistics,
        name: "InteractionStatistics",
      },
      {
        path: "/inform/system",
        component: SystemNotifications,
        name: "SystemNotifications",
      },
      {
        path: "/inform/user",
        component: UserNotifications,
        name: "UserNotifications",
      },
      {
        path: "/system/",
        component: SystemSettings,
        name: "SystemSettings",
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
