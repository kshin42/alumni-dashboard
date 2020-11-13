<template>
  <v-app  id="app">
      <v-content class="bg">
        <transition name='fade' mode="out-in">
          <router-view/>
        </transition>
      </v-content>
  </v-app>
</template>


<script>

export default {
  data: () => ({
      prevHeight: 0,
  })
}
</script>

<style>
.fade-enter-active,
.fade-leave-active {
  transition-duration: 0.3s;
  transition-property: opacity;
  transition-timing-function: ease;
}

.fade-enter,
.fade-leave-active {
  opacity: 0
}

.bg {
  background: #4281A4;
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
</style>