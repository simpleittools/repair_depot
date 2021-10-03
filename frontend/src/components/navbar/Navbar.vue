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