import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    id: 0,
    transitionName: 'swipe-left',
    componentTransitionName: 'component-swipe-left',
    currentComponent: localStorage.getItem('currentComponent'),
    dragging: false,
  },
  mutations: {
  },
  actions: {
  },
  modules: {
  }
})
