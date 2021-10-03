<template>
  <nav class="navbar" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <a class="navbar-item" href="https://bulma.io">
        <img src="https://bulma.io/images/bulma-logo.png" width="112" height="28">
      </a>

      <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>

    <div id="navbarBasicExample" class="navbar-menu">
      <div class="navbar-start">
        <a class="navbar-item">
          <router-link :to="{name: 'Home'}">Home</router-link>
        </a>

        <a class="navbar-item">
          Documentation
        </a>
        <ManagementNav />
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link">
            Dashboard
          </a>
          <div class="navbar-dropdown">
            <a href="http://localhost:3010/dashboard/status" target="_blank" class="navbar-item">
              Status
            </a>
            <a class="navbar-item">
              Jobs
            </a>
            <a class="navbar-item">
              Contact
            </a>
            <hr class="navbar-divider">
            <a class="navbar-item">
              Report an issue
            </a>
          </div>
          <div class="navbar-dropdown">
            <a href="http://localhost:3010/dashboard/status" target="_blank" class="navbar-item">
              Status
            </a>
            <a class="navbar-item">
              Jobs
            </a>
            <a class="navbar-item">
              Contact
            </a>
            <hr class="navbar-divider">
            <a class="navbar-item">
              Report an issue
            </a>
          </div>
        </div>
      </div>

      <div class="navbar-end">
        <div class="navbar-item">
          <div class="buttons">
            <a class="button is-primary">
              <strong>Sign up</strong>
            </a>
            <a class="button is-light">
              Log in
            </a>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>



<script>
import ManagementNav from "@/components/navbar/dropdowns/ManagementNav";
export default {
  name: "Navbar",
  components: {
    ManagementNav
  }
}
</script>

<style scoped>

</style>