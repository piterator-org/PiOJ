const { defineConfig } = require('@vue/cli-service');

module.exports = defineConfig({
  transpileDependencies: true,

  chainWebpack: (config) => {
    config.module
      .rule('vue')
      .use('vue-loader')
      .tap((_options) => {
        const options = _options;
        options.compilerOptions = { whitespace: 'preserve' };
        return options;
      });
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
