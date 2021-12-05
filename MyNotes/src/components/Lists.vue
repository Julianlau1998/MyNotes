<template>
    <div id="app">
        <ul 
            id="listParent"
        >
            <!-- <draggable
                :delay="200"
                :delay-on-touch-only="true"
                v-model="storedLists"
                @start="dragging()"
                @end="save()"
            > -->
                <li
                    v-for="(folder, folderIndex) in folders"
                    v-bind:key="folder.id + folderIndex"
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
                <li v-for="(list, listsIndex) in storedLists" v-bind:key="list.id + listsIndex">
                    <button
                        class="noteDiv"
                        @click="openList(list.id)"
                    >
                        <h5><b>{{list.title.substring(0,11)}}</b></h5>
                    </button>
                    <hr id="redLine">
                </li>
            <!-- </draggable> -->
        </ul>
        <!-- <h3 v-else id="loading">
            Loading...
        </h3> -->
        <div 
            class="plusButton"
            ref="plusButton"
            @click="newButton()"
        >
            +
        </div>
        <div 
            id="newFolder"
            v-if="folderChoice"
        >   
            <a
                @click="newFolder()"
            >
                New Folder
            </a>
        </div>
        <div
            id="newNote"
            v-if="folderChoice"
        >
            <a
                @click="newList()"
            >
                New List
            </a>
        </div>
        <new-folder 
            v-if="newFolderDiv"
            @close="close()"
            v-on:send="newFolderDiv = !newFolderDiv"
            type="Lists" 
        />
    </div>
</template>

<script>
import router from '../router'
// import draggable from 'vuedraggable'
import Vue from 'vue'
import Vue2TouchEvents from 'vue2-touch-events'
import { mapState } from 'vuex' 
import NewFolder from './NewFolder.vue'

Vue.use(Vue2TouchEvents)

export default {
    name: 'Lists',
    // components: {draggable},
    components: {NewFolder},
    data () {
        return {
            titles: [],
            lists: [],
            sorting: false,
            folderChoice: false,
            newFolderDiv: false,
        }
    },
    created () {
        this.$store.dispatch('listsModule/getAll')
        this.$store.dispatch('foldersModule/getAll')

        localStorage.setItem('currentComponent', 'Notes')
    },
    computed: {
    ...mapState(['listsModule']),
    ...mapState(['foldersModule']),
        storedLists () {
            if (!this.listsModule.lists.loading && this.listsModule.lists.data !== null) {
                return (!this.listsModule.lists.loading && this.listsModule.lists.data.filter(el => el.folder_id === '' || el.folder_id === '00000000-0000-0000-0000-000000000000')) || []
            } else {
                return (!this.listsModule.lists.loading && this.listsModule.lists.data) || []
            }
        },
        folders () {
            if (!this.foldersModule.folders.loading && this.foldersModule.folders.data !== null) {
                return (!this.foldersModule.folders.loading && this.foldersModule.folders.data.filter(el => el.type ==='Lists')) || []
            } else {
                return (!this.foldersModule.folders.loading && this.foldersModule.folders.data) || []
            }
        }
    },
    methods: {
        openList (id) {
            this.$store.state.transitionName = 'swipe-left'
            this.$store.state.id = id
            router.push(`/list/${id}`)
        },
        openFolder (id) {
            this.$store.state.transitionName = 'fade'
            this.$store.state.id = id
            localStorage.setItem('folder', 'Lists')
            this.$router.push(`/Folder/${id}`)
        },
        swipeHandler () {
            this.$store.state.transitionName = 'fade'
            this.$router.push('/')
        },
        routeToNotes () {
            this.$store.state.transitionName = 'fade'
            this.$router.push('/')
        },
        newList () {
            this.$store.state.transitionName = 'fade'
            this.$router.push('/newList')
        },
        touchHoldHandler () {
            this.sorting = true
        },
         save () {
            this.$store.state.dragging = false
            localStorage.setItem('lists', JSON.stringify(this.storedLists))
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
    },
    mounted () {
        if(this.storedLists === null) {
            this.storedLists = []
        } else {
            for (let i=0; i<this.storedLists.length; i++) {
            this.titles.push(this.storedLists[i].title)
            this.lists.push(this.storedLists[i].list)
        }
        }
        localStorage.setItem('currentComponent', 'Lists')
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
.noteDiv {
    overflow-wrap: break-word; 
    width: 9rem;
    height: 8rem;
    border-radius: 1rem;
    background-color: transparent;
    border: none;
    box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                1px 1px 5px 2px black;
    margin: 1rem 2rem 2.5rem;
    margin-left: 0rem;
    color: lightgray;
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
.whiteLine{
    border-top: 3px solid rgb(134, 134, 134);
    width: 12.5rem;
    margin-top: -0.5rem;
    margin-bottom: 1rem;
}
#whiteLine{
    border-top: 3px solid rgb(134, 134, 134);
    margin-top: -0.5rem;
    margin-bottom: 1rem;
    margin-left: -2.5rem;
}
#notesButton {
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
 }

@media (min-width: 1000px) { 
    .noteDiv {
        width: 11rem;
        height: 9rem;
    }
 }
</style>