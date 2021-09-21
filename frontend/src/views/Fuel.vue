<template>
  <div class="about">
    <guage />
    <div>
      <h1>Total Fuel Used</h1>
      <h2>Since xx yy dd hh mm ss</h2>
    </div>
  </div>
</template>


<script>
import guage from '../components/guage.vue'
export default ({
  name:"Fuel",
  components: {
   guage,
  }
})
</script>

<style scoped>

</style>
