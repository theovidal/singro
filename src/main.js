import Vue from 'vue'
import App from './renderer/App.vue'

import router from './renderer/plugins/router'
import store from './renderer/store/index'
import vuetify from './renderer/plugins/vuetify'
import './renderer/plugins/registerServiceWorker'

import '@babel/polyfill'
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@mdi/font/css/materialdesignicons.css'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
