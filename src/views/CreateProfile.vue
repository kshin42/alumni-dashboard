<template>
    <div>
        <v-content>
            <v-container class="fill-height" fluid>
                <v-btn to="/" icon color="white"><v-icon dark large style="margin-right: 10px">fab fa-connectdevelop</v-icon></v-btn>
                <v-toolbar-title style="color:white;" >Telam</v-toolbar-title>
                <v-row align="center" justify="center">
                    <v-alert v-if="errorMessage" dense color="error" icon="fas fa-exclamation-triangle" >{{errorMessage}}</v-alert>
                </v-row>
                <v-row align="center" justify="center">
                    <h1 style="padding-top:100px">Create Profile</h1>
                </v-row>
                <v-row align="center" justify="center">
                    <h2 class="subheader">Enter in your information</h2>
                </v-row>
                <v-row align="center" justify="center">
                    <v-form id="createForm">
                        <v-row>
                            <v-col cols="12">
                                <v-text-field v-model="orgCode" dark label="Organization Code" outlined ></v-text-field>
                            </v-col>
                            <v-col cols="6">
                                <v-text-field v-model="firstName" dark label="First Name" outlined ></v-text-field>
                            </v-col>
                            <v-col cols="6">
                                <v-text-field v-model="lastName" dark label="Last Name" outlined ></v-text-field>
                            </v-col>
                            <v-col cols="6">
                                <v-text-field v-model="email" dark label="Email" outlined ></v-text-field>
                            </v-col>
                            <v-col cols="6">
                                <v-text-field v-model="password" dark label="Password" outlined ></v-text-field>
                            </v-col>
                        </v-row>
                        <v-row align="center" justify="center">
                            <v-btn @click="createProfile()">Sign Up</v-btn>
                        </v-row>
                        <v-row align="center" justify="center">
                            <p style="color:white;margin-top:1em;">Already have an account? <v-btn x-small text to="/login">Sign In</v-btn></p>
                        </v-row>
                    </v-form>
                </v-row>
            </v-container>
        </v-content>
    </div>
</template>

<script>
import AuthService from '../services/AuthService'

export default {
    components: {

    },
    data: () => ({
        orgCode: "",
        firstName: "",
        lastName: "",
        email: "",
        password: "",
        errorMessage: "",
    }),
    methods: {
        async createProfile() {
            const userData = {
                orgCode: this.orgCode,
                firstName: this.firstName,
                lastName: this.lastName,
                email: this.email,
                password: this.password
            }

            await AuthService.createMember(userData)
            .then(response => {
                this.$store.dispatch('retrieveToken', {
                    email: this.email,
                    password: this.password,
                })
                .then(response => {
                    this.$router.push('/resume')
                })
                .catch(err => {
                    console.log("failed to login")
                })
            })
            .catch(error => {
                this.errorMessage = error.response.data
            })
        }
    }
}
</script>

<style scoped>
#createForm {
    margin-top: 25px;
}

button {
    font-size: .5em;
}
</style>