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
  created() {
    if (this.$workbox) {
      this.$workbox.addEventListener("waiting", () => {
        this.showUpgradeUI = true;
      });
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
}
body {
  background-color: #0f1820;
  margin-top: 0.5rem;
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
    }

    /*** Note / NewNote ***/
    .arrow {
        width: 2.5rem;
        height: 3rem;
        position: absolute;
        left: 1rem;
        top: 0.3rem;
    }
    .delete {
        width: 2.2rem;
        position: absolute;
        right: 1rem;
        top: 0.6rem;
    }
    #title {
        font-size: 1.5rem;
        background-color: rgb(23, 25, 27);
        background-color: rgb(54, 61, 68);
        color: lightgray;
        border: none;
    }
    #note {
        height: 68vh;
        font-size: 1.8rem;
        background-color: rgb(54, 61, 68);
        color: lightgray;
        border: none;
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
        z-index: -1;

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
    background-color: transparent;
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
    transition: opacity .17s;
  }
  .fade-enter-to, .fade-leave {
    opacity: 1;
  }
  .fade-enter, .fade-leave-to {
    opacity: 0;
  }
  .dg-main-content {
    margin-top: 12rem !important;
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
    margin-left: 3rem !important;
    font-size: 1.2rem !important;
  }

  .dg-btn--ok {
    border: none !important;
    border-radius: 8px !important;
    color: lightgray !important;
    background-color: transparent !important;
    box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                1px 1px 5px 2px black !important;
    margin-right: 3rem !important;
    padding-left: 0.6rem !important;
    padding-right: 0.6rem !important;
    font-size: 1.2rem !important;
  }
</style>
