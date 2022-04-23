<template>
  <div class="container">
    <section class="mt-4 mb-4">
      <h2>A+B Problem</h2>
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
              <span class="badge rounded-pill bg-primary align-text-bottom">admin</span>
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
            <div class="ms-auto">1.0 Second</div>
          </div>
          <div class="d-flex p-1">
            <strong>Memory Limit</strong>
            <div class="ms-auto">128 MiB</div>
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
        <div class="p-4 bg-white rounded-mx shadow mb-3">
          <h5 class="mb-3">Background</h5>
          Under the requirement of most OI contests, it's not allowed to output any extra things
          including "Please input integer A and B" which are messages to lead user to input data.
          <br />
          If your output contain these things, your program will be judged as a
          <code>Wrong Answer</code>, which will be shown as <code>WA</code>. <br />
          Before comparing your output to the standard output, PiOJ will ignore the spaces at the
          end of each lines and the linefeed at the last line in the output. <br />
          Thus, if it is happend to be that it seemed to be AC on your computer but WA on PiOJ,
          please DO NOT consider it as a mistake of PiOJ, it will probably be some extra data in you
          output.
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3">
          <h5 class="mb-3">Description</h5>
          Blah blah blah.
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3">
          <h5 class="mb-3">Input</h5>
          Blah blah blah.
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3">
          <h5 class="mb-3">Output</h5>
          Blah blah blah.
        </div>
        <div class="row" v-for="(example, index) in problem.examples" :key="example[0]">
          <div class="col-md-6">
            <div class="p-4 bg-white rounded-mx shadow mb-3">
              <h5 class="mb-3">Example Input #{{ index + 1 }}</h5>
              <div class="p-3 bg-white rounded-mx shadow-sm font-monospace position-relative">
                <button class="position-absolute top-0 end-0 btn btn-link btn-sm rounded-mx">
                  Copy
                </button>
                <div class="user-select-all">{{ example[0] }}</div>
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
                <pre class="user-select-all">{{ example[1] }}</pre>
              </div>
            </div>
          </div>
        </div>
        <div class="p-4 bg-white rounded-mx shadow mb-3">
          <h5 class="mb-3">Hints</h5>
          Blah blah blah.
        </div>
      </div>
    </section>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';

export default defineComponent({
  name: 'ProblemDetail',
  data: () => ({ problem: {} }),
  mounted() {
    axios
      .post('/api/problem/get', { id: parseInt(this.$route.params.id as string, 10) })
      .then((res) => {
        this.problem = res.data;
      });
  },
});
</script>
