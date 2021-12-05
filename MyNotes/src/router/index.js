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
import { authGuard } from "../auth/authGuard";


Vue.use(VueRouter)

const routes = [
  {
    path: '/newNote',
    name: 'newNote',
    component: NewNote,
    beforeEnter: authGuard 
  },
  {
    path: '/newList',
    name: 'newList',
    component: NewList,
    beforeEnter: authGuard
  },
  { 
    path: '/Note/:id',
    name: 'Note',
    component: Note,
    beforeEnter: authGuard
  },
  {
    path: '/List/:id',
    name: 'SingleList',
    component: SingleList,
    beforeEnter: authGuard
  }
  ,
  {
    path: '/',
    name: 'Home',
    component: Home,
    beforeEnter: authGuard
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
    component: Settings,
    beforeEnter: authGuard
  },
  {
    path: '/folder/:id',
    name: 'Folder',
    component: Folder,
    beforeEnter: authGuard
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
