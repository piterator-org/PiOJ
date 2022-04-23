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
import { defineComponent } from 'vue';
import NoticeBoard from '@/components/home/NoticeBoard.vue';
import RecentDynamics from '@/components/home/RecentDynamics.vue';
import SearchAndLinks from '@/components/home/SearchAndLinks.vue';

export default defineComponent({
  name: 'HomeView',
  components: {
    NoticeBoard,
    RecentDynamics,
    SearchAndLinks,
  },
});
</script>
