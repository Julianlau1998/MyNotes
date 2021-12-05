<template>
  <div id="wrapper">
        <button id="deleteButton">
            <img
                src="../assets/trash.png"
                alt="delete icon"
                class="delete"
                @click="deleteFolder"
                ref="trashcan"
            >
            <router-link to="/">
                    <img
                    src="../assets/arrow.png"
                    alt="back arrow"
                    class="arrow"
                    ref="backArrow"
                    >
            </router-link>
        </button>

        <h1 class="header">
            My<span style="color:rgb(215, 0, 0);">N</span>otes
        </h1>
        <hr class="whiteLine">
        <lists-folder
            v-if="currentFolder=='Lists'"
        />
        <notes-folder 
            v-if="currentFolder=='Notes'"
        />
  </div>
</template>

<script>
import ListsFolder from '../components/ListsFolder.vue'
import NotesFolder from '../components/NotesFolder.vue'

export default {
  components: { NotesFolder, ListsFolder },
    data () {
        return {
            currentFolder: localStorage.getItem('folder'),
            id: this.$route.params.id
        }
    },
    created () {
        this.$store.state.transitionName = 'fade'
    },
    computed: {
    },
    methods: {
        settings () {
            this.$router.push('/settings')
        },
        deleteFolder () {
            const vm = this
            this.$dialog.confirm(this.$t('text.deleteFolderAlert'))
                .then (function () {
                    const payload = {'userID': vm.$auth.user.sub, 'id': vm.id}
                    vm.$store.dispatch('foldersModule/deleteOne', payload)
                    vm.$router.push('/')
                })
                .catch (function () {
                })
        }
    }
}
</script>

<style scoped>
    .delete {
        right: 1rem !important;
        top: 0.6rem !important;
        background: #0f1820;
    }
    #deleteButton {
        margin-left: -210rem;
        width: 0;
    }
    .arrow {
        margin-top: 0.4rem;
    }
    #settingsImage {
        margin-top: 0.2rem;
    }
    .header {
        margin-top: 1rem;
    }
</style>