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
                    v-for="(note, idx) in storedNotes"
                    v-bind:key="idx"
                >
                    <span v-touch:touchhold="touchHoldHandler">
                        <button class="noteDiv" @click="openNote(note.id)">
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
            @click="newNote()"
            class="plusButton"
            ref="plusButton"
        >
            +
        </div>
    </div>
</template>

<script>
import router from '../router'
import Vue from 'vue'
// import draggable from 'vuedraggable'
import Vue2TouchEvents from 'vue2-touch-events'
import { mapState } from 'vuex' 

Vue.use(Vue2TouchEvents)

export default {
    name: 'Notes',
    // components: {draggable},
    data () {
        return {
            titles: [],
            notes: [],
            sorting: false
        }
    },
    created () {
        const payload = {'userID': this.$store.state.userID}
        this.$store.dispatch('notesModule/getAll', payload)

        localStorage.setItem('currentComponent', 'Notes')
    },

    computed: {
    ...mapState(['notesModule']),
        storedNotes () {
            return (!this.notesModule.notes.loading && this.notesModule.notes.data) || []
        }
    },

    methods: {
        openNote (id) {
            this.$store.state.transitionName = 'swipe-left'
            this.$refs.plusButton.style.opacity = 0
            this.$store.state.id = id
            router.push(`/Note/${id}`)
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
.noteDiv {
    overflow-wrap: break-word; 
    width: 9rem;
    height: 8rem;
    border-radius: 1rem;
    background-color: transparent;
    border: none;
    box-shadow: -1px -1px 3px 0px rgb(133, 133, 133),
                1px 1px 4px 2px black;
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
        width: 7rem;
        height: 7rem;
    }
 }
@media (max-width: 327px) { 
    .noteDiv {
        width: 6rem;
        height: 6rem;
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