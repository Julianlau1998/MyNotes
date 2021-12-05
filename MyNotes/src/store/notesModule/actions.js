import axios from 'axios'

const token = sessionStorage.getItem('token')
axios.defaults.headers.common['authorization'] = token

export function getAll ({ commit, state }) {
  commit('GET_NOTES')
  axios
    .get(`${state.localhost}notes`)
    .then(response => {
      commit('RECEIVE_NOTES', response.data)
    })
    .catch(err => {
      console.log(err)
    })
}
export function getByCategory ({ commit, state }, payload) {
  commit('GET_FOLDER_NOTES')
  axios
    .get(`${state.localhost}notes/foldernotes`, {
      headers: {
        'folderID': payload.folderID
      }
    })
    .then(response => {
      commit('RECEIVE_FOLDER_NOTES', response.data)
    })
    .catch(err => {
      console.log(err)
    })
}
export function getOne ({ commit, state }, payload) {
  commit('GET_NOTE')
  axios
    .get(`${state.localhost}note/${payload.id}`)
    .then(response => {
      commit('RECEIVE_NOTE', response.data)
    })
    .catch(err => {
      console.log(err)
    })
}
export function post ({ commit, state }, payload) {
  commit('POST_NOTE')
  axios
    .post(`${state.localhost}notes`, payload.note)
    .then(function () {
      commit('NOTE_POSTED')
    })
    .catch(function (error) {
      console.log(error)
    })
}
export function deleteOne ({ commit, state }, payload) {
  commit('DELETE_NOTE')
  axios
    .delete(`${state.localhost}note/${payload.id}`)
    .then(function () {
      commit('NOTE_DELETED')
    })
    .catch(function (error) {
      console.log(error)
    })
}
export function put ({ commit, state }, payload) {
  commit('PUT_NOTE')
  axios
    .put(`${state.localhost}note/${payload.note.id}`, payload.note)
    .then(function () {
      commit('NOTE_PUT')
    })
    .catch(function (error) {
      console.log(error)
    })
}