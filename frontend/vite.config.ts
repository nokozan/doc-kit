import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "tailwindcss";
import autoprefixer from "autoprefixer";

import * as path from "path";
// import alias from "@rollup/plugin-alias";
// import { resolve } from "path";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    // alias({
    //   entries: [{ find: "@", replacement: resolve(__dirname, "src") }],
    // }),
  ],
  css: {
    postcss: {
      plugins: [tailwindcss, autoprefixer],
    },
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"), // Adjust the path as necessary
    },
  },
});
