<template>
  <ul class="flex border-b">
    <li class="mr-1" :class="{'-mb-px' : newClient}">
      <button class="bg-white inline-block py-2 px-4 text-blue-500 hover:text-blue-800 font-semibold" :class="{'border-l border-t border-r rounded-t border-gray-400 text-blue-700 font-bold': newClient}" @click="newClientButton">New Client</button>
    </li>
    <li class="mr-1" :class="{'-mb-px' : searchClient}">
      <button class="bg-white inline-block py-2 px-4 text-blue-500 hover:text-blue-800 font-semibold" :class="{'border-l border-t border-r rounded-t border-gray-400 text-blue-700 font-bold': searchClient}" @click="searchClientButton">Search Clients</button>
    </li>
    <li class="mr-1">
      <a class="bg-white inline-block py-2 px-4 text-blue-500 hover:text-blue-800 font-semibold" href="#">Tab</a>
    </li>
    <li class="mr-1">
      <a class="bg-white inline-block py-2 px-4 text-gray-400 font-semibold" href="#">Tab</a>
    </li>
  </ul>
  <div v-if="newClient">

      <form class="w-full max-w-md" @submit.prevent="handleSave">
        <div class="flex items-center border-teal-500 py-2">
          <input type="text" class="form-input focus:shadow-outline" v-model="clientDetails.firstName" placeholder="First Name">
        </div>
        <div class="flex items-center border-teal-500 py-2">
          <input type="text" class="form-input focus:shadow-outline" v-model="clientDetails.lastName" placeholder="Last Name">
        </div>
        <div class="flex items-center border-teal-500 py-2">
          <input type="text" class="form-input focus:shadow-outline" v-model="clientDetails.phone1" placeholder="Primary Phone Number">

        </div>
        <div class="flex items-center border-teal-500 py-2">
          <input type="text" class="form-input focus:shadow-outline" v-model="clientDetails.phone2" placeholder="Secondary Phone Number">
        </div>
        <div class="flex items-center border-teal-500 py-2">
          <input type="text" class="form-input focus:shadow-outline" v-model="clientDetails.email" placeholder="Email">
        </div>
        <button type="submit" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Save</button>
        <button type="reset" @click="clearForm" class="cursor-pointer bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 ml-1 rounded">Clear</button>
      </form>


  </div>
  <div v-else-if="searchClient">
    Search Client Form
  </div>

</template>

<script>
import {ref, reactive} from "vue";

export default {
  name: "Clients",
  setup() {
    const isActive = ref(false)
    const newClient = ref(true)
    const searchClient = ref(false)
    const clientDetails = reactive ({
      firstName: '',
      lastName: '',
      phone1: '',
      phone2: '',
      email: ''
    })


    const newClientButton = () => {
      newClient.value = true
      searchClient.value = false
    }
    const searchClientButton = () => {
      newClient.value = false
      searchClient.value = true
    }

    const handleSave = () => {
      clientDetails.firstName
      clientDetails.lastName
      clientDetails.phone1
      clientDetails.phone2
      clientDetails.email

      console.log(clientDetails.firstName)
      console.log(clientDetails.lastName)
      console.log(clientDetails.phone1)
      console.log(clientDetails.phone2)
      console.log(clientDetails.email)
    }

    const clearForm = () => {
      clientDetails.firstName = ""
      clientDetails.lastName = ""
      clientDetails.phone1 = ""
      clientDetails.phone2 = ""
      clientDetails.email = ""
    }

    return {
      isActive,
      newClient,
      searchClient,
      clientDetails,
      handleSave,
      clearForm,
      newClientButton,
      searchClientButton



    }
  }
}
</script>

<style scoped>

</style>