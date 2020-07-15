<template>
<div>
    <Navigation />
    <v-row>
        <v-col cols="4">
            <v-btn><v-icon style="margin-right:10px" class="fillin">far fa-file-alt</v-icon>Upload Resume</v-btn>
            <v-card style="margin-top: 30px;">
                This is my resume and its contents are below
            </v-card>
        </v-col>
        <v-col cols="7">
            <v-card>
                <v-card-title>Kevin Shin</v-card-title>
                <v-card-text>So I think you have a lot of things right here
                    but I think that you can be way more descriptive in your 
                    work items and make sure to always put the results of 
                    your work first.
                </v-card-text>
            </v-card>
        </v-col>
    </v-row>
</div>
</template>

<script>
import Navigation from '../components/Navigation.vue'
export default {
    components: {
        Navigation,
    }
}
</script>

<style scoped>
</style>