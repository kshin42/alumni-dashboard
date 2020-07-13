import Vue from 'vue'
import Vuex from 'vuex'
import router from '../router'
import AuthService from '../services/AuthService'
import RosterService from '../services/RosterService'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    token: null,
    currentUser: null
  },
  getters: {
    loggedIn(state) {
      return state.token !== null
    },
    getToken(state) {
      return state.token
    },
    getCurrentUser(state) {
      return state.currentUser
    }
  },
  mutations: {
    setToken(state, token) {
      state.token = token
    },
    initialiseToken(state) {
      state.token = localStorage.getItem('access_token') || null
      state.currentUser = localStorage.getItem('current_user') || null
    },
    destroyToken(state) {
      state.token = null
      state.currentUser = null
    },
    setCurrentUser(state, email) {
      state.currentUser = email
    }
  },
  actions: {
    async retrieveToken(context, credentials) {
      const token = await AuthService.retrieveToken(credentials)
      localStorage.setItem('access_token', token)
      localStorage.setItem('current_user', credentials.email)
      context.commit('setToken', token)
      context.commit('setCurrentUser', credentials.email)
    },
    destroyToken(context) {
      if (context.getters.loggedIn) {
        AuthService.destroyToken()
        localStorage.removeItem('access_token')
        localStorage.removeItem('current_user')
        context.commit('destroyToken')
        router.push('/login')
      }
    },
    async getRoster(context) {
      const response = await RosterService.getRoster(context.state.token)
      if (response.status == 401) {
        console.log("caught 40111")
        await context.dispatch('destroyToken')
      }
      return response.data
    }
  },
  modules: {
  }
})
