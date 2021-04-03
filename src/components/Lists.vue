<template>
    <div id="app">
        <ul id="listParent">
            <draggable
                v-if="sorting"
                @start="drag=true"
                @end="drag=false"
                v-model="storedLists"
            >
                <li v-for="(list, idx) in storedLists" v-bind:key="idx">
                    <button class="noteDiv" @click="openList(list.id)">
                        <h5><b>{{list.title.substring(0,11)}}</b></h5>
                    </button>
                    <hr id="redLine">
                </li>
            </draggable>

            <li v-else v-for="(list, idx) in storedLists" v-bind:key="idx">
                <span v-touch:touchhold="touchHoldHandler">
                    <button class="noteDiv" @click="openList(list.id)">
                    <h5><b>{{list.title.substring(0,11)}}</b></h5>
                    </button>
                </span>
                <hr id="redLine">
            </li>
        </ul>
        <div 
            class="plusButton"
            ref="plusButton"
            @click="newList"
        >
            +
        </div>
    </div>
</template>

<script>
import router from '../router'
import draggable from 'vuedraggable'
import Vue from 'vue'
import Vue2TouchEvents from 'vue2-touch-events'

Vue.use(Vue2TouchEvents)

export default {
    name: 'Lists',
    components: {draggable},
    data () {
        return {
            storedLists: JSON.parse(localStorage.getItem('lists')),
            titles: [],
            lists: [],
            sorting: false
        }
    },
    methods: {
        openList (id) {
            this.$store.state.transitionName = 'swipe-left'
            this.$store.state.id = id
            router.push(`/list/${id}`)
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
        }
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