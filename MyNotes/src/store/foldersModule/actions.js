import axios from 'axios'

const token = sessionStorage.getItem('token')
axios.defaults.headers.common['authorization'] = token

export function getAll ({ commit, state }) {
  commit('GET_FOLDERS')
  axios
    .get(`${state.localhost}folders`)
    .then(response => {
      commit('RECEIVE_FOLDERS', response.data)
    })
    .catch(err => {
      console.log(err)
    })
}
export function getOne ({ commit, state }, payload) {
  commit('GET_FOLDER')
  axios
    .get(`${state.localhost}folder/${payload.id}`, {
        headers: {
          'userId': payload.userID
        }
    })
    .then(response => {
      commit('RECEIVE_FOLDER', response.data)
    })
    .catch(err => {
      console.log(err)
    })
}
export function post ({ commit, state }, payload) {
  commit('POST_FOLDER')
  axios
    .post(`${state.localhost}folders`, payload.folder, {
      headers: {
        'userId': payload.userID
      }
    })
    .then(function () {
      commit('FOLDER_POSTED', payload.folder)
    })
    .catch(function (error) {
      console.log(error)
    })
}
export function deleteOne ({ commit, state }, payload) {
  commit('DELETE_FOLDER')
  axios
    .delete(`${state.localhost}folder/${payload.id}`, {
      headers: {
        'userId': payload.userID
      }
    })
    .then(function () {
      commit('FOLDER_DELETED')
    })
    .catch(function (error) {
      console.log(error)
    })
}
export function put ({ commit, state }, payload) {
  commit('PUT_FOLDER')
  axios
    .put(`${state.localhost}folder/${payload.folder.id}`, payload.folder, {
      headers: {
        'userId': payload.userID
      }
    })
    .then(function () {
      commit('FOLDER_PUT')
    })
    .catch(function (error) {
      console.log(error)
    })
}