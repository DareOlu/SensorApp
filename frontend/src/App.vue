<template>
  <div id="app">
     <!-- <Navbar />  -->
     <navbar/>
   
        
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue'
//import Navbar from './components/Navbar.vue'

export default {
  name:'App',
  components:{
    Navbar,  
  },
}
</script>
<style scoped>
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
      -webkit-font-smoothing: antialiased;
      -moz-osx-font-smoothing: grayscale;
  }
  * {
    color: #ffffff;
    background-color: #000911;
    width: 100%;
    margin: 0px;
    padding: 0px;  
  }

  .body{
    background-color: #000911;
  }
</style>
