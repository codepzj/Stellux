import router, { routes } from "./router";
import { useUserStore } from "@/store/user";
import { useTabStore } from "@/store/tab";
import { useSidebarStore } from "@/store/sidebar";
import type { RouteRecordRaw } from "vue-router";

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

router.beforeEach((to, _from, next) => {
  const sidebarStore = useSidebarStore();

  // 扁平化所有可见的路由（排除隐藏的路由）
  const getVisibleRoutes = (routes: RouteRecordRaw[]): RouteRecordRaw[] => {
    const visibleRoutes: RouteRecordRaw[] = [];
    routes.forEach(route => {
      if (route.meta?.hidden !== true && route.name) {
        visibleRoutes.push(route);
      }
      if (route.children) {
        visibleRoutes.push(...getVisibleRoutes(route.children));
      }
    });
    return visibleRoutes;
  };

  const visibleRoutes = getVisibleRoutes(routes);
  const matchedRoute = visibleRoutes.find(route => route.name === to.name);

  if (matchedRoute && to.name) {
    // 如果路由匹配到侧边栏项，则高亮它
    sidebarStore.setSelectedKeys([to.name.toString()]);

    // 根据路由层次结构确定打开的键
    const pathSegments = to.path.split("/").filter(Boolean);
    const openKeys: string[] = [];
    let currentPath = "";
    const mainRoute = routes.find(r => r.path === "/");

    pathSegments.forEach(segment => {
      currentPath += `/${segment}`;
      const parentRoute = mainRoute?.children?.find(
        r => r.path === currentPath && r.name && r.meta?.hidden !== true
      );
      if (parentRoute?.name) {
        openKeys.push(parentRoute.name.toString());
      }
    });

    // 如果最后一个段是叶节点（选中的项），则移除它
    if (
      openKeys.length > 0 &&
      openKeys[openKeys.length - 1] === to.name.toString()
    ) {
      openKeys.pop();
    }
    sidebarStore.setOpenKeys(openKeys);
  } else {
    // 如果没有匹配（例如，/login, /error），则折叠侧边栏并清除选择
    sidebarStore.setSelectedKeys([]);
    sidebarStore.setOpenKeys([]);
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
