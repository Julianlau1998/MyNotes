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
                v-for="(list, idx) in lists"
                v-bind:key="idx"            >
                <span>
                    <button
                        class="noteDiv" 
                        v-bind:style="{backgroundColor: folder.color}"
                        @click="openList(list.id)"
                    >
                        <h5><b>{{list.title.substring(0,11)}}</b></h5>
                    </button>
                </span>
                <hr id="redLine">
            </li>
        </ul>
        <div 
            @click="newList()"
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
    name: 'ListsFolder',
    data () {
        return {
            id: this.$route.params.id
        }
    },
    created () {
        let payload = {'userID': this.$auth.user.sub, 'folderID': this.id}
        this.$store.dispatch('listsModule/getByCategory', payload)
        payload = {'userID': this.$auth.user.sub, 'id': this.id}
        this.$store.dispatch('foldersModule/getOne', payload)
    },
    computed: {
    ...mapState(['listsModule']),
    ...mapState(['foldersModule']),
        lists () {
            return (!this.listsModule.lists.loading && this.listsModule.lists.data) || []
        },
        folder () {
            return (!this.foldersModule.folder.loading && this.foldersModule.folder.data) || []
        },
    },
    methods: {
        settings () {
            this.$router.push('/settings')
        },
        openList (id) {
            this.$store.state.transitionName = 'swipe-left'
            this.$store.state.id = id
            this.$router.push(`/List/${id}`)
        },
        newList () {
            this.$store.state.transitionName = 'fade'
            localStorage.setItem('currentFolder', this.id)
            this.$router.push('/newList')
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
    .noteDiv {
        outline: none;
    }
    @media (min-width: 1000px) { 
      .noteDiv {
          width: 11rem;
          height: 9rem;
      }
    }
</style>