<template>
    <div id="app">
        <h1 class="header">
            My<span style="color:rgb(215, 0, 0);">N</span>otes
        </h1>
        <hr class="whiteLine">
        <br>
        <div>
            <hr class="splitLine"> 
            <h3 @click="deleteData()" class="setting">
                <img src="../assets/trash.png" alt="trash icon" id="trash">
                Delete App Data
            </h3>
            <hr class="splitLine">
        </div>
        <button
            @click="$router.push('/')"
            class="saveButton"
            id="home"
        >
                Home
        </button>
    </div>
</template>

<script>
export default {
    data () {
        return {
        }
    },
    methods: {
        deleteData () {
            const vm = this
            this.$dialog.confirm("Are You sure you want to delete all notes and lists?\n \n You won't be able to restore them.")
                .then (function () {
                    localStorage.clear()
                    vm.$router.push('/')
                })
                .catch (function () {
                })
        }
    },
    mounted () {

    },
    created () {
        // this.$store.state.transitionName = 'swipe-right'
    }
}
</script>

<style scoped>
    .splitLine{
      border-top: 3px solid rgb(134, 134, 134);
      width: 100%;
      opacity: 0.5;
      margin-bottom: 1.3rem;
    }
    #trash {
        width: 2.2rem;
        margin-bottom: 0.5rem;
        margin-right: 0.5rem;
    }
    .setting {
        text-align: left;
        margin-left: 1rem;
    }
    #home {
        margin-bottom: 1.5rem;
    }
</style>