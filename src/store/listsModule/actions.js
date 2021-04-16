import axios from 'axios'

export function getAll ({ commit }, payload) {
  commit('GET_LISTS')
  axios
    .get(`${this.$store.state.localhost}lists`, {
      headers: {
        'userId': payload.userID
      }
    })
    .then(response => {
      commit('RECEIVE_LISTS', response.data)
    })
    .catch(err => {
      console.log(err)
    })
}
export function getOne ({ commit }, payload) {
  commit('GET_LIST')
  axios
    .get(`${this.$store.state.localhost}list/${payload.id}`, {
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
export function post ({ commit }, payload) {
  commit('POST_LIST')
  axios
    .post(`${this.$store.state.localhost}lists`, payload.list, {
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
export function deleteOne ({ commit }, payload) {
  commit('DELETE_LIST')
  axios
    .delete(`${this.$store.state.localhost}list/${payload.id}`, {
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
export function put ({ commit }, payload) {
  commit('PUT_LIST')
  axios
    .put(`${this.$store.state.localhost}list/${payload.list.id}`, payload.list, {
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