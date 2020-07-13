const axios = require('axios').default.create();
import store from '../store'

export default {
    
    getRoster(session) {
        let config = this.setUpSessionHeader(session)
        return axios
          .get(process.env.VUE_APP_BASE_API_URL + "/getAlumni", config)
          .then((response) => response)
          .catch((err) => {
              if (err.response && err.response.status === 401) {
                store.dispatch('destroyToken')
              }
          })
    },

    setUpSessionHeader(session) {
        let config = {
            headers: {
                'Authorization': store.getters.getToken,
                'X-user-email': store.getters.getCurrentUser
            }
        }
        return config
    }

}