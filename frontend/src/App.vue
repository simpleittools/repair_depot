<template>
  <Navbar />
  <div class="container">
    <router-view/>
  </div>
</template>

<script>
  import Navbar from "@/components/navbar/Navbar";
  export default {
    components: {
      Navbar
    }
  }
</script>
<style>

</style>
