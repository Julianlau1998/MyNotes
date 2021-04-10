import Vue from 'vue'
import Vuex from 'vuex'
import notesModule from './notesModule'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    id: 0,
    userID: '3',
    transitionName: 'swipe-left',
    componentTransitionName: 'component-swipe-left',
    currentComponent: localStorage.getItem('currentComponent')
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    notesModule
  }
})
