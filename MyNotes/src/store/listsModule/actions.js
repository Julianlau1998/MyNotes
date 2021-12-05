import axios from 'axios'

const token = sessionStorage.getItem('token')
axios.defaults.headers.common['authorization'] = token

export function getAll ({ commit, state }) {
  commit('GET_LISTS')
  axios
    .get(`${state.localhost}lists`)
    .then(response => {
      commit('RECEIVE_LISTS', response.data)
    })
    .catch(err => {
      console.log(err)
    })
}

export function getByCategory ({ commit, state }, payload) {
  commit('GET_FOLDER_LISTS')
  axios
    .get(`${state.localhost}lists/folderlists`, {
      headers: {
        'userId': payload.userID,
        'folderID': payload.folderID
      }
    })
    .then(response => {
      commit('RECEIVE_FOLDER_LISTS', response.data)
    })
    .catch(err => {
      console.log(err)
    })
}

export function getOne ({ commit, state }, payload) {
  commit('GET_LIST')
  axios
    .get(`${state.localhost}list/${payload.id}`, {
        headers: {
          'userId': payload.userID
        }
    })
    .then(response => {
      commit('RECEIVE_LIST', response.data)
    })
    .catch(err => {
      console.log(err)
    })
}
export function post ({ commit, state }, payload) {
  commit('POST_LIST')
  axios
    .post(`${state.localhost}lists`, payload.list, {
      headers: {
        'userId': payload.userID
      }
    })
    .then(function () {
      commit('LIST_POSTED')
    })
    .catch(function (error) {
      console.log(error)
    })
}
export function deleteOne ({ commit, state }, payload) {
  commit('DELETE_LIST')
  axios
    .delete(`${state.localhost}list/${payload.id}`, {
      headers: {
        'userId': payload.userID
      }
    })
    .then(function () {
      commit('LIST_DELETED')
    })
    .catch(function (error) {
      console.log(error)
    })
}
export function put ({ commit, state }, payload) {
  commit('PUT_LIST')
  axios
    .put(`${state.localhost}list/${payload.list.id}`, payload.list, {
      headers: {
        'userId': payload.userID
      }
    })
    .then(function () {
      commit('LIST_PUT')
    })
    .catch(function (error) {
      console.log(error)
    })
}