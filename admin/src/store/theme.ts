import { defineStore } from "pinia";
import { ref } from "vue";

export const useThemeStore = defineStore(
  "theme",
  () => {
    const tailwindTheme = ref("light");

    const toggleTheme = () => {
      tailwindTheme.value = tailwindTheme.value === "light" ? "dark" : "light";
      document.documentElement.classList.toggle(
        "dark",
        tailwindTheme.value === "dark"
      );
    };

    const initTheme = () => {
      if (tailwindTheme.value === "dark")
        document.documentElement.classList.add("dark");
      else document.documentElement.classList.remove("dark");
    };

    return { tailwindTheme, toggleTheme, initTheme };
  },
  {
    persist: true,
  }
);

export default useThemeStore;
