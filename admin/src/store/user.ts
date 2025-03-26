import { defineStore } from "pinia";

import type { User } from "@/types/user";
import { ref } from "vue";

export const useUserStore = defineStore(
  "user",
  () => {
    const user = ref<User | null>(null);
    const token = ref<string | null>(null);
    const setUser = (newUser: User) => {
      user.value = newUser;
    };

    const setToken = (newToken: string) => {
      token.value = newToken;
    };

    const clearToken = () => {
      token.value = null;
    };

    return { user, token, setUser, setToken, clearToken };
  },
  {
    persist: true,
  }
);
