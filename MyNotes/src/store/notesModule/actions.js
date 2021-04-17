import axios from 'axios'

export function getAll ({ commit, state }, payload) {
  commit('GET_NOTES')
  axios
    .get(`${state.localhost}notes`, {
      headers: {
        'userId': payload.userID
      }
    })
    .then(response => {
      commit('RECEIVE_NOTES', response.data)
    })
    .catch(err => {
      console.log(err)
    })
}
export function getOne ({ commit, state }, payload) {
  commit('GET_NOTE')
  axios
    .get(`${state.localhost}note/${payload.id}`, {
        headers: {
          'userId': payload.userID
        }
    })
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
    .post(`${state.localhost}notes`, payload.note, {
      headers: {
        'userId': payload.userID
      }
    })
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
    .delete(`${state.localhost}note/${payload.id}`, {
      headers: {
        'userId': payload.userID
      }
    })
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
    .put(`${state.localhost}note/${payload.note.id}`, payload.note, {
      headers: {
        'userId': payload.userID
      }
    })
    .then(function () {
      commit('NOTE_PUT')
    })
    .catch(function (error) {
      console.log(error)
    })
}