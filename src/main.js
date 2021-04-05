import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import "./registerServiceWorker";
import VuejsDialog from 'vuejs-dialog';
import Vue2TouchEvents from 'vue2-touch-events'
import VueI18n from 'vue-i18n'
import langs from './lang'

Vue.config.productionTip = false

Vue.use(VueI18n)
Vue.use(Vue2TouchEvents)
 
import 'vuejs-dialog/dist/vuejs-dialog.min.css';

let chosenLanguage = localStorage.getItem('language')
if (chosenLanguage === null) {
  chosenLanguage = 'English'
}

let lang = langs.en
if (chosenLanguage == 'German') {
  lang = langs.de
}
 
Vue.use(VuejsDialog, {
  html: true,
  okText: lang.text.continue,
  cancelText: lang.text.cancel,
  animation: 'bounce'
});

Vue.config.productionTip = false

import uuid from '@estudioliver/vue-uuid-v4'
Vue.use(uuid)

const messages = {
  English: langs.en,
  German: langs.de
}

const i18n = new VueI18n({
  locale: chosenLanguage, // set locale
  messages // set locale messages
})

new Vue({
  i18n,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
