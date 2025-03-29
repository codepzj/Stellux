import { MdThemes } from "@/types/theme";
import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

type Store = {
  mdTheme: MdThemes;
  setMdTheme: (theme: MdThemes) => void;
};

const useMdThemeStore = create<Store>()(
  persist(
    set => ({
      mdTheme: "react-md",
      setMdTheme: (theme: MdThemes) => set({ mdTheme: theme }),
    }),
    {
      name: "md-theme",
      storage: createJSONStorage(() => localStorage),
    }
  )
);
export default useMdThemeStore;
