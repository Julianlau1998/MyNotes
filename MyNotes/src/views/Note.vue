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
            <div class="form-group">
                <ValidationProvider name="email" rules="required" v-slot="{ errors }">
                    <button to="/" v-if="focusValue" class="saveHidden">
                        <img
                        src="../assets/haken.png"
                        alt="delete icon"
                        class="deleteNew"
                        v-if="focusValue"
                        >
                    </button>
                    <input
                    type="test"
                    class="form-control"
                    id="title"
                    autocomplete="off"
                    :placeholder="$t('text.note.title')"
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
                  :placeholder="$t('text.note.note')"
                  value="note"
                  v-model="note"
                  ref="note"
                  @focus="focusValue=true"
                > 
                </textarea>        
            </div>
            <button class="saveButton" type="submit">
                {{ $t('text.note.save') }}
            </button>
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
import { mapState } from 'vuex' 


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
            title: '',
            note: '',
            postNote: {
                id: '',
                title: '',
                body: ''
            },
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
            this.postNote.id = this.id
            this.postNote.title = this.title
            this.postNote.body = this.note
            const payload = {'note': this.postNote}
            this.$store.dispatch('notesModule/put', payload)
            .then (() => {
                this.save = true
                router.push('/')
            })
            .catch((err) => {
                console.log(err)
            })
        },

        deleteNote () {
            const payload = {'id': this.id }
            this.$store.dispatch('notesModule/deleteOne', payload)
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
    },
    created () {
        const payload = {'id': this.id}
        this.$store.dispatch('notesModule/getOne', payload)
        
        document.getElementById('body').style.overflow = 'hidden'
        if(navigator.share !== undefined) {
            this.shareAvailable = true
        }
        this.$store.state.currentComponent = 'Notes'
    },

    computed: {
        ...mapState(['notesModule']),
        storedNote () {
            return (!this.notesModule.note.loading && this.notesModule.note.data) || []
        }
    },

    watch: {
        storedNote: function (val) {
            this.title = val.title
            this.note = val.body
            this.originalTitle = this.title
            this.originalNote = this.note
        }
    },

    beforeRouteLeave (to, from, next) {
        if(
            this.originalNote !== this.note ||
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
    top: 0.2rem;
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
textarea {
    z-index: 10;
}
</style>