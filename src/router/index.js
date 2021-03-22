import Vue from 'vue'
import VueRouter from 'vue-router'
import Notes from '../views/Notes.vue'
import NewNote from '../views/NewNote.vue'
import Note from '../views/Note.vue'
import Lists from '../views/Lists.vue'
import NewList from '../views/NewList.vue'
import SingleList from '../views/List.vue'


Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'notes',
    component: Notes
  },
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
    path: '/Lists',
    name: 'Lists',
    component: Lists
  }
  ,
  {
    path: '/List/:id',
    name: 'SingleList',
    component: SingleList
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
