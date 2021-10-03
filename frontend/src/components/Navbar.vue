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
<!--        todo: change this to https://javascript.plainenglish.io/how-to-set-up-a-dropdown-navbar-router-with-vuejs-6e0ef48625f6-->
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link">
            Manage
          </a>
          <div class="navbar-dropdown">
            <router-link :to="{name: 'Clients'}" class="navbar-item">Client Management</router-link>
            <a class="navbar-item" href="/clients">
              Device Management
            </a>
            <hr class="navbar-divider">
            <a class="navbar-item">
              Add Job
            </a>
          </div>
        </div>

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
export default {
  name: "Navbar"
}
</script>

<style scoped>

</style>