<template>
  <footer class="bg-white mt-5 shadow">
    <div class="container py-5">
      <div class="row">
        <div class="col-12 col-md">
          <h5>PiOJ <span class="text-muted fs-6">Next Generation</span></h5>
          Powered by <a class="link-secondary" href="#">PiOJ</a>
          <small class="d-block mb-3 text-muted">
            Release v1.0.0 <br />
            &copy; 2022 Piterator <br />
            沪ICP备10000000号 <br />
            沪公网安备31010100000000号 <br />
          </small>
        </div>
        <div class="col-6 col-md">
          <h5>Social</h5>
          <ul class="list-unstyled text-small">
            <li><a class="link-secondary" href="#">Twitter</a></li>
            <li><a class="link-secondary" href="#">Facebook</a></li>
            <li><a class="link-secondary" href="#">YouTube</a></li>
            <li><a class="link-secondary" href="#">Bilibili</a></li>
          </ul>
        </div>
        <div class="col-6 col-md">
          <h5>Developers</h5>
          <ul class="list-unstyled text-small">
            <li><a class="link-secondary" href="#">GitHub</a></li>
            <li><a class="link-secondary" href="#">Documentation</a></li>
            <li><a class="link-secondary" href="#">API</a></li>
          </ul>
        </div>
        <div class="col-6 col-md">
          <h5>Supports</h5>
          <ul class="list-unstyled text-small">
            <li><a class="link-secondary" href="#">Help</a></li>
            <li><a class="link-secondary" href="#">QQ Group</a></li>
          </ul>
        </div>
        <div class="col-6 col-md">
          <h5>About</h5>
          <ul class="list-unstyled text-small">
            <li>
              <router-link :to="{ name: 'about' }" class="link-secondary">About</router-link>
            </li>
            <li><a class="link-secondary" href="#">Contact Us</a></li>
            <li><a class="link-secondary" href="#">Privacy</a></li>
            <li><a class="link-secondary" href="#">Terms of Service</a></li>
          </ul>
        </div>
      </div>
    </div>
  </footer>
</template>
