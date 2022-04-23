<template>
  <div class="bg-light">
    <SiteHeader />
    <router-view />
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
.rounded-mx {
  border-radius: 0.5em;
}
</style>
