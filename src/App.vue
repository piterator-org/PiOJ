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
import { Options, Vue } from 'vue-class-component';

@Options({
  components: {
    SiteHeader,
    SiteFooter,
  },
})
export default class App extends Vue {}
</script>

<style lang="scss">
.rounded-mx {
  border-radius: 0.5em;
}
</style>
