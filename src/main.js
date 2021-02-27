import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import "./registerServiceWorker";

Vue.config.productionTip = false

import uuid from '@estudioliver/vue-uuid-v4'
Vue.use(uuid)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
