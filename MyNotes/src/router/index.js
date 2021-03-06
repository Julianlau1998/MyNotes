import Vue from 'vue'
import VueRouter from 'vue-router'
import NewNote from '../views/NewNote.vue'
import Note from '../views/Note.vue'
import NewList from '../views/NewList.vue'
import SingleList from '../views/List.vue'
import Home from '../views/Home.vue'
import Settings from '../views/Settings.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Folder from '../views/FolderView.vue'


Vue.use(VueRouter)

const routes = [
  {
    path: '/newNote',
    name: 'newNote',
    component: NewNote
  },
  {
    path: '/newList',
    name: 'newList',
    component: NewList
  },
  { 
    path: '/Note/:id',
    name: 'Note',
    component: Note
  },
  {
    path: '/List/:id',
    name: 'SingleList',
    component: SingleList
  }
  ,
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/Login',
    name: 'Login',
    component: Login
  },
  {
    path: '/Register',
    name: 'Register',
    component: Register
  },
  {
    path: '/settings',
    name: 'Settings',
    component: Settings
  },
  {
    path: '/folder/:id',
    name: 'Folder',
    component: Folder
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
