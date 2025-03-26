<template>
  <div
    class="flex items-center justify-center h-screen w-full dark:bg-gray-900"
  >
    <a-card class="shadow-md w-96">
      <img
        class="m-auto my-6 w-32"
        :src="themeStore.tailwindTheme === 'dark' ? darkLogo : lightLogo"
      />
      <a-form
        :model="loginForm"
        class="w-full mx-auto pt-4 px-4"
        @finish="onFinish"
      >
        <a-form-item
          label="用户名"
          name="username"
          :rules="[{ required: true, message: '请输入用户名' }]"
        >
          <a-input v-model:value="loginForm.username">
            <template #prefix>
              <UserOutlined class="site-form-item-icon" />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item
          label="密码"
          name="password"
          :rules="[{ required: true, message: '请输入密码' }]"
        >
          <a-input-password v-model:value="loginForm.password">
            <template #prefix>
              <LockOutlined class="site-form-item-icon" />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item>
          <a-form-item name="remember" no-style>
            <a-checkbox class="text-sm" v-model:checked="loginForm.remember"
              >记住我</a-checkbox
            >
          </a-form-item>
          <span class="float-right space-x-2 w-32">
            <a class="text-sm" href="">忘记密码</a>
            <a class="text-sm" href="">立即注册</a>
          </span>
        </a-form-item>

        <a-form-item class="flex justify-center">
          <a-button
            :disabled="disabled"
            type="primary"
            html-type="submit"
            class="w-48"
          >
            登录
          </a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>
<script lang="ts" setup>
import { reactive, computed } from "vue";
import { UserOutlined, LockOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import type { LoginForm, LoginVO } from "@/types/user";
import { userLogin } from "@/api/user";
import { useUserStore } from "@/store/user";
import darkLogo from "@/assets/logo/logo-dark.png";
import lightLogo from "@/assets/logo/logo-light.png";
import type { Response } from "@/types/response";
import { useRouter } from "vue-router";
import { useThemeStore } from "@/store/theme";

const themeStore = useThemeStore();
const { setUser, setToken } = useUserStore();
const router = useRouter();

const loginForm: LoginForm = reactive({
  username: "",
  password: "",
  remember: true,
});

// 登录校验逻辑
const onFinish = async (loginForm: LoginForm) => {
  const key = "updatable";
  try {
    message.loading({ content: "登录中", key });
    const res: Response<LoginVO> = await userLogin(loginForm);
    // 设置用户信息和token
    setUser(res.data.user);
    setToken(res.data.token);
    message.success({ content: "登录成功", key });
    router.push("/");
  } catch (error: any) {
    message.error({ content: error, key });
  }
};

const disabled = computed(() => {
  return !(loginForm.username && loginForm.password);
});
</script>
