<template>
  <div id="app">
      <Navigation />
      <router-view></router-view>
  </div>
</template>


<script>
import Navigation from "./components/Navigation.vue"
export default {
 data: () => ({}),
 components: {
   Navigation
 }
}
</script>
