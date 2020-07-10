import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: { transitionName: 'slide' }
  },
  {
    path: '/createProfile',
    name: 'CreateProfile',
    component: () => import('../views/CreateProfile.vue'),
    meta: {
      transitionName: 'slide',
      requiresVisitor: true
    }
  },
  {
    path: '/signIn',
    name: 'SignIn',
    component: () => import('../views/SignIn.vue'),
    meta: {
      transitionName: 'slide',
      requiresVisitor: true
    }
  },
  {
    path: '/roster',
    name: 'Roster',
    component: () => import('../views/Roster.vue'),
    meta: {
      transitionName: 'slide',
      requiresAuth: true
    }
  }
]

const router = new VueRouter({
  routes
})

export default router
