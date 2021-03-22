<template>
    <div id="app">
        <router-link to="/">
            <img src="../assets/arrow.png" alt="back arrow" class="arrow">
        </router-link>
        <br><br>
        <ValidationObserver v-slot="{ handleSubmit }">
        <form @submit.prevent="handleSubmit(onSubmit)">
            <button to="/" class="saveHidden">
                <img src="../assets/haken.png" alt="delete icon" class="deleteNew" type="submit">
            </button>
            <div class="form-group">
                <ValidationProvider name="email" rules="required" v-slot="{ errors }">
                    <input type="test" autocomplete="off" class="form-control" id="title"  placeholder="Title" v-model="title" ref="title" autofocus>        
                    <span class="errorMessage">{{ errors[0] }}</span>
                </ValidationProvider>
            </div>
            <div class="form-group">
                <textarea type="test" class="form-control" id="note"  placeholder="Note" v-model="note"> </textarea>        
            </div>
            <button type="submit" class="saveButton">Save</button>
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
                title: '',
                note: '',
                id: ""
            },
            notesList: JSON.parse(localStorage.getItem('notes')),
            id: this.$uuidKey()
        }
    },
    methods: {
        onSubmit () {
            if(this.notesList===null) {
                this.notesList = []
            }
            this.currentObject.title = this.title
            this.currentObject.note = this.note
            this.currentObject.id = this.id
            this.notesList.push(this.currentObject)
            localStorage.setItem('id', this.id)
            localStorage.setItem('notes',JSON.stringify(this.notesList))
            this.$router.push('/')
        }
    },
    mounted () {
        if (this.id === null) {
            this.id = 0
        }
        this.$refs.title.focus();
    }
}
</script>

<style scoped>
.saveHidden {
    position: absolute;
    right: 0.5rem;
    top: 0.6rem;
    background: transparent;
    border: none;;
}
.deleteNew {
    width: 2.2rem;
    background: none;
}
</style>