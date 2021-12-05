export function GET_FOLDERS (state) {
  state.folders.loading = true
  state.err = null
}

export function RECEIVE_FOLDERS  (state, folders) {
  state.folders.data = folders
  state.folders.loading = false
}

export function GET_FOLDER (state) {
  state.folder.loading = true
  state.err = null
}

export function RECEIVE_FOLDER (state, folder) {
  state.folder.data = folder
  state.folder.loading = false
}
export function POST_FOLDER (state) {
  state.folder.submitting = true
  state.err = null
}
export function FOLDER_POSTED (state, folder) {
  state.folders.data.unshift(folder)
  state.folder.submitting = false
}
export function DELETE_FOLDER (state) {
  state.folder.loading = true
  state.err = null
}
export function FOLDER_DELETED (state) {
  state.folder.loading = false
}
export function PUT_FOLDER (state) {
  state.folder.submitting = true
  state.err = null
}
export function FOLDER_PUT (state) {
  state.folder.submitting = false
}