import { defineConfig } from '@rsbuild/core';
import { pluginReact } from '@rsbuild/plugin-react';
import { UnoCSSRspackPlugin } from '@unocss/webpack/rspack';

export default defineConfig({
  output: {
    assetPrefix: '/web',
    distPath: {
      root: "./dist/web",
    }
  },
  tools: {
    rspack: {
      plugins: [ UnoCSSRspackPlugin() ],
    },
  },
  plugins: [pluginReact()],
});
