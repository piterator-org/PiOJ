<template>
  <nav class="navbar navbar-expand-md navbar-light sticky-top bg-white shadow p-3">
    <div class="container">
      <a class="navbar-brand" href="#">PiOJ</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarCollapse"
        aria-controls="navbarCollapse"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav me-auto mb-2 mb-md-0">
          <li class="nav-item">
            <router-link :to="{ name: 'home' }" class="nav-link active">Home</router-link>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Problems</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Contests</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Submissions</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Discussions</a>
          </li>
        </ul>
        <ul class="navbar-nav mb-2 mb-md-0 d-flex">
          <!-- not acthenticated -->
          <li class="nav-item" v-if="!this.$store.state.user">
            <a class="nav-link" href="#">Login</a>
          </li>
          <li class="nav-item" v-if="!this.$store.state.user">
            <a class="nav-link" href="#">Signup</a>
          </li>
          <!-- acthenticated -->
          <li class="nav-item dropdown" v-if="this.$store.state.user">
            <a
              class="nav-link dropdown-toggle"
              href="#"
              id="navbarScrollingDropdown"
              role="button"
              data-bs-toggle="dropdown"
              data-bs-auto-close="outside"
              aria-expanded="false"
              >{{ this.$store.state.user.username }}&nbsp;</a
            >
            <div
              class="dropdown-menu dropdown-menu-end p-4 text-muted border-0 shadow rounded-mx"
              aria-labelledby="navbarScrollingDropdown"
              id="user-card"
            >
              <a href="#">
                <img
                  :src="this.$store.state.user.avatar"
                  class="rounded-circle mx-auto d-none d-md-block border border-1 shadow-sm mb-2"
                  alt="avatar"
                  width="80"
                  height="80"
                />
              </a>
              <p class="text-center d-none d-md-block">
                <span class="fs-5">{{ this.$store.state.user.username }}</span
                >&nbsp;<span class="badge rounded-mx-pill bg-secondary align-text-bottom">{{
                  this.$store.state.user.role
                }}</span>
              </p>
              <ul class="list-group list-group-flush text-center">
                <li class="list-group-item">
                  <a class="text-decoration-none link-secondary" href="#">Profile</a>
                </li>
                <li class="list-group-item">
                  <a class="text-decoration-none link-secondary" href="#">Submissions</a>
                </li>
                <li class="list-group-item">
                  <a class="text-decoration-none link-secondary" href="#">Activities</a>
                </li>
                <li class="list-group-item">
                  <a class="text-decoration-none link-secondary" href="#">Teams</a>
                </li>
                <li class="list-group-item">
                  <a class="text-decoration-none link-secondary" href="#">Blog</a>
                </li>
              </ul>
              <div class="text-center mt-2">
                <a class="text-decoration-none link-secondary" href="#">Switch</a>
                |
                <a class="text-decoration-none link-secondary" href="#">Logout</a>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<style scoped>
#user-card {
  width: 15em;
}
</style>
