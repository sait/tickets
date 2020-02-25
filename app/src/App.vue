<template>
  <div id="app">
    <nav-bar></nav-bar>
    <br>
    <router-view/>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue"
export default {
  components: {
    NavBar
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

#multifilter_estado{
  visibility: hidden;
}

#format-pre {
  font-size: 1rem;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
}

#alert-text {
  color: red;
}
</style>
