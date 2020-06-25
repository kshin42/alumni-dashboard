import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import CreateOrg from '../views/CreateOrg.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/createOrg',
    name: 'CreateOrg',
    component: () => import('../views/CreateOrg.vue')
  },
  {
    path: '/roster',
    name: 'Roster',
    component: () => import('../views/Roster.vue')
  },
  {
    path: '/userImport',
    name: 'UserImport',
    component: () => import('../views/UserImport.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
