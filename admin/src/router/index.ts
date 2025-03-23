import router from "./router";
import { useUserStore } from "@/store/user";
import { useTabStore } from "@/store/tab";

// 监听标签栏
router.beforeEach((to, _from, next) => {
  const tabStore = useTabStore();
  if (to.meta.title) {
    tabStore.addTab({
      key: to.name as string,
      title: to.meta.title as string,
    });
  }
  next();
});

// 监听用户是否登录
router.beforeEach((to, _from, next) => {
  const userStore = useUserStore();
  if (to.path === "/login") {
    next();
  } else {
    if (userStore.token) {
      next();
    } else {
      next({ path: "/login" });
    }
  }
});

export default router;
