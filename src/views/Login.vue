<template>
    <div>
        <h1 class="header">My<span style="color:rgb(215, 0, 0);">N</span>otes</h1>
        <hr class="whiteLine">
        <br>
        <h2 class="whiteText">
            Welcome!
        </h2>
        <p class="whiteText">
            You can simply sign in with google by using the Button below.
        </p>
        <GoogleLogin id="googleLogin" :params="params" :onSuccess="onSuccess" :onFailure="onFailure"><img id="googleImage" src="../assets/google.png" alt=""></GoogleLogin>
        <br><br><br><br><br><br><br><br>
        <!--<GoogleLogin :params="params" :logoutButton=true>Logout</GoogleLogin>-->        
    </div>
</template>

<script>
import GoogleLogin from 'vue-google-login';

export default {
        name: 'App',
        data() {
            return {
                // client_id is the only required property but you can add several more params, full list down bellow on the Auth api section
                params: {
                    client_id: "780824923030-217pr3nsqqg6rs8lt488q1oigb99tci9.apps.googleusercontent.com"
                },
                // only needed if you want to render the button with the google ui
                renderParams: {
                    width: 250,
                    height: 50,
                    longtitle: true
                },
                email: '',
                password: '',
                correct: false
            }
        },
        components: {
            GoogleLogin
        },
        created () {
            if (localStorage.userID) {
                this.$router.push('/')
            }
        },
        methods: {
        onSuccess(googleUser) {
            var user = {
                id: googleUser.getBasicProfile().getId(),
            }
            localStorage.setItem('userID', user.id)
            this.$store.state.userID = user.id
            this.$router.push('/') 
        },
        onFailure () {
            alert("Login not possible. Please contact us, if this happens repeatedly.")
        }
    }
    }
</script>

<style scoped>
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
    #login {
        width: 6rem;
        height: 3rem;
        border-radius: 1rem;
        background-color: transparent;
        border: none;
        box-shadow: -1px -1px 4px 0px rgb(133, 133, 133),
                    1px 1px 5px 2px black;
        color: lightgray;
        font-weight: 960;
        font-size: larger;
        margin-top: 1rem;
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
    .whiteText {
        color: lightgray;
        margin: 1rem;
    }
    @media (min-width: 1000px) { 
    }
</style>