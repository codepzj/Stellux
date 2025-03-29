<template>
  <div>
    <a-menu
      v-model:selectedKeys="sidebarStore.selectedKeys"
      mode="inline"
      :open-keys="sidebarStore.openKeys"
      :items="menuItems"
      @openChange="onOpenChange"
      @select="onSelect"
      :style="{ height: '100%', borderRight: 0 }"
    ></a-menu>
  </div>
</template>

<script lang="ts" setup>
import { computed, h } from "vue";
import { useRouter } from "vue-router";
import { routes } from "@/router/router";
import { useSidebarStore } from "@/store/sidebar";
import type { ItemType } from "ant-design-vue";
import type { RouteRecordRaw } from "vue-router";
import type { VNodeChild } from "vue";
import * as Icons from "@ant-design/icons-vue"; // Import all icons dynamically

type MenuItemType = {
  key: string;
  label: string;
  children?: MenuItemType[];
  icon?: () => VNodeChild;
  type?: "group";
} & Omit<ItemType, "key" | "label" | "children" | "icon" | "type">;

function getItem(
  label: string,
  key: string,
  icon?: string,
  children?: MenuItemType[],
  type?: "group"
): MenuItemType {
  return {
    key,
    icon:
      icon && Icons[icon as keyof typeof Icons]
        ? () => h(Icons[icon as keyof typeof Icons])
        : undefined,
    children,
    label,
    type,
  };
}

const menuItems = computed(() => {
  const mainRoute = routes.find((r): r is RouteRecordRaw => r.path === "/");
  if (!mainRoute || !mainRoute.children) return [];

  const generateMenuItems = (route: RouteRecordRaw): MenuItemType | null => {
    if (!route.name || !route.meta?.title || route.meta.hidden === true) {
      return null;
    }

    const hasComponent = !!route.component;
    const childrenRoutes = route.children?.filter(
      (child): child is RouteRecordRaw =>
        !!child.name && !!child.meta?.title && child.meta.hidden !== true
    );

    const itemChildren = childrenRoutes
      ?.map(child => generateMenuItems(child))
      .filter((item): item is MenuItemType => item !== null);

    if (!hasComponent && (!itemChildren || itemChildren.length === 0)) {
      return null;
    }

    return getItem(
      route.meta.title as string,
      route.name.toString(),
      route.meta.icon as string | undefined,
      itemChildren?.length ? itemChildren : undefined
    );
  };

  return mainRoute.children
    .map(route => generateMenuItems(route))
    .filter((item): item is MenuItemType => item !== null);
});

const sidebarStore = useSidebarStore();

const rootSubmenuKeys = computed(() =>
  menuItems.value
    .filter((item): item is MenuItemType => !!item.children)
    .map(item => item.key)
);

const router = useRouter();

const onSelect = ({ key }: { key: string }) => {
  sidebarStore.setSelectedKeys([key]);
  router.push({ name: key });
};

const onOpenChange = (openKeys: string[]) => {
  const latestOpenKey = openKeys.find(
    key => sidebarStore.openKeys.indexOf(key) === -1
  );
  if (
    typeof latestOpenKey === "string" &&
    rootSubmenuKeys.value.indexOf(latestOpenKey) === -1
  ) {
    sidebarStore.setOpenKeys(openKeys);
  } else {
    sidebarStore.setOpenKeys(latestOpenKey ? [latestOpenKey] : []);
  }
};
</script>
