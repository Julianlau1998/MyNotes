<template>
    <div id="app">
        <ul id="listParent">
            <li v-for="(note, idx) in storedNotes" v-bind:key="idx">
                <button class="noteDiv" @click="openNote(note.id)">
                    <h5><b>{{note.title.substring(0,11)}}</b></h5>
                </button>
                <hr id="redLine">
            </li>
        </ul>
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
import Vue2TouchEvents from 'vue2-touch-events'
import { mapState } from 'vuex' 

Vue.use(Vue2TouchEvents)

export default {
    name: 'Notes',
    data () {
        return {
            // storedNotes: JSON.parse(localStorage.getItem('notes')),
            titles: [],
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
            this.$router.push('/newNote')
        }
    },
    created () {
        const payload = {'userID': this.$store.state.userID}
        this.$store.dispatch('notesModule/getAll', payload)

        if(this.storedNotes === null || this.storedNotes === undefined) {
            this.storedNotes = []
        } else {
            for (let i=0; i<this.storedNotes.length; i++) {
            this.titles.push(this.storedNotes[i].title)
            this.notes.push(this.storedNotes[i].note)
        }
        }
        localStorage.setItem('currentComponent', 'Notes')
    },
    computed: {
    ...mapState(['notesModule']),
    storedNotes () {
      return (!this.notesModule.notes.loading && this.notesModule.notes.data) || []
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