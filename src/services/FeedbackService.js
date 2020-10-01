const axios = require('axios').default.create();
import store from '../store'

export default {
    
    submitFeedback(data) {
        let config = this.setUpSessionHeader()
        return axios
          .post(process.env.VUE_APP_BASE_API_URL + "/submitFeedback", data, config)
          .then((response) => response)
          .catch((err) => {
            if (err.response && err.response.status === 401) {
              store.dispatch('destroyToken')
            } else {
                console.log(err)
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