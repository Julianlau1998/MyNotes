export function GET_NOTES (state) {
  state.notes.loading = true
  state.err = null
}

export function GET_FOLDER_NOTES (state) {
  state.notes.loading = true
  state.err = null
}

export function RECEIVE_FOLDER_NOTES  (state, notes) {
  state.notes.data = notes
  state.notes.loading = false
}

export function RECEIVE_NOTES  (state, notes) {
  state.notes.data = notes
  state.notes.loading = false
}

export function GET_NOTE (state) {
  state.note.loading = true
  state.err = null
}

export function RECEIVE_NOTE (state, note) {
  state.note.data = note
  state.note.loading = false
}
export function POST_NOTE (state) {
  state.note.submitting = true
  state.err = null
}
export function NOTE_POSTED (state) {
  state.note.submitting = false
}
export function DELETE_NOTE (state) {
  state.note.loading = true
  state.err = null
}
export function NOTE_DELETED (state) {
  state.note.loading = false
}
export function PUT_NOTE (state) {
  state.note.submitting = true
  state.err = null
}
export function NOTE_PUT (state) {
  state.note.submitting = false
}