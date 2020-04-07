const axios = require('axios').default.create();


export default {
    
    getRoster() {
        return axios
          .get(process.env.VUE_APP_BASE_API_URL + "/getAlumni")
          .then((response) => response.data);
    }
}