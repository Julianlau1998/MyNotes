<template>
    <div id="app">
        <router-link to="/">
            <img
                src="../assets/arrow.png"
                alt="back arrow"
                class="arrow"
                ref="backArrow"
            >
        </router-link>

        <router-link to="/" v-if="!focusValue">
            <img
                src="../assets/trash.png"
                alt="delete icon"
                class="delete"
                @click="deleteList"
                ref="trashcan"
            >
        </router-link>
        <br><br>
        <ValidationObserver v-slot="{ handleSubmit }">
        <form @submit.prevent="handleSubmit(onSubmit)" class="form">
            <div class="form-group">
                <ValidationProvider rules="required" v-slot="{ errors }">
                    <button to="/" class="saveHidden" v-if="focusValue">
                        <img
                            src="../assets/haken.png"
                            alt="delete icon"
                            class="deleteNew"
                        >
                    </button>
                    <input 
                        type="test"
                        class="form-control"
                        id="title"
                        :placeholder="$t('text.list.title')"
                        v-model="title"
                        @click="focusValue=true"
                        autocomplete="off"
                        >    
                    <span class="errorMessage">{{ errors[0] }}</span>
                </ValidationProvider>    
            </div>
            <br>
            <div class="form-group scroll">
                <div type="test" class="form-control" id="note"  placeholder="List">
                    <span 
                        v-if="listElements && listElements.length != 0"
                        class="subTitle"
                    >
                        {{ $t('text.list.toDo') }}
                    </span>
                    <ul id="itemsList">
                        <draggable
                            v-model="listElements"
                            :delay="200"
                            :delay-on-touch-only="true"
                        >
                            <li v-for="(item, itemKey) in listElements" :key="itemKey" ref="list">
                                <span
                                    v-touch:swipe="swipeItem(itemKey)"
                                >
                                    <div 
                                        id="checkbox"
                                        @click="itemDone(item)"
                                        slot="footer"
                                    >
                                    </div>
                                        <span
                                            @click="edit(item)"
                                            class="marginLeft bottom item"
                                            id="item"
                                        >   
                                            {{item}} 
                                        </span>
                                    <hr class="whiteLine">
                                </span>
                            </li>
                        </draggable>
                    </ul>
                    <span
                    v-if="doneItems && doneItems.length != 0"
                    class="subTitle"
                    >
                        <br>
                        {{ $t('text.list.done') }}
                    </span>
                    <ul>
                        <draggable
                            v-model="doneItems"
                            :delay="200"
                            :delay-on-touch-only="true"
                        >
                            <li v-for="(doneItem, itemKey) in doneItems" id="done" :key="itemKey">
                                <span 
                                  v-touch:swipe="swipeDoneItem(itemKey)"
                                >
                                    <div 
                                        id="checkboxChecked"
                                        class="checked"
                                        @click="itemNotDone(doneItem)"
                                    >
                                        <img
                                            src="../assets/haken.png"
                                            alt=""
                                            class="checkImage"
                                            >
                                    </div> 
                                    &nbsp;
                                            <del
                                                class="marginLeft"
                                                id="doneItem"
                                            >
                                                {{doneItem}}
                                            </del>
                                    <hr class="whiteLine">
                                </span>
                            </li>
                        </draggable>
                    </ul>
                    <br><br>
                </div>        
            </div>
        </form>
        </ValidationObserver>
        <img
            src="../../public/img/share.png"
            alt="share"
            v-on:click="share"
            v-if="shareAvailable"
            id="share"
            ref="share"
        >
        <input
            class="form-control newNote"
            :placeholder="$t('text.list.newItem')"
            v-model="listItem"
            @click="focusValue=true"
            v-on:keyup.enter="addItem()"
            ref="add"
        />
        <button
            class="neomorph"
            id="addButton"
            @click="addItem"
            ref="addButton"
        > 
            {{ $t('text.list.add') }} 
        </button>
    </div>
</template>

<script>
import { ValidationProvider, ValidationObserver, extend } from 'vee-validate';
import { required } from 'vee-validate/dist/rules';
import draggable from 'vuedraggable'
import { mapState } from 'vuex' 


extend('required', {
  ...required,
  message: 'This field is required'
});
export default {
    name: 'SingleList',
    components: {
        ValidationProvider,
        ValidationObserver,
        draggable,
    },
    data () {
        return {
            id: this.$route.params.id,
            title: '',
            originalTitle: '',
            postList: {
                id: '',
                user_id: '',
                title: '',
                listElements: [],
                doneItems: [],
            },
            listElements: [],
            originalListElements: [],
            listItem: '',
            doneItems: [],
            originalDoneItems: [],
            focusValue: false,
            sharedList: '',
            shareAvailable: false,
            alreadyAsked: false,
            save: false,
            listElementsChanged: false,
            watcherCounter: 0,
        }
    },

    mounted () {
        //set animation type
        this.$store.state.transitionName = 'swipe-right'
    },
    created () {
        const payload = {'id': this.id}
        this.$store.dispatch('listsModule/getOne', payload)
        this.lists = this.listsList

        if(navigator.share !== undefined) {
            this.shareAvailable = true
        }
        this.$store.state.currentComponent = 'Lists'
        document.getElementById('body').style.overflow = 'hidden'
    },

    computed: {
        ...mapState(['listsModule']),
        storedList () {
            return (!this.listsModule.list.loading && this.listsModule.list.data) || []
        }
    },

    watch: {
        storedList: function (val) {
            this.title = val.title
            this.listElements = val.list || []
            this.doneItems = val.doneItems
            this.originalTitle = this.title
            // this.originalListElements = val.listElements.filter(el => el == el)
            this.originalNote = this.note
            this.originalDoneItems = this.doneItems
        },
        listElements: function () {
            if (this.watcherCounter > 1) {
                this.listElementsChanged = true
                this.watcherCounter++
                this.focusValue=true
            } else {
                this.watcherCounter++
            }
        }
    },

    methods: {
        onSubmit (error) {
            if (error === undefined) {
                this.postList.id = this.id
                this.postList.title = this.title
                this.postList.list = this.listElements
                this.postList.doneItems = this.doneItems
                const payload = {'list': this.postList}
                this.$store.dispatch('listsModule/put', payload)
                .then (() => {
                    this.save = true
                    this.$router.push('/')
                })
                .catch((err) => {
                    console.log(err)
                })
            }
        },
        addItem () {
            if (this.$refs.add.value !== '') {
                this.listElements.unshift(this.$refs.add.value)
                this.listItem = ''
                this.$refs.add.value = ''
                this.$refs.add.focus();
            }
        },
        itemDone(item) {
            this.listElements = this.listElements.filter(el => el != item)
            this.doneItems.unshift(item)
            this.focusValue=true
        },
        itemNotDone(item) {
            this.doneItems = this.doneItems.filter(el => el != item)
            this.listElements.unshift(item)
            this.focusValue=true
        },
        deleteList () {
            const payload = {'id': this.id}
            this.$store.dispatch('listsModule/deleteOne', payload)
            this.$router.push('/')
        },
        edit(item) {
            if (this.$refs.add.value === '') {
                this.listElements = this.listElements.filter(el => el != item)
                this.$refs.add.focus()
                this.$refs.add.value = item
                this.listItem = item
            }
        },
        share () {
            this.listElements.forEach(el => {
                console.log(el)
                this.sharedList += ('- ' + el + '\n')
            })
            navigator.share({
                "title": this.title,
                "text": this.sharedList
            })
            this.sharedList = ''
        },
        swipeItem (itemIndex) {
            const vm = this
            return function(direction) {
                if (direction==='left' || direction==='right') {
                    const item = vm.listElements[itemIndex]
                    vm.listElements = vm.listElements.filter(el => el != item)
                }
            }
        },
        swipeDoneItem (itemIndex) {
            const vm = this
            return function(direction) {
                if (direction==='left' || direction==='right') {
                    const item = vm.doneItems[itemIndex]
                    vm.doneItems = vm.doneItems.filter(el => el != item)
                }
                
            }
        }
    },

    beforeRouteLeave (to, from, next) {
        if(
            this.listElementsChanged ||
            this.originalDoneItems !== this.doneItems ||
            this.originalTitle !== this.title
        ) {
            if (this.save===false) {
                this.$dialog.confirm(this.$t('text.saveAlert'))
                .then (function () {
                    next()
                })
                .catch (function () {
                    next(false)
                })
            } else {
                next ()
            }
        }
        else {
            next()
        }
    },
    beforeMount() {
        window.addEventListener("beforeunload", event => {
            if (this.listElementsChanged ||
                this.originalDoneItems !== this.doneItems ||
                this.originalTitle !== this.title) {
                    setTimeout(50)
                    event.preventDefault()
                    // Chrome requires returnValue to be set.
                    return event.returnValue = ""
            } else {
                return
            }
        })
    }
}
</script>




// *** Styles ***


<style scoped>
  ul li {
      text-align: left;
      list-style-type: none;
      margin-left: -2rem;
      margin-bottom: 0.8rem;
  }
  input {
      color: lightgray;
      border: none;
      box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                1px 1px 5px 2px black;
  }
  .scroll{
      overflow-x: hidden;
      overflow-y: scroll;
  }
  #listElement {
      background: none;
      width: 75vw;
      margin-bottom: 0.5rem;
      background-color:  rgb(54, 61, 68);
  }
  #note{
      background-color: #0f1820;
      z-index: 10;
      height: 73vh;
      text-align: left;
  }
  #trashcan {
      width: 2rem;
  }
  .deleteNew {
      width: 2.8em;
      background: none;
  }
  .newNote{
      width: 75%;
      margin-left: 3%;
      position: absolute;
      bottom: 1rem;
      font-size: 1.5rem;
      background-color: rgb(23, 25, 27);
      background-color: rgb(54, 61, 68);
      color: lightgray;
      border: none;
  }
  #addButton {
      position: absolute;
      bottom: 1.2rem;
      right: 1rem;
      height: 2.5rem;
      border-radius: 5px;
      border: none;
      box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                    1px 1px 5px 2px black;
      color: lightgray;
      font-weight: 900;
      font-size: larger;
  }
  #addButton:focus {
      outline: none;
  }
  .whiteLine {
      width: 100vw;;
      margin-left: -1.2rem;
      border-top: 3px solid rgb(89, 89, 89);
      opacity: 0.5;
  }
  .saveHidden {
    position: absolute;
    right: 0.5rem;
    top: -0.1rem;
    background: transparent;
    border: none;;
}
input[type="checkbox"]:checked {
  color: red;
}
input[type="checkbox"] {
    width: 1rem;
    height: 1.2rem;
    box-shadow: none;
    margin-right: 10px;
}
#checkbox {
    width: 1.6rem;
    height: 1.6rem;
    background-color: red;
    border: 2px solid lightgray;
    border-radius: 20px;
    position: relative;
    top: 1.9rem;
    opacity: 1;
    padding-top: 0rem;
    cursor: pointer;
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
#done {
    margin-top: -0.5rem;
}
.subTitle {
    color: rgb(0, 215, 215);
    font-weight: 500;
}
.marginLeft {
    margin-left: 2.7rem
}
#item,
#doneItem {
    position: relative;
    bottom: 0.1rem;
    overflow-wrap: break-word; 
}
#share {
    width: 5rem;
    position: absolute;
    top: -1.15rem;
    left: 50%;
    transform: translateX(-50%);
    cursor: pointer;
}

 @media (min-width: 600px) { 
  .newNote {
      width: 83%;
  }   
 }
 @media (min-width: 800px) { 
  .newNote {
      width: 85%;
  }   
  #addButton {
      margin-right: 0.8rem;
  }
  /* Custom Scrollbar */
  ::-webkit-scrollbar {
    width: 5px;
    height: 0px;
  }

  ::-webkit-scrollbar-track {
    background: transparent;
  }

  ::-webkit-scrollbar-thumb {
    background: rgb(215, 0, 0);
    border-radius: 20px;
  }

  ::-webkit-scrollbar-thumb:hover {
    background: rgb(0, 215, 215);
    cursor: pointer;
  }
 }
 @media (min-width: 1300px) { 
  .newNote {
      width: 88%;
  }
  #addButton {
      margin-right: 1rem;
  }
 }


.list-enter-active, .list-leave-active {
  transition: all 0.5s;
}
.list-enter, .list-leave-to {
  opacity: 0;
  transform: translateX(-100px);
}
</style>