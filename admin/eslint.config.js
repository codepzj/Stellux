import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import pluginVue from "eslint-plugin-vue";


/** @type {import('eslint').Linter.Config[]} */
export default [
  { files: ["**/*.{js,mjs,cjs,ts,vue}"] },
  { languageOptions: { globals: globals.browser } },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  ...pluginVue.configs["flat/essential"],
  {
    files: ["**/*.vue"],
    languageOptions: { parserOptions: { parser: tseslint.parser } },
  },
  {
    rules: {
      "vue/multi-word-component-names": "off", // 禁用 vue/multi-word-component-names 规则
      "@typescript-eslint/no-explicit-any": ["off"], // 禁用 @typescript-eslint/no-explicit-any 规则
    },
  },
];