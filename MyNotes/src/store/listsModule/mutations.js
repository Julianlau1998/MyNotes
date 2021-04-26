export function GET_LISTS (state) {
  state.lists.loading = true
  state.err = null
}

export function RECEIVE_LISTS  (state, lists) {
  state.lists.data = lists
  state.lists.loading = false
}

export function GET_FOLDER_LISTS (state) {
  state.lists.loading = true
  state.err = null
}

export function RECEIVE_FOLDER_LISTS  (state, lists) {
  state.lists.data = lists
  state.lists.loading = false
}

export function GET_LIST (state) {
  state.list.loading = true
  state.err = null
}

export function RECEIVE_LIST (state, list) {
  state.list.data = list
  state.list.loading = false
}
export function POST_LIST (state) {
  state.list.submitting = true
  state.err = null
}
export function LIST_POSTED (state) {
  state.list.submitting = false
}
export function DELETE_LIST (state) {
  state.list.loading = true
  state.err = null
}
export function LIST_DELETED (state) {
  state.list.loading = false
}
export function PUT_LIST (state) {
  state.list.submitting = true
  state.err = null
}
export function LIST_PUT (state) {
  state.list.submitting = false
}