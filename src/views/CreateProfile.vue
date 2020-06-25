<template>
    <div class="bg">
        <v-content>
            <v-container class="fill-height" fluid>
                <v-row align="center" justify="center">
                    <h1>Create Profile</h1>
                </v-row>
                <v-row align="center" justify="center">
                    <h2 class="subheader">Enter in your information</h2>
                </v-row>
                <v-row align="center" justify="center">
                    <v-form>
                        <v-row>
                            <v-col cols="12">
                                <v-text-field label="Solo" placeholder="First Name" solo></v-text-field>
                            </v-col>
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
    }
}
</script>

<style scoped>
.text-input {
    color: white;
}
</style>