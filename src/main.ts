import { createApp } from "vue";
import * as bootstrap from "bootstrap";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import i18n from "./i18n";

const app = createApp(App).use(store).use(i18n).use(router);
app.config.globalProperties.$bootstrap = bootstrap;
app.mount("#app");

declare module "@vue/runtime-core" {
  interface ComponentCustomProperties {
    $bootstrap: typeof bootstrap;
    $store: typeof store;
  }
}
