import Vue from 'vue'
import App from './App.vue'
import store from './store'
import vuetify from './plugins/vuetify'
import router from './router'
import firebase from 'firebase/app'
import 'firebase/auth'

let app= '';
const config = {
  apiKey: "AIzaSyAeTTXHUehypyHl1Lvokc90hWXsjsRIX2E",
  authDomain: "alumni-dashboard-1586121981449.firebaseapp.com",
  databaseURL: "https://alumni-dashboard-1586121981449.firebaseio.com",
  projectId: "alumni-dashboard-1586121981449",
  storageBucket: "alumni-dashboard-1586121981449.appspot.com",
  messagingSenderId: "740306678659",
  appId: "1:740306678659:web:745126f8387551a6eafcb6",
  measurementId: "G-RLRB7EXCKT"
}

firebase.initializeApp(config)

Vue.config.productionTip = false

firebase.auth().onAuthStateChanged((user) => {
  if (!app) {
    app = new Vue({
      store,
      vuetify,
      router,
      render: h => h(App)
    }).$mount('#app')
  }
})
