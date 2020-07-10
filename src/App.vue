<template>
  <v-app  id="app">
      <v-content class="bg">
        <transition :name="transitionName" mode="out-in">
          <router-view/>
        </transition>
      </v-content>
  </v-app>
</template>


<script>
const DEFAULT_TRANSITION = 'fade';

export default {
  data: () => ({
      prevHeight: 0,
      transitionName: DEFAULT_TRANSITION,
  }),
  created() {
    this.$router.beforeEach((to, from, next) => {
      // set up fancy transitions
      let transitionName = to.meta.transitionName || from.meta.transitionName;

      if (transitionName === 'slide') {
        const toDepth = to.path.split('/').length;
        const fromDepth = from.path.split('/').length;
        transitionName = toDepth < fromDepth ? 'slide-right' : 'slide-left';
      }

      this.transitionName = transitionName || DEFAULT_TRANSITION;

      // check auth
      if (to.matched.some(record => record.meta.requiresAuth)) {
        if (!this.$store.getters.loggedIn) {
          next({
            path: '/login',
            query: { redirect: to.fullPath}
          })
        } else {
          next()
        }
      } else {
        next()
      }
    });
  }
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

.slide-left-enter-active,
.slide-left-leave-active,
.slide-right-enter-active,
.slide-right-leave-active {
  transition-duration: 0.3s;
  transition-property: height, opacity, transform;
  transition-timing-function: cubic-bezier(0.55, 0, 0.1, 1);
  overflow: hidden;
}

.slide-left-enter,
.slide-right-leave-active {
  opacity: 0;
  transform: translate(2em, 0);
}

.slide-left-leave-active,
.slide-right-enter {
  opacity: 0;
  transform: translate(-2em, 0);
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