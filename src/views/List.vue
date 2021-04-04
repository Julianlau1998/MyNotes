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
            <button to="/" class="saveHidden" v-if="focusValue">
                <img src="../assets/haken.png" alt="delete icon" class="deleteNew" type="submit">
            </button>
            <div class="form-group">
                <ValidationProvider rules="required" v-slot="{ errors }">
                    <input 
                        type="test"
                        class="form-control"
                        id="title"
                        placeholder="Title"
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
                        v-if="listElements.length != 0"
                        class="subTitle"
                    >
                        To-Do:
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
                    v-if="doneItems.length != 0"
                    class="subTitle"
                    >
                        <br>
                        Done:
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
            placeholder="New Item"
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
            add 
        </button>
    </div>
</template>

<script>
import { ValidationProvider, ValidationObserver, extend } from 'vee-validate';
import { required } from 'vee-validate/dist/rules';
import draggable from 'vuedraggable'

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
            lists: JSON.parse(localStorage.getItem('lists')),
            currentObject: {
                title: '',
                listElements: [],
                doneItems: [],
                id: this.$uuidKey()
            },
            listsList: JSON.parse(localStorage.getItem('lists')),
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
    methods: {
        onSubmit () {
            if(this.listsList===null) {
                this.listsList = []
            }
            for (let i = 0; i < this.listsList.length; i++) {
                if (this.listsList[i].id === this.id) {
                    this.listsList[i].title = this.title
                    this.listsList[i].listElements = this.listElements
                    this.listsList[i].doneItems = this.doneItems                }
            }
            localStorage.setItem('id', this.currentObject.id)
            localStorage.setItem('lists',JSON.stringify(this.listsList))
            this.save = true
            this.$router.push('/')
        },
        addItem () {
            if (this.$refs.add.value !== '') {
                this.listElements.unshift(this.listItem)
                console.log(this.listElements)
                this.listItem = ''
                this.$refs.add.focus()
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
            for (let i in this.lists) {
                if (this.lists[i].id === this.id) {
                    this.lists.splice(i, 1)
                }
            }
            localStorage.setItem('lists', JSON.stringify(this.lists))
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
    mounted () {
        //set animation type
        this.$store.state.transitionName = 'swipe-right'
        //Get data from local storage and create backup
        for (let i = 0; i < this.lists.length; i++) {
            if (this.lists[i].id === this.id) {
                this.title = this.lists[i].title
                this.listElements = this.lists[i].listElements
                this.doneItems = this.lists[i].doneItems
                this.originalTitle = this.title
                this.originalListElements = this.listElements.filter(el => el == el)
                this.originalDoneItems = this.doneItems
            }
        }
    },
    created () {
        if(navigator.share !== undefined) {
            this.shareAvailable = true
        }
        this.$store.state.currentComponent = 'Lists'
        document.getElementById('body').style.overflow = 'hidden'
    },
    beforeRouteLeave (to, from, next) {
        if(
            this.listElementsChanged ||
            this.originalDoneItems !== this.doneItems ||
            this.originalTitle !== this.title
        ) {
            if (this.save===false) {
                this.$dialog.confirm('Are You sure you want to leave without saving? \n \n All changes would be lost.')
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
    },
    watch: {
        listElements: function () {
            if (this.watcherCounter > 0) {
                this.listElementsChanged = true
                this.watcherCounter++
                this.focusValue=true
            } else {
                this.watcherCounter++
            }
        }
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
    background-color: transparent;
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
#note {
    height: 73vh;
    text-align: left;
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
.checkImage {
    width: 1.7rem;
    background-color: transparent;
    position: relative;
    bottom: 0.7rem;
    right: 0.21rem;
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
::-webkit-scrollbar {
  width: 5px;
  height: 0px;
}

/* Custom Scrollbar */

::-webkit-scrollbar-track {
  background: transparent;
  cursor: pointer;
}

::-webkit-scrollbar-thumb {
  background: rgb(215, 0, 0);
  border-radius: 20px;
}

::-webkit-scrollbar-thumb:hover {
  background: rgb(0, 215, 215);
  cursor: pointer;
}
</style>