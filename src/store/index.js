import Vue from 'vue'
import Vuex from 'vuex'
import router from '../router'
import AuthService from '../services/AuthService'
import RosterService from '../services/RosterService'
import ResumeService from '../services/ResumeService'
import FeedbackService from '../services/FeedbackService'

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
      return response.data
    },
    async uploadResume(context, payload) {
      const response = await ResumeService.uploadResume(payload)
      return response
    },
    async getResume(context) {
      const response = await ResumeService.getResume()
      return response
    },
    async getResumes(context) {
      const response = await ResumeService.getResumes()
      return response
    },
    async submitFeedback(context, payload) {
      const response = await FeedbackService.submitFeedback(payload)
      return response
    }
  },
  modules: {
  }
})
