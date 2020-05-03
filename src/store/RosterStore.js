const Vuex = require('vuex');
const axios = require('axios').default.create();
const baseApiUrl = process.env.VUE_APP_BASE_API_URL;

const rosterStore = new Vuex.Store({
    state: {
        roster: [],
    },
    mutations: {
        addAlumni() { },
        updateAlumni() { },
        deleteAlumni() { },
        saveRoster(state, newRoster) {
            state.roster = newRoster;
        }
    },
    actions: {
        getAlumni() {
            return axios.get(baseApiUrl + '/getAlumni');
        }
    },
    modules: {},
});