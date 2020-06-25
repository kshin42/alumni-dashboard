import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/createProfile',
    name: 'CreateProfile',
    component: () => import('../views/CreateProfile.vue')
  },
  {
    path: '/roster',
    name: 'Roster',
    component: () => import('../views/Roster.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
