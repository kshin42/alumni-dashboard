const axios = require('axios').default.create();

export default {
    
    createMember(data) {
        return axios
          .post(process.env.VUE_APP_BASE_API_URL + "/createMember", data)
          .then((response) => response);
    }
}