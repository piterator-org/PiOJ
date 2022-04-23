<template>
  <div class="row mb-2" id="major-section">
    <div class="col-md-4 mb-3" id="user-info-board">
      <div class="p-5 bg-white rounded-mx shadow">
        <a href="#">
          <img
            :src="this.$store.state.user.avatar"
            class="rounded-circle mx-auto d-block mb-2 border border-1"
            alt="avatar"
            width="100"
            height="100"
          />
        </a>
        <p class="text-center">
          <span class="fs-5">{{ this.$store.state.user.username }}</span>&nbsp;<span
            class="badge rounded-mx-pill bg-secondary align-text-bottom"
            >{{ this.$store.state.user.role }}</span
          >
        </p>
        <div class="text-center">
          <a class="btn btn-outline-primary btn-sm" href="#">My blog</a>
        </div>
      </div>
    </div>
    <div class="col-md-8 mb-3 d-flex" id="notice-board">
      <div class="p-5 bg-white rounded-mx overflow-auto shadow">
        {{ notice }}
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@include media-breakpoint-up(md) {
  #major-section {
    position: relative;
  }

  #notice-board {
    position: absolute;
    width: calc(100% - 33.333333333%);
    height: calc(100% - 1rem);
    height: -webkit-fill-available;
  }

  #user-info-board {
    position: relative;
    left: calc(100% - 33.333333333%);
  }
}
</style>

<script lang="ts">
import { Vue } from 'vue-class-component';

export default class NoticeBoard extends Vue {
  notice = `
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lacinia arcu quam, quis vulputate
    tellus accumsan vel. Morbi gravida arcu vitae ligula congue aliquet. Curabitur accumsan enim
    nibh, sed venenatis libero scelerisque a. Curabitur ac viverra odio. Pellentesque a porttitor
    elit, non hendrerit magna. Aenean porttitor, ipsum ac pharetra volutpat, neque ante ultrices
    dolor, vitae suscipit magna dolor eget libero. Vestibulum ante ipsum primis in faucibus orci
    luctus et ultrices posuere cubilia curae.
  `;
}
</script>
