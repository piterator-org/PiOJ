<template>
  <div class="col-md-3 mb-3 d-flex">
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
          <button class="btn btn-outline-success mb-1" type="submit">
            Search
          </button>
        </form>
      </div>
      <h5 class="mb-3">Links</h5>
      <div class="mb-3">
        <div>
          <a class="text-decoration-none" href="#">Link 1</a>&nbsp;<span
            class="text-muted"
            >Description.</span
          >
        </div>
        <div>
          <a class="text-decoration-none" href="#">Link 2</a>&nbsp;<span
            class="text-muted"
            >Another description.</span
          >
        </div>
        <div>
          <a class="text-decoration-none" href="#">Link 3</a>&nbsp;<span
            class="text-muted"
            >A very very long description that it needs to be wrapped.</span
          >
        </div>
      </div>
    </div>
  </div>
</template>
