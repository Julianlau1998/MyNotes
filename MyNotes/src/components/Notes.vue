<template>
    <div id="app">
        <ul
            id="listParent"
        >
            <!-- <draggable
                :delay="200"
                :delay-on-touch-only="true"
                v-model="storedNotes"
                @start="dragging()"
                @end="save()"
            > -->
                <li
                    v-for="(folder, index) in folders"
                    v-bind:key="folder.id + index"
                >
                    <span v-touch:touchhold="touchHoldHandler">
                        <button
                            class="noteDiv" 
                            v-bind:style="{backgroundColor: folder.color}"
                            @click="openFolder(folder.id)"
                        >
                            <h5><b>{{folder.title.substring(0,11)}}</b></h5>
                        </button>
                    </span>
                    <hr id="redLine">
                </li>
                <li
                    v-for="(note, idx) in storedNotes"
                    v-bind:key="note.id + idx"
                >
                    <span
                        v-touch:touchhold="touchHoldHandler"
                        v-if="note.folder_id === '' || note.folder_id == '00000000-0000-0000-0000-000000000000'"
                    >
                        <button
                            class="noteDiv"
                            @click="openNote(note.id)"
                        >
                            <h5><b>{{note.title.substring(0,11)}}</b></h5>
                        </button>
                    </span>
                    <hr id="redLine">
                </li>
            <!-- </draggable> -->
        </ul>
        <!-- <h3 v-else id="loading">
            Loading...
        </h3> -->
        <div 
            @click="newButton()"
            class="plusButton"
            ref="plusButton"
        >
            +
        </div>
        <div 
            class="newFolder"
            v-if="folderChoice"
        >   
            <a
                id="newFolder"
                @click="newFolder()"
            >
                New Folder
            </a>
        </div>
        <div
            class="newNote"
            v-if="folderChoice"
        >
            <a
                id="newNote"
                @click="newNote()"
            >
                New Note
            </a>
        </div>
        <transition
            name="new-folder"
        >
            <new-folder 
            v-if="newFolderDiv"
            v-on:send="newFolderDiv = !newFolderDiv"
            @close="close()"
            type="Notes"
        />
        </transition>
    </div>
</template>

<script>
import router from '../router'
import Vue from 'vue'
// import draggable from 'vuedraggable'
import Vue2TouchEvents from 'vue2-touch-events'
import { mapState } from 'vuex' 
import NewFolder from './NewFolder.vue'

Vue.use(Vue2TouchEvents)

export default {
    name: 'Notes',
    // components: {draggable},
    components: {NewFolder},
    data () {
        return {
            titles: [],
            notes: [],
            sorting: false,
            folderChoice: false,
            newFolderDiv: false,
        }
    },
    created () {
        this.$auth.getTokenSilently()
            .then((token) => {
                console.log(token)
            })
        this.$store.dispatch('notesModule/getAll')
        this.$store.dispatch('foldersModule/getAll')

        localStorage.setItem('currentComponent', 'Notes')
    },

    computed: {
    ...mapState(['notesModule']),
    ...mapState(['foldersModule']),
        storedNotes () {
            if (!this.notesModule.notes.loading && this.notesModule.notes.data !== null) {
                console.log(this.notesModule.notes.data)
                return (!this.notesModule.notes.loading && this.notesModule.notes.data.filter(el => el.folder_id === '' || el.folder_id === '00000000-0000-0000-0000-000000000000')) || []
            } else {
                return (!this.notesModule.notes.loading && this.notesModule.notes.data) || []
            }
        },
        folders () {
            if (!this.foldersModule.folders.loading && this.foldersModule.folders.data !== null) {
                return (!this.foldersModule.folders.loading && this.foldersModule.folders.data.filter(el => el.type === 'Notes')) || []
            } else {
                return (!this.foldersModule.folders.loading && this.foldersModule.folders.data) || []
            }
        },
    },
    methods: {
        openNote (id) {
            this.$store.state.transitionName = 'swipe-left'
            this.$refs.plusButton.style.opacity = 0
            this.$store.state.id = id
            router.push(`/Note/${id}`)
        },
        openFolder (id) {
            this.$store.state.transitionName = 'fade'
            this.$store.state.id = id
            localStorage.setItem('folder', 'Notes')
            this.$router.push(`/Folder/${id}`)
        },
        swipeHandler () {
            this.$store.state.transitionName = 'fade'
            this.$router.push('/lists')
        },
        routeToLists () {
            this.$store.state.transitionName = 'fade'
            this.$router.push('/lists')
        },
        newNote () {
            this.$store.state.transitionName = 'fade'
            this.$router.push('/newNote')
        },
        touchHoldHandler () {
            this.sorting = true
        },
        save () {
            this.$store.state.dragging = false
            localStorage.setItem('notes', JSON.stringify(this.storedNotes))
        },
        dragging () {
            this.$store.state.dragging = true
        },
        newButton () {
            if (this.$refs.plusButton.innerHTML === ' + ') {
                this.$refs.plusButton.innerHTML = 'x'
            } else {
                this.$refs.plusButton.innerHTML = ' + '
            }
            this.folderChoice = !this.folderChoice
        },
        newFolder () {
            this.newFolderDiv=true 
            this.folderChoice=false
            this.newFolderDiv = true
        },
        close () {
            this.folderChoice = !this.folderChoice
            this.newFolderDiv=false
        },
        closeFolderChoiceDiv () {
            this.folderChoice=!this.folderChoice
        },
        closeNewFolderDiv () {
            this.newFolderDiv=!this.newFolderDiv
        }
    }
}
</script>

<style scoped>
body {
    margin: 0;
}
ul {
    list-style-type: none;
}
ul li {
    display: inline-block;
}

.noteDiv:active {
    box-shadow: -1px -1px 3px 0px black,
                1px 1px 3px 0px rgb(133, 133, 133);
}
.noteDiv:focus {
    outline: none;
}
.header {
    font-size: 3.5rem;
    color: lightgray;
}
#listParent{
    margin-left: -0.5rem;
}
#listsButton {
    background: transparent;
    border: none;
    outline: none;
}
#listParent {
    min-height: 79vh;
}
#loading {
    margin-top: 3rem;
}
#folder {
    width: 2rem;
    margin-top: -1rem;
}
#newFolderBackground  {
    position: absolute;
    top: -11.3rem;
    left: 0;
    width:100vw;
    height: 100vh;
    background-color: black;
    z-index: 0;
    opacity: 0.5;
}
.plusButton {
    z-index: 10;
}

@media (min-width: 1000px) { 
    .noteDiv {
        width: 7rem;
        height: 7rem;
    }
    #listParent {
        width: 80vw;
        margin-left: 10vw;
    }
 }
 @media (min-width: 1500px) { 
    #listParent {
        width: 70vw;
        margin-left: 15vw;
    }
 }
@media (max-width: 390px) { 
    .noteDiv {
        width: 8rem;
        height: 7rem;
    }
 }
@media (max-width: 327px) { 
    .noteDiv {
        width: 8rem;
        height: 7rem;
    }
    body {
        margin: 100rem;
    }
 }

@media (min-width: 1000px) { 
    .noteDiv {
        width: 11rem;
        height: 9rem;
    }
 }
</style>