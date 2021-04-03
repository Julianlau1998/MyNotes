import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import "./registerServiceWorker";
import VuejsDialog from 'vuejs-dialog';
import VueAnalytics from 'vue-analytics';
import Vue2TouchEvents from 'vue2-touch-events'

Vue.use(Vue2TouchEvents)

// Configuration VueAnalytics
Vue.use(VueAnalytics, {
  id: 'UA-xxxxxxxxx-x'
});
 
import 'vuejs-dialog/dist/vuejs-dialog.min.css';
 
Vue.use(VuejsDialog, {
  html: true,
  okText: 'Continue',
  cancelText: 'Cancel',
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
