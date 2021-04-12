import axios from 'axios'

export function getAll ({ commit }, payload) {
  commit('GET_NOTES')
  axios
    .get('/api/notes', {
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
export function getOne ({ commit }, payload) {
  commit('GET_NOTE')
  axios
    .get(`/api/note/${payload.id}`, {
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
export function post ({ commit }, payload) {
  commit('POST_NOTE')
  axios
    .post('/api/notes', payload.note, {
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
export function deleteOne ({ commit }, payload) {
  commit('DELETE_NOTE')
  axios
    .delete(`/api/note/${payload.id}`, {
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
export function put ({ commit }, payload) {
  commit('PUT_NOTE')
  axios
    .put(`/api/note/${payload.note.id}`, payload.note, {
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
