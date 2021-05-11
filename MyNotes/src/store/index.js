import Vue from 'vue'
import Vuex from 'vuex'
import notesModule from './notesModule'
import listsModule from './listsModule'
import foldersModule from './foldersModule'


Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    id: 0,
    userID: '',
    transitionName: 'swipe-left',
    componentTransitionName: 'component-swipe-left',
    currentComponent: localStorage.getItem('currentComponent'),
    currentFolder: '',
    dragging: false,
    // localhost: 'https://mynotes-312216.ey.r.appspot.com/api/',
    localhost: '/api/',
    // localhost: 'https://mynotes-backend-f7xxtlrmbq-ey.a.run.app/api/',
    // localhost: 'http://139.162.158.148:1323/api/',
    // localhost: 'http://192.168.178.58:1323/api/',
    folder: ''
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    notesModule,
    listsModule,
    foldersModule
  }
})