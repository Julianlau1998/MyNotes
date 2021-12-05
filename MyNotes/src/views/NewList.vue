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
        
        <br><br>
        <ValidationObserver v-slot="{ handleSubmit }">
        <form @submit.prevent="handleSubmit(onSubmit)">
            <div class="form-group">
                <ValidationProvider rules="required" v-slot="{ errors }">
                    <button to="/" class="saveHidden">
                        <img
                            src="../assets/haken.png"
                            alt="delete icon"
                            class="deleteNew"
                            ref="safe"
                        >
                    </button>
                    <input
                        type="test"
                        class="form-control"
                        id="title"
                        :placeholder="$t('text.list.title')"
                        v-model="title"
                        ref="title"
                        autocomplete="off"
                    >    
                    <span class="errorMessage">{{ errors[0] }}</span>
                </ValidationProvider>    
            </div>
            <br>
            <div class="form-group paddingTop scroll">
                <div type="test" class="form-control" id="note"  placeholder="List">
                    <span
                    v-if="listElements.length != 0"
                    class="subTitle"
                    >   
                        {{ $t('text.list.toDo') }}
                    </span>
                    <ul>
                        <draggable
                            v-model="listElements"
                            :delay="200"
                            :delay-on-touch-only="true"
                        >
                            <li v-for="(item, itemKey) in listElements" :key="itemKey">
                                <span 
                                  v-touch:swipe="swipeItem(itemKey)"
                                >
                                    <div 
                                        id="checkbox"
                                        @click="itemDone(item)"
                                    >
                                    </div>
                                    <span
                                    @click="edit(item)"
                                    class="marginLeft bottom"
                                    id="item"
                                    >
                                        {{item}} 
                                    </span>
                                    <hr class="whiteLine">
                                </span>
                            </li>
                        </draggable>
                    </ul>
                    <br>
                    <span
                    v-if="doneItems.length != 0"
                    class="subTitle"
                    >
                        {{ $t('text.list.done') }}
                    </span>
                    <ul>
                        <draggable
                            v-model="doneItems"
                            :delay="200"
                            :delay-on-touch-only="true"
                        >
                            <li v-for="(doneItem, itemKey) in doneItems" :key="itemKey">
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
                </div>        
            </div>
        </form>
        </ValidationObserver>
        <input
            class="form-control newNote"
            :placeholder="$t('text.list.newItem')"
            v-model="listItem"
            ref="add"
            v-on:keyup.enter="addItem()"
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

extend('required', {
  ...required,
  message: 'This field is required'
});

export default {
    components: {
        ValidationProvider,
        ValidationObserver,
        draggable
    },
    data () {
        return {
            title: '',
            list: '',
            postList: {
                id: '',
                user_id: '',
                folder_id: localStorage.getItem('currentFolder'),
                title: '',
                listElements: [],
                doneItems: [],
            },
            listsList: JSON.parse(localStorage.getItem('lists')),
            id: this.$uuidKey(),
            listElements: [],
            listItem: '',
            doneItems: [],
            save: false
        }
    },
    methods: {
        onSubmit (error) {
            if (error === undefined) {
                this.postList.id = this.id
                this.postList.userID = this.$auth.user.sub
                this.postList.title = this.title
                this.postList.list = this.listElements
                this.postList.doneItems = this.doneItems
                console.log(this.doneItems)

                const payload = {'list': this.postList,'userID': this.$auth.user.sub}
                this.$store.dispatch('listsModule/post', payload)
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
        },
        itemNotDone(item) {
            this.doneItems = this.doneItems.filter(el => el != item)
            this.listElements.unshift(item)
        },
        edit(item) {
            if (this.$refs.add.value === '') {
                this.listElements = this.listElements.filter(el => el != item)
                this.$refs.add.focus()
                this.$refs.add.value = item
                this.listItem = item
            }
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
    created () {
        document.getElementById('body').style.overflow = 'hidden'
        if (this.listsList === undefined) {
            this.listsList = []
        }
        if (this.id === null) {
            this.id = 0
        }
        this.$store.state.currentComponent = 'Lists'
    },
    mounted () {
        this.$refs.title.focus();
    },
    beforeRouteLeave (to, from, next) {
        if(
            this.listElements[0] !== undefined ||
            this.doneItems[0] !== undefined ||
            this.title !== ''
        ) {
            if (this.save === false) {
                this.$dialog.confirm(this.$t('text.saveAlert'))
                .then (function () {
                    next()
                })
                .catch (function () {
                    next(false)
                })
            } else {
                next()
            }
        }
        else {
            next()
        }
    }
}
</script>

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
      text-align: left;
      height: 73vh
  }
  #trashcan {
      width: 2rem;
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
      background-color: transparent;
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
  .deleteNew {
      width: 2.8em;
      background: none;
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
#blocker {
    width: 1.6rem;
    height: 1.6rem;
    background-color: black;
    position: relative;
    top: 1.9rem;
    opacity: 0;
    padding-top: 0rem;
}
.subTitle {
    color: rgb(0, 215, 215);
    font-weight: 500;
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
</style>