<template>
  <div class="col-md-4 col-lg-3 mb-3 d-flex">
    <div class="p-4 bg-white rounded-mx shadow">
      <h5 class="mb-3">Search</h5>
      <div class="mb-3">
        <form class="d-block d-xl-flex" role="search">
          <input
            class="form-control me-2 mb-1"
            type="search"
            placeholder="Search"
            aria-label="Search"
          />
          <button class="btn btn-outline-success mb-1" type="submit">Search</button>
        </form>
      </div>
      <h5 class="mb-3">Links</h5>
      <div class="mb-3">
        <div v-for="link in links" :key="link.name">
          <a class="text-decoration-none" href="#">{{ link.name }}</a>
          <span class="text-muted">{{ link.description }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'SearchAndLinks',
  data: () => ({
    links: [
      {
        name: 'Link 1',
        description: 'Description.',
      },
      {
        name: 'Link 2',
        description: 'Another description.',
      },
      {
        name: 'Link 3',
        description: 'A very very long description that it needs to be wrapped.',
      },
    ],
  }),
});
</script>
