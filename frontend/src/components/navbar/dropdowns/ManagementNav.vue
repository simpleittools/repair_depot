<template>
  <button class="flex flex-row items-center w-full px-4 py-4 mt-2 text-base font-bold text-left uppercase bg-transparent rounded-lg md:w-auto md:inline md:mt-0 md:ml-4 focus:outline-none font-montserrat" @mouseover="isOpen = true">
    <span>Manage</span>
  </button>
  <div class="absolute z-10 hidden bg-grey-400 group-hover:block">

    <div class="px-2 pt-2 pb-4 bg-white bg-gray-200 shadow-lg" v-if="isOpen">
      <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
        <p>
          <router-link :to="{name: 'Clients'}" class="navbar-item" @click="showDrop">Client Management</router-link>
        </p>
      </div>
    </div>
  </div>

</template>

<script>
import {ref, reactive} from 'vue'
export default {
  name: "ManagementNav",
  setup() {
    const isOpen = ref(true)

    const showDrop = () => {
      isOpen.value = !isOpen.value
    }


    return {
      isOpen,
      showDrop
    }
  }
}


</script>

<style scoped>

</style>