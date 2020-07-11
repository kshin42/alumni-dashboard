<template>
    <div>
        <v-content>     
            <v-container class="fill-height" fluid>
                <v-btn to="/" icon color="white"><v-icon dark large style="margin-right: 10px">fab fa-connectdevelop</v-icon></v-btn>
                <v-toolbar-title style="color:white;" >Telam</v-toolbar-title>
                <v-row align="center" justify="center">
                    <h1>Login</h1>
                </v-row>
                <v-row align="center" justify="center">
                    <h2 class="subheader">Enter in your login information</h2>
                </v-row>
                <v-row align="center" justify="center">
                    <v-form id="createForm">
                        <v-row>
                            <v-col cols="6">
                                <v-text-field v-model="email" dark label="Email" outlined ></v-text-field>
                            </v-col>
                            <v-col cols="6">
                                <v-text-field v-model="password" dark label="Password" outlined ></v-text-field>
                            </v-col>
                        </v-row>
                        <v-row align="center" justify="center">
                            <v-btn @click="login()" >Login</v-btn>
                        </v-row>
                        <v-row align="center" justify="center">
                            <p style="color:white;margin-top:1em;">Don't have an account? <v-btn x-small text to="/createProfile">Sign Up</v-btn></p>
                        </v-row>
                    </v-form>
                </v-row>
            </v-container>
        </v-content>
    </div>
</template>

<script>
import Navigation from '../components/Navigation'
export default {
    components: {
        Navigation,
    },
    computed: {
        loggedIn() {
            return this.$store.getters.loggedIn
        }
    },
    data: () => ({
        email: '',
        password: '',
    }),
    methods: {
        async login() {
            this.$store.dispatch('retrieveToken', {
                email: this.email,
                password: this.password,
            })
              .then(response => {
                  this.$router.push('/roster')
              })
              .catch(err => {
                  console.log("failed to login")
              })
        }
    }
}
</script>

<style scoped>
#createForm {
    margin-top: 25px;
}
</style>