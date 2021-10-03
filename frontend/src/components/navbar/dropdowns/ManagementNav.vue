<template>
  <div class="navbar-item has-dropdown is-hoverable" @mouseover="isOpen = true">
    <a class="navbar-link">
      Manage
    </a>
    <transition name="fade" apear>
      <div class="navbar-dropdown" v-if="isOpen">
        <router-link :to="{name: 'Clients'}" class="navbar-item" @click="showDrop">Client Management</router-link>

<!--          <router-link class="navbar-link" to="/services/design">Design</router-link>-->

<!--          <router-link class="navbar-link" to="/services/videos">Videos</router-link>-->
      </div>
    </transition>
  </div>
</template>

<script>
import {ref, reactive} from 'vue'
export default {
  name: "ManagementNav",
  setup() {
    const isOpen = ref(false)

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