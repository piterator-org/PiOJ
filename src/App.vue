<template>
  <div class="bg-light">
    <SiteHeader />
    <router-view v-slot="{ Component }">
      <transition name="fade-in">
        <component :is="Component" />
      </transition>
    </router-view>
    <SiteFooter />
  </div>
</template>

<script lang="ts">
import 'bootstrap';
import SiteHeader from '@/components/SiteHeader.vue';
import SiteFooter from '@/components/SiteFooter.vue';
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'HomeView',
  components: {
    SiteHeader,
    SiteFooter,
  },
});
</script>

<style lang="scss">
 ;
</style>

<style>
.rounded-mx {
  border-radius: 0.5em;
}
</style>

<style scoped>
.fade-in-enter-active {
  transition: opacity 0.5s ease;
}

.fade-in-enter-from {
  opacity: 0;
}
.fade-in-leave-to {
  display: none;
}
</style>
