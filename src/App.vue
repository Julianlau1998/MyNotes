<template>
  <router-view/>
</template>

<script>
export default {
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
        top: 0.2rem;
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
</style>
