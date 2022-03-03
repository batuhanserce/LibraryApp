<template>
  <div>
    <Navbar></Navbar>
  </div>
</template>

<script>
import Navbar from "@/component/navbar";
export default {
  components: {
    Navbar
  },
  name: "home",
  data() {
    return {
      form: {
        isim: "",
        soyisim: "",
        kod: ""
      },
      isBusy: false,
      disable: false,
      kodDisable: false,
    }
  },

}
</script>

<style scoped>
</style>
