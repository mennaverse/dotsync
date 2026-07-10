import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { tanstackRouter } from "@tanstack/router-plugin/vite";
import UnoCSS from "unocss/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    tanstackRouter({
      target: "react",
      autoCodeSplitting: true,
      generatedRouteTree: "./src/route-tree.gen.ts",
      quoteStyle: "double",
    }),
    UnoCSS(),
    react(),
  ],

  resolve: { tsconfigPaths: true },

  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8067",
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
