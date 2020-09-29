const axios = require('axios').default.create();

export default {
    
    submitFeedback(data) {
        return axios
          .post(process.env.VUE_APP_BASE_API_URL + "/submitFeedback", data)
          .then((response) => response);
    }
}