import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import firebase from 'firebase/app'
import 'firebase/auth'

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
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: {
      transitionName: 'slide',
    }
  },
  {
    path: '/roster',
    name: 'Roster',
    component: () => import('../views/Roster.vue'),
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/resume',
    name: 'Resume',
    component: () => import('../views/Resume.vue'),
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/resumeWorkshop',
    name: 'ResumeWorkshop',
    component: () => import('../views/ResumeWorkshop.vue'),
    meta: {
      requiresAuth: true
    }
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  const currentUser = firebase.auth().currentUser;
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);

  if (requiresAuth && !currentUser) {
    next('login');
  } else if (!requiresAuth && currentUser) {
    next('resume');
  } else {
    next();
  }

});

export default router
