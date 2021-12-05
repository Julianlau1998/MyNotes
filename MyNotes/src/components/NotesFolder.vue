<template>
  <div id="wrapper">
        <h2
            class="headerTwo"
            >
            {{folder.title}}
        </h2>
        <ul
            id="listParent"
        >
            <li
                v-for="(note, idx) in notes"
                v-bind:key="idx"            >
                <span>
                    <button
                        class="noteDiv" 
                        v-bind:style="{backgroundColor: folder.color}"
                        @click="openNote(note.id)"
                    >
                        <h5><b>{{note.title.substring(0,11)}}</b></h5>
                    </button>
                </span>
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
import { mapState } from 'vuex' 

export default {
    name: 'NotesFolder',
    data () {
        return {
            id: this.$route.params.id
        }
    },
    created () {
        let payload = {'userID': this.$auth.user.sub, 'folderID': this.id}
        this.$store.dispatch('notesModule/getByCategory', payload)
        payload = {'userID': this.$auth.user.sub, 'id': this.id}
        this.$store.dispatch('foldersModule/getOne', payload)
    },
    computed: {
    ...mapState(['notesModule']),
    ...mapState(['foldersModule']),
        notes () {
            return (!this.notesModule.notes.loading && this.notesModule.notes.data) || []
        },
        folder () {
            return (!this.foldersModule.folder.loading && this.foldersModule.folder.data) || []
        },
    },
    methods: {
        settings () {
            this.$router.push('/settings')
        },
        openNote (id) {
            this.$store.state.transitionName = 'swipe-left'
            this.$store.state.id = id
            this.$router.push(`/Note/${id}`)
        },
        newNote () {
            this.$store.state.transitionName = 'fade'
            localStorage.setItem('currentFolder', this.id)
            this.$router.push('/newNote')
        }
    }
}
</script>

<style scoped>
    .headerTwo {
        font-size: 2.2rem;
        color: lightgray;
        margin-top: 1.5rem;
        margin-bottom: 0.8rem;
    }
    ul li {
        display: inline-block !important;
    }
    .delete {
        left: 1rem !important;
        top: 0.5rem !important;
    }
    #settingsImage {
        top: 0.5rem !important;
    }
    @media (min-width: 1000px) { 
      .noteDiv {
          width: 11rem;
          height: 9rem;
      }
    }
</style>