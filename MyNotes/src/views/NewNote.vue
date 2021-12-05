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
                <ValidationProvider name="email" rules="required" v-slot="{ errors }">
                    <button class="saveHidden">
                        <img
                            src="../assets/haken.png"
                            alt="delete icon"
                            class="deleteNew"
                            ref="safe"
                            >
                    </button>
                    <input
                        type="test"
                        autocomplete="off"
                        class="form-control"
                        id="title" 
                        :placeholder="$t('text.note.title')"
                        v-model="title"
                        ref="title" autofocus
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
                    v-model="note"
                >
                </textarea>        
            </div>
            <button
                type="submit"
                class="saveButton"
            >
                {{ $t('text.note.save') }}
            </button>
        </form>
        </ValidationObserver>
    </div>
</template>

<script>
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
            title: '',
            note: '',
            currentObject: {
                id: '',
                user_id: '',
                title: '',
                body: '',
                folder_id: localStorage.getItem('currentFolder')
            },
            id: this.$uuidKey(),
            save: false,
            errors: []
        }
    },
    methods: {
        onSubmit (error) {
            if (error === undefined) {
                if(this.notesList===null) {
                    this.notesList = []
                }
            }
            this.currentObject.user_id = this.$auth.user.sub
            this.currentObject.title = this.title
            this.currentObject.body = this.note
            const payload = {'note': this.currentObject,'userID': this.$auth.user.sub}
            this.$store.dispatch('notesModule/post', payload)
            .then(() => {
                this.save = true
                this.$router.push('/')
            })
            .catch ((err) => {
                console.log(err)
            })
        }
    },
    mounted () {
        if (this.id === null) {
            this.id = 0
        }
        this.$refs.title.focus();
        this.$store.state.currentComponent = 'Notes'
    },
    created () {
        document.getElementById('body').style.overflow = 'hidden'
    },
    beforeRouteLeave (to, from, next) {
        if(
            this.note !== '' ||
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
.saveHidden {
    position: absolute;
    right: 0.5rem;
    top: -0.15rem;
    background: transparent;
    border: none;;
}
.deleteNew {
    width: 2.8em;
}
</style>