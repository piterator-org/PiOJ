<template>
  <div class="container">
    <section class="mt-4 mb-4">
      <h2>{{ t('title', 'Loading...') }}</h2>
      <p class="fs-5 text-muted">#{{ problem.id }}</p>
    </section>
    <section class="row mb-2">
      <div class="col-md-5 col-lg-4 col-xl-3 mb-3 order-md-2">
        <div class="mb-3">
          <div class="row">
            <div class="col-xl-6 d-grid mb-2">
              <a class="btn btn-primary" href="#">Submit</a>
            </div>
            <div class="col-xl-6 d-grid mb-2">
              <a class="btn btn-secondary" href="#">Submissions</a>
            </div>
            <div class="col-xl-6 d-grid mb-2">
              <a class="btn btn-secondary" href="#">Solutions</a>
            </div>
            <div class="col-xl-6 d-grid mb-2">
              <a class="btn btn-secondary" href="#">Discussions</a>
            </div>
          </div>
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3">
          <div class="d-flex p-1">
            <strong>Provider</strong>
            <a class="ms-auto text-decoration-none" href="#">
              wbh07
              <span class="badge bg-primary align-text-bottom">admin</span>
            </a>
          </div>
          <div class="d-flex p-1">
            <strong>Difficulty</strong>
            <div class="ms-auto">Noob</div>
          </div>
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3">
          <div class="d-flex p-1">
            <strong>Input</strong>
            <div class="ms-auto font-monospace">{{ problem.input_file }}</div>
          </div>
          <div class="d-flex p-1">
            <strong>Output</strong>
            <div class="ms-auto font-monospace">{{ problem.output_file }}</div>
          </div>
          <div class="d-flex p-1">
            <strong>Time Limit</strong>
            <div class="ms-auto">{{ problem.time_limit }}</div>
          </div>
          <div class="d-flex p-1">
            <strong>Memory Limit</strong>
            <div class="ms-auto">{{ problem.memory_limit }}</div>
          </div>
          <div class="d-flex p-1">
            <strong>Submitted</strong>
            <div class="ms-auto">1247</div>
          </div>
          <div class="d-flex p-1">
            <strong>Accepted</strong>
            <div class="ms-auto">996</div>
          </div>
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3">
          <a href="#"><span class="badge bg-success">Some Tag</span></a>
          <a href="#"><span class="badge bg-success">Another Tag</span></a>
          <a href="#"><span class="badge bg-success">Third</span></a>
          <a href="#"><span class="badge bg-success">Also Tnother</span></a>
        </div>
      </div>
      <div class="col-md-7 col-lg-8 col-xl-9 mb-3 order-md-1">
        <div class="p-4 bg-white rounded-mx shadow mb-3" v-if="t('background', ' ').trim()">
          <h5 class="mb-3">Background</h5>
          {{ t('background') }}
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3" v-if="t('description', ' ').trim()">
          <h5 class="mb-3">Description</h5>
          {{ t('description') }}
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3" v-if="t('input_format', ' ').trim()">
          <h5 class="mb-3">Input</h5>
          {{ t('input_format') }}
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3" v-if="t('output_format', ' ').trim()">
          <h5 class="mb-3">Output</h5>
          {{ t('output_format') }}
        </div>
        <div class="row" v-for="(example, index) in problem.examples" :key="example[0]">
          <div class="col-md-6">
            <div class="p-4 bg-white rounded-mx shadow mb-3">
              <h5 class="mb-3">Example Input #{{ index + 1 }}</h5>
              <div class="p-3 bg-white rounded-mx shadow-sm font-monospace position-relative">
                <button class="position-absolute top-0 end-0 btn btn-link btn-sm rounded-mx">
                  Copy
                </button>
                <pre class="user-select-all mb-0">{{ example[0] }}</pre>
              </div>
            </div>
          </div>
          <div class="col-md-6">
            <div class="p-4 bg-white rounded-mx shadow mb-3">
              <h5 class="mb-3">Example Output #{{ index + 1 }}</h5>
              <div class="p-3 bg-white rounded-mx shadow-sm font-monospace position-relative">
                <button class="position-absolute top-0 end-0 btn btn-link btn-sm rounded-mx">
                  Copy
                </button>
                <pre class="user-select-all mb-0">{{ example[1] }}</pre>
              </div>
            </div>
          </div>
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3" v-if="t('hints', ' ').trim()">
          <h5 class="mb-3">Hints</h5>
          {{ t('hints') }}
        </div>
      </div>
    </section>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios, { AxiosResponse } from 'axios';
import { useI18n } from 'vue-i18n';

export default defineComponent({
  name: 'ProblemDetail',
  data: () => ({ problem: {} }),
  setup() {
    return useI18n({
      useScope: 'local',
    });
  },
  created() {
    this.$watch(
      () => this.$route.params,
      () => {
        if (this.$route.name === 'problem_detail') {
          this.fetchData(this.$route.params.id as string);
        }
      },
      { immediate: true },
    );
  },
  methods: {
    fetchData(pid: string) {
      this.problem = {};
      this.availableLocales.forEach((locale) => {
        this.setLocaleMessage(locale, {});
      });
      axios
        .post('/api/problem/get', { id: parseInt(pid, 10) })
        .then((response) => {
          this.problem = response.data;
          const locales: { [key: string]: { [key: string]: string } } = {};
          ['title', 'background', 'description', 'input_format', 'output_format', 'hints']
            .filter((key) => response.data[key])
            .forEach((key) => {
              Object.keys(response.data[key]).forEach((lang) => {
                locales[lang] = locales[lang] || {};
                locales[lang][key] = response.data[key][lang];
              });
            });
          Object.entries(locales).forEach(([locale, message]) => {
            this.setLocaleMessage(locale, message);
          });
        })
        .catch((err) => {
          const { response }: { response: AxiosResponse } = err;
          const { fullPath } = this.$route;
          if (response.status === 404) {
            this.$router
              .replace({ name: '404', params: { pathMatch: '404' } })
              .then(() => window.history.replaceState({}, '', fullPath));
          }
        });
    },
  },
});
</script>
