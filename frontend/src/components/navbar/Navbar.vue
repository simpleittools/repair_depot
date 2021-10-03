<template>
  <nav aria-label="primary" class="relative z-20 flex-col flex-grow hidden pb-4 md:pb-0 md:flex md:flex-row ml-4">
    <div class="flex items-center flex-shrink-0 mr-6 mt-2 font-montserrat">
      <svg class="fill-current h-8 w-8 mr-2" width="54" height="54" viewBox="0 0 54 54" xmlns="http://www.w3.org/2000/svg"><path d="M13.5 22.1c1.8-7.2 6.3-10.8 13.5-10.8 10.8 0 12.15 8.1 17.55 9.45 3.6.9 6.75-.45 9.45-4.05-1.8 7.2-6.3 10.8-13.5 10.8-10.8 0-12.15-8.1-17.55-9.45-3.6-.9-6.75.45-9.45 4.05zM0 38.3c1.8-7.2 6.3-10.8 13.5-10.8 10.8 0 12.15 8.1 17.55 9.45 3.6.9 6.75-.45 9.45-4.05-1.8 7.2-6.3 10.8-13.5 10.8-10.8 0-12.15-8.1-17.55-9.45-3.6-.9-6.75.45-9.45 4.05z"/></svg>
      <router-link :to="{name: 'Home'}"><span class="font-semibold text-xl tracking-tight text-base">Home</span></router-link>
    </div>
    <div class="relative group">
      <ManagementNav />

    </div>

    <div class="relative group">
      <button class="flex flex-row items-center w-full px-4 py-4 mt-2 text-base font-bold text-left uppercase bg-transparent rounded-lg md:w-auto md:inline md:mt-0 md:ml-4 focus:outline-none font-montserrat">
        <span>Second Dropdown</span>
      </button>
      <div class="absolute z-10 hidden bg-grey-200 group-hover:block">

        <div class="px-2 pt-2 pb-4 bg-white bg-gray-200 shadow-lg">
          <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
            <p>
              stuff
            </p>
          </div>
        </div>
      </div>
    </div>
  </nav>
<!--  <nav class="navbar" role="navigation" aria-label="main navigation">-->
<!--    <div class="navbar-brand">-->
<!--      <a class="navbar-item" href="https://bulma.io">-->
<!--        <img src="https://bulma.io/images/bulma-logo.png" width="112" height="28">-->
<!--      </a>-->

<!--      <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">-->
<!--        <span aria-hidden="true"></span>-->
<!--        <span aria-hidden="true"></span>-->
<!--        <span aria-hidden="true"></span>-->
<!--      </a>-->
<!--    </div>-->

<!--    <div id="navbarBasicExample" class="navbar-menu">-->
<!--      <div class="navbar-start">-->
<!--        <a class="navbar-item">-->
<!--          <router-link :to="{name: 'Home'}">Home</router-link>-->
<!--        </a>-->

<!--        <a class="navbar-item">-->
<!--          Documentation-->
<!--        </a>-->
<!--        <ManagementNav />-->
<!--        <div class="navbar-item has-dropdown is-hoverable">-->
<!--          <a class="navbar-link">-->
<!--            Dashboard-->
<!--          </a>-->
<!--          <div class="navbar-dropdown">-->
<!--            <a href="http://localhost:3010/dashboard/status" target="_blank" class="navbar-item">-->
<!--              Status-->
<!--            </a>-->
<!--            <a class="navbar-item">-->
<!--              Jobs-->
<!--            </a>-->
<!--            <a class="navbar-item">-->
<!--              Contact-->
<!--            </a>-->
<!--            <hr class="navbar-divider">-->
<!--            <a class="navbar-item">-->
<!--              Report an issue-->
<!--            </a>-->
<!--          </div>-->
<!--          <div class="navbar-dropdown">-->
<!--            <a href="http://localhost:3010/dashboard/status" target="_blank" class="navbar-item">-->
<!--              Status-->
<!--            </a>-->
<!--            <a class="navbar-item">-->
<!--              Jobs-->
<!--            </a>-->
<!--            <a class="navbar-item">-->
<!--              Contact-->
<!--            </a>-->
<!--            <hr class="navbar-divider">-->
<!--            <a class="navbar-item">-->
<!--              Report an issue-->
<!--            </a>-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->

<!--      <div class="navbar-end">-->
<!--        <div class="navbar-item">-->
<!--          <div class="buttons">-->
<!--            <a class="button is-primary">-->
<!--              <strong>Sign up</strong>-->
<!--            </a>-->
<!--            <a class="button is-light">-->
<!--              Log in-->
<!--            </a>-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->
<!--  </nav>-->
<!--  <nav class="flex items-center justify-between flex-wrap bg-teal-500 p-6">-->
<!--    <div class="flex items-center flex-shrink-0 mr-6">-->
<!--      <svg class="fill-current h-8 w-8 mr-2" width="54" height="54" viewBox="0 0 54 54" xmlns="http://www.w3.org/2000/svg"><path d="M13.5 22.1c1.8-7.2 6.3-10.8 13.5-10.8 10.8 0 12.15 8.1 17.55 9.45 3.6.9 6.75-.45 9.45-4.05-1.8 7.2-6.3 10.8-13.5 10.8-10.8 0-12.15-8.1-17.55-9.45-3.6-.9-6.75.45-9.45 4.05zM0 38.3c1.8-7.2 6.3-10.8 13.5-10.8 10.8 0 12.15 8.1 17.55 9.45 3.6.9 6.75-.45 9.45-4.05-1.8 7.2-6.3 10.8-13.5 10.8-10.8 0-12.15-8.1-17.55-9.45-3.6-.9-6.75.45-9.45 4.05z"/></svg>-->
<!--      <span class="font-semibold text-xl tracking-tight">Tailwind CSS</span>-->
<!--    </div>-->
<!--    <div class="block lg:hidden">-->
<!--      <button class="flex items-center px-3 py-2 border rounded text-teal-200 border-teal-400 hover:text-white hover:border-white">-->
<!--        <svg class="fill-current h-3 w-3" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>Menu</title><path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"/></svg>-->
<!--      </button>-->
<!--    </div>-->
<!--    <div class="w-full block flex-grow lg:flex lg:items-center lg:w-auto">-->
<!--      <div class="text-sm lg:flex-grow">-->
<!--        <a href="#responsive-header" class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4">-->
<!--          Clients-->
<!--        </a>-->
<!--        <a href="#responsive-header" class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4">-->
<!--          Examples-->
<!--        </a>-->
<!--        <a href="#responsive-header" class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white">-->
<!--          Blog-->
<!--        </a>-->
<!--        <ul class="flex border-b">-->
<!--          <li class="-mb-px mr-1">-->
<!--            <a class="bg-white inline-block border-l border-t border-r rounded-t py-2 px-4 text-blue-700 font-semibold" href="#">Home</a>-->
<!--          </li>-->
<!--          <li class="mr-1">-->
<!--            <router-link class="bg-white inline-block py-2 px-4 text-blue-500 hover:text-blue-800 font-semibold" :to="{name: 'Clients'}">Clients</router-link>-->
<!--          </li>-->
<!--          <li class="mr-1">-->
<!--            <a class="bg-white inline-block py-2 px-4 text-blue-500 hover:text-blue-800 font-semibold" href="#">Tab</a>-->
<!--          </li>-->
<!--          <li class="mr-1">-->
<!--            <a class="bg-white inline-block py-2 px-4 text-gray-400 font-semibold" href="#">Tab</a>-->
<!--          </li>-->
<!--        </ul>-->
<!--      </div>-->
<!--      <div>-->
<!--        <a href="#" class="inline-block text-sm px-4 py-2 leading-none border rounded border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0">Download</a>-->
<!--      </div>-->

<!--    </div>-->
<!--  </nav>-->

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