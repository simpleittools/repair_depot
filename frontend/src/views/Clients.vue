<template>
  <button class="button is-primary" :class="{'is-light': !newClient}" @click="newClientButton">New Client</button>
  <button class="button is-primary" :class="{'is-light': !searchClient}" @click="searchClientButton">Search Clients</button>
  <div v-if="newClient">
    <form action="">
      <div class="control">
        <input type="text" class="input" placeholder="First Name">
      </div>
    </form>
  </div>
  <div v-else-if="searchClient">
    Search Client Form
  </div>

</template>

<script>
import {ref} from "vue";

export default {
  name: "Clients",
  setup() {
    const isActive = ref(false)
    const newClient = ref(true)
    const searchClient = ref(false)

    const newClientButton = () => {
      newClient.value = true
      searchClient.value = false
    }
    const searchClientButton = () => {
      newClient.value = false
      searchClient.value = true
    }


    return {
      isActive,
      newClient,
      searchClient,
      newClientButton,
      searchClientButton



    }
  }
}
</script>

<style scoped>

</style>