<template>
  <div class="col-md-8 col-lg-9">
    <div class="container-fuild mb-4">
      <div class="p-4 bg-white rounded-mx shadow">
        <h5 class="mb-1">Discussions</h5>
        <div class="row">
          <div
            class="col-lg-6 col-xxl-4 mt-3"
            v-for="discussion in discussions"
            :key="discussion.id"
          >
            <div class="card">
              <div class="card-body">
                <div class="row">
                  <div class="col-3 col-md-3 d-none d-md-block">
                    <a href="#">
                      <img
                        :src="discussion.user.avatar"
                        class="rounded-circle mx-auto d-block border border-1 discussion-avatar"
                        alt="avatar"
                      />
                    </a>
                  </div>
                  <div class="col-12 col-md-9">
                    <a class="text-decoration-none" href="#">
                      <h5 class="card-title text-ellipsis">
                        {{ discussion.title }}
                      </h5>
                    </a>
                    <p class="card-text">
                      <a class="text-decoration-none" href="#">
                        {{ discussion.user.username }}
                        <span class="badge rounded-mx-pill bg-primary align-text-bottom">{{
                          discussion.user.role
                        }}</span> </a
                      ><br />
                      in
                      <a class="text-decoration-none" href="#">{{ discussion.catagory }}</a
                      ><br />
                      <span class="text-muted"
                        >{{ discussion.pub_date }} &middot; {{ discussion.replies }} replies</span
                      >
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-lg-6 mb-3 d-flex">
        <div class="p-4 bg-white rounded-mx shadow">
          <h5 class="mb-1">Latest Problems</h5>
          <div class="row">
            <div class="col-12 mt-3" v-for="problem in problems" :key="problem.id">
              <div class="card">
                <div class="card-body">
                  <router-link
                    class="text-decoration-none"
                    :to="{ name: 'problem_detail', params: { id: problem.id } }"
                  >
                    <h5 class="card-title text-ellipsis">
                      {{ problem.title }}
                    </h5>
                  </router-link>
                  <p class="card-text">
                    <span class="text-muted">#{{ problem.id }}</span> &middot;
                    {{ problem.difficulty }} &middot; by {{ problem.provider.username }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-lg-6 mb-3 d-flex">
        <div class="p-4 bg-white rounded-mx shadow">
          <h5 class="mb-1">Recent Contests</h5>
          <div class="row">
            <div class="col-12 mt-3" v-for="contest in contests" :key="contest.id">
              <div
                class="card"
                :class="
                  contest.status
                    ? contest.status > 0
                      ? 'border-primary'
                      : 'border-secondary'
                    : 'border-success'
                "
              >
                <div class="card-header fw-bold">
                  <a
                    class="text-decoration-none"
                    :class="
                      contest.status
                        ? contest.status > 0
                          ? 'link-primary'
                          : 'link-secondary'
                        : 'link-success'
                    "
                    href="#"
                  >
                    <span
                      class="badge align-text-bottom"
                      :class="
                        contest.status
                          ? contest.status > 0
                            ? 'bg-primary'
                            : 'bg-secondary'
                          : 'bg-success'
                      "
                      >{{
                        contest.status ? (contest.status > 0 ? 'Pending' : 'Ended') : 'In progress'
                      }}</span
                    >
                    {{ contest.name }}
                  </a>
                </div>
                <div class="card-body">
                  <p class="card-text">
                    <span class="text-muted">{{ contest.id }}</span> &middot; by
                    <a class="text-decoration-none" href="#">
                      {{ contest.organizer }}
                      <span class="badge rounded-mx-pill bg-primary align-text-bottom">{{
                        contest.type
                      }}</span> </a
                    ><br />
                    {{ contest.date }} &middot; {{ contest.duration }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@import "~bootstrap/scss/functions";
@import "~bootstrap/scss/variables";
@import "~bootstrap/scss/mixins";

@include media-breakpoint-up(md) {
  .discussion-avatar {
    width: 90px;
    height: 90px;
  }
}

@include media-breakpoint-up(lg) {
  .discussion-avatar {
    width: 64px;
    height: 64px;
  }
}

@include media-breakpoint-up(xl) {
  .discussion-avatar {
    width: 76px;
    height: 76px;
  }
}

@include media-breakpoint-up(xxl) {
  .discussion-avatar {
    width: 60px;
    height: 60px;
  }
}

.text-ellipsis {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
</style>

<script lang="ts">
import { defineComponent } from 'vue';

const wbh07 = {
  username: 'wbh07',
  avatar: 'https://q1.qlogo.cn/g?b=qq&nk=3415751684&s=640',
  role: 'admin',
};
const wxh06 = {
  username: 'wxh06',
  avatar: 'https://q1.qlogo.cn/g?b=qq&nk=1659133940&s=640',
  role: 'admin',
};
const bohanjun = {
  username: 'bohanjun',
  avatar: 'https://q1.qlogo.cn/g?b=qq&nk=3416885501&s=640',
  role: 'admin',
};
const ForkKILLET = {
  username: 'ForkKILLET',
  avatar: 'https://q1.qlogo.cn/g?b=qq&nk=1096694717&s=640',
};

export default defineComponent({
  name: 'RecentDynamics',
  data: () => ({
    discussions: [
      {
        id: 4,
        title: 'Long Discussion Title',
        catagory: 'Academic',
        pub_date: '9:40 Apr 21',
        replies: 81,
        user: wbh07,
      },
      {
        id: 3,
        title: 'Another Discussion Title',
        catagory: 'Entertainment',
        pub_date: '9:25 Apr 21',
        replies: 12,
        user: wxh06,
      },
      {
        id: 2,
        title: 'Short Two',
        catagory: 'Academic',
        pub_date: '8:53 Apr 21',
        replies: 0,
        user: bohanjun,
      },
      {
        id: 1,
        title: 'Short One',
        catagory: 'Entertainment',
        pub_date: '8:07 Apr 21',
        replies: 51,
        user: ForkKILLET,
      },
    ],
    problems: [
      {
        id: 4,
        title: 'Problem Title 4',
        difficulty: 'Easy',
        provider: wbh07,
      },
      {
        id: 3,
        title: 'Problem Title 3',
        difficulty: 'Master',
        provider: wxh06,
      },
      {
        id: 2,
        title: 'Problem Title 2',
        difficulty: 'Terrible',
        provider: wbh07,
      },
      {
        id: 1,
        title: 'Problem Title 1',
        difficulty: 'Noob',
        provider: wbh07,
      },
    ],
    contests: [
      {
        id: 'pi-round-3',
        name: 'Pi Round III',
        status: 1,
        organizer: 'piterator-org',
        type: 'official',
        date: '9:40 04/21/2022',
        duration: '3 hours 30 minutes',
      },
      {
        id: 'pi-round-2',
        name: 'Pi Round II',
        status: 0,
        organizer: 'piterator-org',
        type: 'official',
        date: '15:20 04/15/2022',
        duration: '4 hours',
      },
      {
        id: 'af-2022',
        name: 'Completed April Fools 2022',
        status: -1,
        organizer: 'april-fools-team',
        date: '0:00 04/01/2022',
        duration: '8 hours',
      },
    ],
  }),
});
</script>
