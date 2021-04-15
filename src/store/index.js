import Vue from 'vue'
import Vuex from 'vuex'
import notesModule from './notesModule'
import listsModule from './listsModule'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    id: 0,
    userID: '',
    transitionName: 'swipe-left',
    componentTransitionName: 'component-swipe-left',
    currentComponent: localStorage.getItem('currentComponent'),
    localhost: 'http://139.162.158.148:1324/api/users'
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    notesModule,
    listsModule
  }
})