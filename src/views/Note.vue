<template>
    <div id="app">
        <div>
            <router-link to="/">
                    <img src="../assets/arrow.png" alt="back arrow" class="arrow">
            </router-link>
        </div>
        <div>
            <router-link to="/">
                    <img src="../assets/trash.png" alt="delete icon" class="delete" @click="deleteNote">
            </router-link>
        </div>
        <br><br>
        <ValidationObserver v-slot="{ handleSubmit }">
        <form @submit.prevent="handleSubmit(onSubmit)">
            <div class="form-group">
                <ValidationProvider name="email" rules="required" v-slot="{ errors }">
                    <input type="test" class="form-control" id="title" autocomplete="off" placeholder="Title" value="title" v-model="title">        
                <span class="errorMessage">{{ errors[0] }}</span>
                </ValidationProvider>
            </div>
            <div class="form-group">
                <textarea type="test" class="form-control" id="note"  placeholder="Note" value="note" v-model="note" autofocus> 
                </textarea>        
            </div>
            <button class="saveButton" type="submit">Save</button>
        </form>
        </ValidationObserver>
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
            id: this.$store.state.id,
            notes: JSON.parse(localStorage.getItem('notes')),
            title: '',
            note: '',
            currentObject: {title: '', note: '', id: ""},
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
            router.push('/')
        },
        deleteNote () {
            for (let i in this.notes) {
                if (this.notes[i].id === this.id) {
                    this.notes.splice(i, 1)
                }
            }
            localStorage.setItem('notes', JSON.stringify(this.notes))
        }
    },
    mounted () {
        for (let i = 0; i < this.notes.length; i++) {
            if (this.notes[i].id === this.id) {
                this.title = this.notes[i].title
                this.note = this.notes[i].note
            }
        }
        this.$refs.title.focus();
    }
}
</script>