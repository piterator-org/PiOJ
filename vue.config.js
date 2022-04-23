const { defineConfig } = require('@vue/cli-service');

module.exports = defineConfig({
  transpileDependencies: true,

  css: {
    loaderOptions: {
      sass: {
        additionalData: '@import "~bootstrap/scss/bootstrap";',
      },
    },
  },

  pluginOptions: {
    i18n: {
      // locale: 'en',
      // fallbackLocale: 'en',
      localeDir: 'locales',
      enableLegacy: false,
      runtimeOnly: true,
      compositionOnly: false,
      fullInstall: true,
    },
  },
});
