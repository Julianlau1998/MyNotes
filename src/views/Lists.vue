<template>
    <div id="app">
        <h1 class="header">My<span style="color:rgb(215, 0, 0);">N</span>otes</h1>
        <hr class="whiteLine">
        <h2><router-link to="/" style="color: white;">Notes</router-link>
        <span style="color: rgb(0, 215, 215); opacity: 0.7; margin-left: 0.6rem;">Lists</span></h2>
        <ul id="listParent">
            <li v-for="(list, idx) in storedLists" v-bind:key="idx">
                <button class="noteDiv" @click="openList(list.id)">
                    <h5><b>{{list.title}}</b></h5>
                </button>
                <hr id="redLine">
            </li>
        </ul>
        <br><br><br><br><br><br>

        <router-link to="/NewList"><div class="plusButton">+</div></router-link>
    </div>
</template>

<script>
import router from '../router'
/* import Vue from 'vue'
import Vue2TouchEvents from 'vue2-touch-events'

Vue.use(Vue2TouchEvents) */

export default {
    name: 'Lists',
    data () {
        return {
            storedLists: JSON.parse(localStorage.getItem('lists')),
            titles: [],
            lists: []
        }
    },
    methods: {
        openList (id) {
            this.$store.state.id = id
            router.push(`/list/${id}`)
        },
        swipeHandler () {
            this.$router.push('/')
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