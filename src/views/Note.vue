<template>
    <div id="app">
        <div>
            <router-link to="/">
                    <img
                    src="../assets/arrow.png"
                    alt="back arrow"
                    class="arrow"
                    ref="backArrow"
                    >
            </router-link>
        </div>
        <div>
            <router-link to="/" v-if="!focusValue">
                    <img
                    src="../assets/trash.png"
                    alt="delete icon"
                    class="delete"
                    @click="deleteNote"
                    ref="trashcan"
                    >
            </router-link>
        </div>
        <br><br>
        <ValidationObserver v-slot="{ handleSubmit }">
        <form @submit.prevent="handleSubmit(onSubmit)">
            <button to="/" v-if="focusValue" class="saveHidden">
                <img
                src="../assets/haken.png"
                alt="delete icon"
                class="deleteNew"
                type="submit"
                v-if="focusValue"
                >
            </button>
            <div class="form-group">
                <ValidationProvider name="email" rules="required" v-slot="{ errors }">
                    <input
                    type="test"
                    class="form-control"
                    id="title"
                    autocomplete="off"
                    placeholder="Title"
                    value="title"
                    v-model="title"
                    @focus="focusValue=true"
                    >        
                <span class="errorMessage">{{ errors[0] }}</span>
                </ValidationProvider>
            </div>
            <div class="form-group">
                <textarea
                  type="test"
                  class="form-control"
                  id="note"
                  placeholder="Note"
                  value="note"
                  v-model="note"
                  ref="note"
                  @focus="focusValue=true"
                > 
                </textarea>        
            </div>
            <button class="saveButton" type="submit">Save</button>
        </form>
        </ValidationObserver>
        <img
        v-if="shareAvailable"
        src="../../public/img/share.png"
        alt="share"
        v-on:click="share"
        id="share"
        ref="share"
        >
    </div>
</template>

<script>
import router from '../router'
import { ValidationProvider, ValidationObserver, extend } from 'vee-validate';
import { required } from 'vee-validate/dist/rules';

extend('required', {
  ...required,
  message: 'This field is required'
});

export default {
    components: {
        ValidationProvider,
        ValidationObserver
    },
    data () {
        return {
            id: this.$route.params.id,
            notes: JSON.parse(localStorage.getItem('notes')),
            title: '',
            note: '',
            currentObject: {title: '', note: '', id: ""},
            focusValue: false,
            shareAvailable: false,
            shareNote: '',
            originalTitle: '',
            originalNote: '',
            save: false
        }
    },
    methods: {
        onSubmit () {
            for (let i = 0; i < this.notes.length; i++) {
                if (this.notes[i].id === this.id) {
                    this.notes[i].title = this.title
                    this.notes[i].note = this.note                }
            }
            localStorage.setItem('notes', JSON.stringify(this.notes))
            this.save = true
            router.push('/')
        },
        deleteNote () {
            for (let i in this.notes) {
                if (this.notes[i].id === this.id) {
                    this.notes.splice(i, 1)
                }
            }
            localStorage.setItem('notes', JSON.stringify(this.notes))
        },
        share () {
            navigator.share({
                "title": this.title,
                "text": this.note
            })
        }
    },
    mounted () {
        this.$store.state.transitionName = 'swipe-right'
        for (let i = 0; i < this.notes.length; i++) {
            if (this.notes[i].id === this.id) {
                this.title = this.notes[i].title
                this.note = this.notes[i].note
                this.originalTitle = this.title
                this.originalNote = this.note
            }
        }
        setTimeout(() => {
            this.$refs.backArrow.style.opacity = 1
            this.$refs.trashcan.style.opacity = 1
            this.$refs.share.style.opacity = 1
        }, 200);
    },
    created () {
        if(navigator.share !== undefined) {
            this.shareAvailable = true
        }
        this.$store.state.currentComponent = 'Notes'
    },
    beforeRouteLeave (to, from, next) {
        if(
            this.originalNote !== this.note ||
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
body {
    overflow-x: hidden;
} 
.saveHidden {
    position: absolute;
    right: 0.5rem;
    top: 0.75rem;
    background: transparent;
    border: none;;
}
.deleteNew {
      width: 2.8em;
      background: none;
      margin-top: -0.3rem;
  }
#share {
    width: 7rem;
    position: absolute;
    bottom: -0.5rem;
    right: -0.7rem;
    cursor: pointer;
}
.arrow,
.delete,
#share {
    opacity: 0;
}
</style>