<template>
    <transition
      :name= transitionName 
      mode="out-in"
      ref="transition"
    >
        <router-view
        ></router-view>
    </transition>
</template>

<script>
export default {
  data () {
    return {
    }
  },
  async created() {
    setInterval(() => {
      this.$auth.getTokenSilently()
      .then((token) => {
         sessionStorage.setItem('token', `bearer ${token}`)
      })
    }, 3600)

    if (this.$workbox) {
      this.$workbox.addEventListener("waiting", () => {
        this.showUpgradeUI = true;
      });
    }
    if (this.$store.state.userID == '') {
      this.$store.state.userID = localStorage.getItem('userID')
    }
  },

  methods: {
    async accept() {
      this.showUpgradeUI = false
      await this.$workbox.messageSW({ type: "SKIP_WAITING" });
    }
  },
  computed: {
    transitionName () {
      return this.$store.state.transitionName
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: rgb(201, 201, 201);
  position: absolute;
  width: 100%;
  height: 100%;
}
body {
  background-color: #0f1820;
  margin-top: 0.5rem;
  overflow: hidden;
  text-align: center;
}
    .plusButton {
        position: fixed;
        bottom: 1rem;
        padding-bottom: 3.5rem;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 4.9rem;
        height: 4.9rem;
        font-size: 3.2em;
        background-color: transparent;
        border-radius: 100%;
        border: none;
        box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                     1px 1px 5px 2px black;
        color: white;
        cursor: pointer;
    }
    .plusButton:active {
      box-shadow: -1px -1px 4px 0px black,
                  1px 1px 3px 0px rgb(133, 133, 133);
      }
    .line {
        border-top: 6px dashed rgb(158, 63, 0);
        border-bottom: none;
        border-right: transparent;
        border-left: transparent;
        color: transparent;
    }
    .header {
      margin: 2rem 0 1rem;
      font-family: Avenir, Helvetica, Arial, sans-serif;
      letter-spacing: 5px;
      font-size: 3.5rem;
      color: lightgray;
    }

    .noteDiv {
        outline: none;
    }

    /*** Note / NewNote ***/
    .arrow {
        width: 2.5rem;
        height: 3rem;
        position: absolute;
        left: 1rem;
        top: -0.1rem;
    }
    .delete {
        width: 2.2rem;
        position: absolute;
        right: 1rem;
        top: 0.2rem;
    }
    #title {
        font-size: 1.5rem;
        background-color: rgb(23, 25, 27);
        background-color: rgb(54, 61, 68);
        color: lightgray;
        border: none;
    }
    #note {
        height: 84vh;
        font-size: 1.8rem;
        background-color: rgb(54, 61, 68);
        color: lightgray;
        border: none;
        z-index: 100;
    }
    .saveButton {
        background-color: rgb(4, 82, 17);
        width: 10rem;
        height: 3rem;
        color: white;
        font-weight: bolder;
        font-size: 1.2rem;
        position: absolute;
        bottom: 1rem;
        left: 50%;
        transform: translateX(-50%);
        z-index: -1;
        outline: none;
    }
    .saveButton:active {
      box-shadow: -1px -1px 4px 0px black,
                  1px 1px 3px 0px rgb(133, 133, 133);
      }
      .saveButton:focus {
        outline: none;
      }
    .errorMessage{
      color: red;
    }
    .saveButton {
        width: 10rem;
        height: 3rem;
        font-weight: bolder;
        font-size: 1.2rem;
        position: absolute;
        bottom: 1.5rem;
        left: 50%;
        transform: translateX(-50%);
        z-index: 1;
        border-radius: 0.7rem;
        background-color: transparent;
        border: none;
        box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                    1px 1px 5px 2px black;
        color: lightgray;
        font-weight: 900;
        font-size: larger;
        margin-top: 5rem;
    }
    .saveHidden {
      background: transparent;
    }
    .saveHidden:focus {
      outline: none;
    }
    .whiteLine{
      border-top: 3px solid rgb(134, 134, 134);
      width: 12.5rem;
      margin-top: -0.5rem;
      margin-bottom: 1rem;
    }
    #redLine{
        border-top: 3px solid rgb(215, 0, 0);
        margin-top: -0.5rem;
        margin-bottom: 1rem;
        margin-left: -2.5rem;
    }

    /* Lists */
    #checkbox,
    #checkboxChecked {
    width: 1.6rem;
    height: 1.6rem;
    background-color: transparent !important;
    border: 2px solid lightgray;
    border-radius: 20px;
    position: relative;
    top: 1.9rem;
    opacity: 0.5;
    padding-top: 0rem;
    cursor: pointer;
  }
  #checkboxChecked {
    margin-left: -0.47rem
  }
  .checked {
      width: 7px;
      height: 7px;
      border-radius: 20px;
      background-color: rgb(0, 215, 215);
      position: relative;
      left: 0.47rem;
      top: 0.47rem;
  }
  .marginLeft {
      margin-left: 3.5rem
  }
  #item,
  #doneItem {
      position: relative;
      bottom: 0.1rem;
  }

  .neomorph {
    border: none;
    box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                 1px 1px 5px 2px black;
    color: lightgray;
    font-weight: 500;
    background-color: #0f1820 !important;
  }
  .reverseNeomorph {
    border: none;
    box-shadow: -1px -1px 4px 0px black,
                  1px 1px 3px 0px rgb(133, 133, 133);
    color: lightgray;
    font-weight: 500;
  }
  #settingsImage {
    width: 2.2rem;
    position: absolute;
    right: 0.7rem;
    top: 0.3rem;
    cursor: pointer;
  }

  .checkImage {
    width: 1.7rem;
    background-color: transparent;
    position: relative;
    bottom: 0.7rem;
    right: 0.21rem;
  }

  #login {
        width: 6rem;
        height: 3rem;
        border-radius: 1rem;
        background-color: transparent;
        border: none;
        box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                    1px 1px 5px 2px black;
        color: lightgray;
        font-weight: 960;
        font-size: larger;
        margin-top: 1rem;
  }

  .disabled {
    opacity: 0.3;
  }
  .whiteText {
    color: lightgray;
    margin: 1rem;
  }

  #newFolder,
  #newNote {
    width: 9.5rem;
    height: 3.5rem;
    padding: 0.6rem 0 0rem 0;
    background-color: aliceblue;
    position: fixed;
    bottom: 7.3rem;
    transform: translate(-50%, -50%);
    background-color: #0f1820;
    border: none;
    box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                 1px 1px 5px 2px black;
    color: white;
    font-weight: 700;
    font-size: 1.5rem;
    cursor: pointer;
    border-radius: 5px;
    z-index: 20;
  }

#newFolder{
    left: 50%;
    margin-left: -6rem;
}
#newNote{
    left: 50%;
    margin-left: 6rem;
}
.newFolderElement {
    color: lightgray;
    font-weight: 600;
    width: 8rem;
}
.divide {
    border: 3px solid black
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

  /* Route change animation */
  .swipe-left-enter-active,
  .swipe-left-leave-active {
    transition: transform 0.15s;
  }
  .swipe-left-enter {
    transform: translateX(100%);
  }
  .swipe-left-enter-to {
    transform: translateX(0%);
  }
  /* .swipe-left-leave {
    transform: translateX(0%);
  }
  .swipe-left-leave-to {
    transform: translateX(-100%);
  } */

  /* Swipe Right */
  .swipe-right-enter-active,
  .swipe-right-leave-active {
    transition: transform 0.15s;
  }
  /* .swipe-right-enter {
    transform: translateX(-100%);
  }
  .swipe-right-enter-to {
    transform: translateX(0%);
  } */
  .swipe-right-leave {
    transform: translateX(0%);
  }
  .swipe-right-leave-to {
    transform: translateX(100%);
  }

  /* Fade In-Out */
  .fade-enter-active, .fade-leave-active {
    transition: opacity .2s;
  }
  .fade-enter-to, .fade-leave {
    opacity: 1;
  }
  .fade-enter, .fade-leave-to {
    opacity: 0;
  }

  /* New Folder modal animation */
  .new-folder-enter-active, .fade-leave-active {
    transition: all .2s;
  }
  .new-folder-enter-to, .new-folder-leave {
    opacity: 1;
  }
  .new-folder-enter, .new-folder-leave-to {
    opacity: 0;
  }

  .dg-main-content {
    position: absolute;
    top: 0rem;
    left: 50%;
    transform: translate(-50%, 0%);
    background-color: #0f1820 !important;
    color: rgb(201, 201, 201) !important;
    border-radius: 20px !important;
  }
  .dg-content {
    font-size: 1.6rem !important;
  }
  .dg-btn-content {
    border: none !important;
    background-color: transparent !important;
  }
  .dg-btn--cancel {
    border: none !important;
    border-radius: 8px !important;
    color:  rgb(0, 215, 215) !important;
    background-color: transparent !important;
    box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                1px 1px 5px 2px black !important;
    margin-left: 1rem !important;
    font-size: 1.2rem !important;
  }

  .dg-btn--ok {
    border: none !important;
    border-radius: 8px !important;
    color: lightgray !important;
    background-color: transparent !important;
    box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                1px 1px 5px 2px black !important;
    margin-right: 3.8rem !important;
    padding-left: 0.6rem !important;
    padding-right: 0.6rem !important;
    font-size: 1.2rem !important;
  }

  @media (max-height: 1300px) {
    #note {
        height: 80.5vh;
    }
  }
  @media (max-height: 1000px) {
    #note {
        height: 77.5vh;
    }
  }
  @media (max-height: 886px) {
    #note {
        height: 75vh;
    }
  }
  @media (max-height: 790px) {
    #note {
        height: 73.5vh;
    }
  }
  @media (max-height: 750px) {
    #note {
        height: 72.5vh;
    }
  }

  @media (max-height: 729px) {
    #note {
        height: 70.5vh;
    }
  }

  @media (max-height: 686px) {
    #note {
        height: 68.5vh;
    }
  }

  @media (max-height: 631px) {
    #note {
        height: 64.5vh;
    }
  }
  @media (max-height: 547px) {
    #note {
        height: 59.5vh;
    }
  }
  @media (max-height: 463px) {
    #note {
        height: 54.5vh;
    }
  }
  @media (max-height: 421px) {
    #note {
        height: 49.5vh;
    }
  }
  @media (max-height: 381px) {
    #note {
        height: 44.5vh;
    }
  }
  @media (max-width: 370px)  {
    #newFolder,
    #newNote {
      width: 7.3rem;
      height: 1.7rem;
      padding-bottom: 2rem;
      font-size: 1.1rem;
    }
    #newFolder {
      margin-left: -5rem;
    }
    #newNote {
      margin-left: 5rem;
    }
  }
</style>
