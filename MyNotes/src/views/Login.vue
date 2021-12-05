<template>
    <div>
        <h1 class="header">My<span style="color:rgb(215, 0, 0);">N</span>otes</h1>
        <hr class="whiteLine">
        <br>
        <h2 class="whiteText">
            Welcome!
        </h2>
        <p class="whiteText">
            You can simply log in below.
        </p>
        <!-- <div class="form"> <br><br>
            <input
                type="text"
                placeholder="Username"
                v-model="username"
            >
            <br><br>
            <input
                type="password"
                placeholder="Password"
                v-model="password"
                v-on:keyup.enter="logIn()"
            >
            <br><br>
            <button
                type="button"
                v-on:click="logIn()"
                id="login"
            >
                Log In
            </button>
        </div> -->

        <span v-if="!$auth.loading" class="is-log-in-buttons">
          <button
            v-if="!$auth.isAuthenticated"
            @click="login"
            class="button is-log-in-button is-outlined is-dark mr-3"
          >
            Log in
          </button>
          <button
            v-if="!$auth.isAuthenticated"
            @click="register"
            id="login"
          >
            Register
          </button>
          <button
            v-if="$auth.isAuthenticated"
            @click="logout"
            id="login"
          >
            Log out
          </button>
        </span>
               
        <br><br><br>
        <router-link to="/register" id="register">Or register new Account</router-link>
        <br><br><br><br><br><br>
        <!-- <h2 class="whiteText">This App is currently under construction. <br>
            It will be available again soon.</h2>
        <img src="../../public/img/settings.png" alt="construction" class="settings"> -->
    </div>
</template>

<script>
export default {
    name: 'App',
    data() {
        return {
            username: '',
            password: '',
            correct: false,
            isActive: false
        }
    },
    computed: {
        authenticated () {
            if (this.$auth.isAuthenticated) {
                // return this.$auth.user.sub
                return true
            } 
            return false
        }
    },
    watch: {
        authenticated (val) {
            if (val) this.$router.push('/')
        }
    },
    methods: {
    login() {
        this.$auth.loginWithRedirect();
    },
    register() {
        this.$auth.loginWithRedirect({
        screen_hint: 'signup'
        });
    },
        // logIn () {
        //     if (this.username !== '' || this.password !== '') {
        //         axios.get(`${this.$store.state.localhost}users`, {
        //         headers: {
        //             'username': this.username,
        //             'password': ""+this.password.hashCode()
        //         }
        //         })
        //         .then((response) => {
        //             if (response.data.username == this.username) {
        //                 this.$store.state.userID = response.data.id
        //                 alert(response.data.id)
        //                 localStorage.setItem('userID',(response.data.id))
        //                 this.$router.push('/') 
        //             } else {
        //                 alert("Wrong Password or Username")
        //             }
        //         })
        //         .error((err) => {
        //             console.log(err)
        //         })
        //     } else {
        //         alert("Wrong Password or Username")
        //     }
        // }
    }
}
</script>

<style scoped>
    /* .whiteText {
        margin-top: 7.5rem;
        color: lightgray
    }
    .settings {
        margin-top: 3rem;
        width: 15rem;
    } */
    body {
        margin:1rem;
    }
    #googleLogin {
        display: flex;
        justify-content: center;
        margin: 0 auto;
        align-items: center;
        margin-top: 4rem;
        width: 12rem;
        height: 2.6rem;
        border: none;
    } 
    #googleImage {
        width: 12.3rem;
    }
    .invisible {
        opacity: 0;
        height: 2rem;
    }
    .form {
        display: block;
        text-align: center;

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
    @media (min-width: 1000px) { 
    }
</style>