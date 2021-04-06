<template>
    <div>
        <div id="bottomNav">
            <h3>
                <img src="../assets/home.png" alt="Home Icon" id="home">
                <img src="../assets/folder.png" alt="Folder Icon" id="folder">
            </h3>
        </div>
    </div>
</template>

<script>
export default {
    name: 'bottomNav'
}
</script>

<style scoped>
    #bottomNav {
        width: 100%;
        height: 4rem;
        background-color: transparent;
        position: fixed;
        bottom: 0rem;
        border-top: 1px solid lightgray;
        padding-top: 0.5rem;
        z-index: -1;
    }
    #home,
    #folder {
        margin-bottom: -0.7rem;
        opacity: 0.5;
    }
    #home {
        width: 2.2rem;
        margin-right: 5rem;
    }
    #folder {
        width: 2.2rem;
        margin-left: 5rem;
    }
</style>