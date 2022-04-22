<template>
  <main class="pt-4">
    <div class="container">
      <NoticeBoard />
      <div class="row mb-2">
        <RecentDynamics />
        <SearchAndLinks />
      </div>
    </div>
  </main>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import NoticeBoard from '@/components/home/NoticeBoard.vue';
import RecentDynamics from '@/components/home/RecentDynamics.vue';
import SearchAndLinks from '@/components/home/SearchAndLinks.vue';

@Options({
  components: {
    NoticeBoard,
    RecentDynamics,
    SearchAndLinks,
  },
})
export default class HomeView extends Vue {}
</script>
