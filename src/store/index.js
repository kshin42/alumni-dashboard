import Vue from 'vue'
import Vuex from 'vuex'
import AuthService from '../services/AuthService'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    token: null,
  },
  getters: {
    loggedIn(state) {
      return state.token !== null
    },
    getToken(state) {
      return state.token
    }
  },
  mutations: {
    setToken(state, token) {
      state.token = token
    },
    initialiseToken(state) {
      state.token = localStorage.getItem('access_token') || null
    },
    destroyToken(state) {
      state.token = null
    }
  },
  actions: {
    async retrieveToken(context, credentials) {
      const token = await AuthService.retrieveToken(credentials)
      console.log(token)
      localStorage.setItem('access_token', token)
      context.commit('setToken', token)
    },
    destroyToken(context) {
      if (context.getters.loggedIn) {
        AuthService.destroyToken()
        localStorage.removeItem('access_token')
        context.commit('destroyToken')
      }
    }
  },
  modules: {
  }
})
