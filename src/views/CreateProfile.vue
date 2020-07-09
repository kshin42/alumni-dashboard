<template>
    <div>
        <v-content>
            <v-container class="fill-height" fluid>
                <v-btn to="/" icon color="white"><v-icon dark large style="margin-right: 10px">fab fa-connectdevelop</v-icon></v-btn>
                <v-toolbar-title style="color:white;" >Telam</v-toolbar-title>
                <v-row align="center" justify="center">
                    <h1>Create Profile</h1>
                </v-row>
                <v-row align="center" justify="center">
                    <h2 class="subheader">Enter in your information</h2>
                </v-row>
                <v-row align="center" justify="center">
                    <v-form id="createForm">
                        <v-row>
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
                            <p style="color:white;margin-top:1em;">Already have an account? <v-btn x-small text to="/signIn">Sign In</v-btn></p>
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
    data: () => ({
        firstName: "",
        lastName: "",
        email: "",
        password: "",
    }),
    methods: {
        createProfile() {
            
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