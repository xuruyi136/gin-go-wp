<template>
  <div id="app">
    <!--   <img src="./assets/logo.png">
    <Register />
    <Home /> -->
    <Navbar />
    <b-container>
      <router-view />
    </b-container>
    <!-- <router-link to="/register" tag="button">register</router-link> -->
  </div>
</template>

<script>
// import Home from "./components/HelloWorld.vue";
// import Register from "./components/Register.vue";
import Navbar from "./components/layout/Navbar.vue";
export default {
  name: "App",
  //组件
  components: { Navbar },
};
</script>

<style>
</style>
