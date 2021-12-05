<template>
    <div>
         <h1 class="header">My<span style="color:rgb(215, 0, 0);">N</span>otes</h1>
        <hr class="whiteLine">
        
        <h2 class="whiteText marginTop">
            Welcome!
        </h2>
        <p class="whiteText">
            You can register a new account here.
        </p>
         <ValidationObserver v-slot="{ invalid }">
            <form class="form" @submit.prevent="onSubmit"><br>

                <ValidationProvider name="E-mail" rules="required" v-slot="{ errors }">
                    <input v-model="username" type="text" placeholder="Username"><br>
                    <span class="errorMessage">{{ errors[0] }}</span>
                    <span v-if="exists" class="errorMessage">
                        Username exists already.
                    </span>
                </ValidationProvider><br>

                <ValidationProvider name="Password" rules="required|min:10" v-slot="{ errors }">
                    <input v-model="password" type="password" placeholder="Password"><br>
                    <span class="errorMessage">{{ errors[0] }}</span>
                </ValidationProvider><br>

                <ValidationProvider name="Password" rules="required">
                    <input v-model="repeatedPassword" type="password" placeholder="Repeat Password"><br>
                    <span v-if="passwordError" class="errorMessage">
                        The passwords are not equal.
                    </span>
                </ValidationProvider>
                <br>

                <button
                    v-if="invalid || passwordError"
                    id="login"
                    :disabled="true"
                    class="disabled"
                >
                    Register
                </button>
                <button
                    v-else
                    type="submit"
                    :disabled="false"
                    id="login"
                >
                    Register
                </button>
            </form>
        </ValidationObserver>
        <br><br><br>
        <router-link to="/logIn" id="register">Log In instead</router-link>
        

    </div>
</template>

<script>
import axios from 'axios'
import { ValidationProvider, ValidationObserver, extend } from 'vee-validate';
import { required, min } from 'vee-validate/dist/rules';

extend('required', {
  ...required,
  message: 'This field is required'
});

extend('min', {
  ...min,
  message: 'Password must be at least 10 characters long'
});

String.prototype.hashCode = function() {
    var hash = 0;
    if (this.length == 0) {
        return hash;
    }
    for (var i = 0; i < this.length; i++) {
        var char = this.charCodeAt(i);
        hash = ((hash<<5)-hash)+char;
        hash = hash & hash; // Convert to 32bit integer
    }
    return hash;
}

export default {
    data () {
        return {
            username: '',
            password: '',
            repeatedPassword: '',
            passwordError: false,
            exists: false
        }
    },
    components: {
        ValidationProvider,
        ValidationObserver
    },
    methods: {
        onSubmit () {
            axios.post(`${this.$store.state.localhost}users`, {
                'username': this.username,
                'password': ""+this.password.hashCode()
            })
            .then((response) => {
                if(response.data == "exists") {
                    this.exists = true
                } else {
                    this.exists = false
                    this.$store.state.userID = response.data
                    localStorage.setItem('userID', response.data); 
                    this.$router.push('/')
                }
            })
        }
    },
    watch: {
        repeatedPassword (val) {
            if (val === this.password) {
                this.passwordError = false
            } else {
                this.passwordError = true
            }
        }
    }
}
</script>

<style scoped>
    .invisible {
        opacity: 0;
    }
    .form {
        display: block;
        text-align: center;
    }
    #register {
        width: 6rem;
        margin-top: 2.5rem;
    }
    input {
        border-radius: 5px;
        height: 2.5rem;
         background-color: transparent;
        border: none;
        box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                    1px 1px 20px 2px rgb(0, 0, 0);
        color: lightgray;
        font-weight: 400;
        font-size: larger;
        padding-left: 1rem;
        padding-right: -1rem;
    }
    .marginTop {
        margin-top: 2rem !important;
    }
</style>