import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import "./registerServiceWorker";
import VuejsDialog from 'vuejs-dialog';
 
import 'vuejs-dialog/dist/vuejs-dialog.min.css';
 
Vue.use(VuejsDialog, {
  html: true,
  okText: 'Continue',
  cancelText: 'Stop',
  animation: 'bounce'
});

Vue.config.productionTip = false

import uuid from '@estudioliver/vue-uuid-v4'
Vue.use(uuid)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
