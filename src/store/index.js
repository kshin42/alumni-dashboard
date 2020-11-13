import Vue from 'vue'
import Vuex from 'vuex'
import router from '../router'
import RosterService from '../services/RosterService'
import ResumeService from '../services/ResumeService'
import FeedbackService from '../services/FeedbackService'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    currentUser: null
  },
  getters: {
    getCurrentUser(state) {
      return state.currentUser
    }
  },
  mutations: {
    setCurrentUser(state, email) {
      state.currentUser = email
    }
  },
  actions: {
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
