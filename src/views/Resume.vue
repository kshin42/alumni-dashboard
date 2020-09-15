<template>
<div>
    <Navigation />
    <v-row align="center" justify="center">
        <v-alert v-if="errorMessage" dense color="error" icon="fas fa-exclamation-triangle" >{{errorMessage}}</v-alert>
    </v-row>
    <v-dialog v-model="dialog" max-width="600px">
      <template v-slot:activator="{ on, attrs }">
        <v-btn v-bind="attrs" v-on="on" style="margin-top:20px;margin-left:163px"><v-icon style="margin-right:10px" class="fillin">far fa-file-alt</v-icon>Link Resume</v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="headline">How to Link Your Resume</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row style="color:black">
            <div>
              1. Upload your resume to google docs<br />
              2. Hit the share button in the top right corner<br />
              3. Select the "Get link" option<br />
              4. Below the link is a dropdown with access options. Make sure it is <strong>"Anyone with the link"</strong> and NOT "Restricted"<br />
              5. To the right of the access dropdown in a permissions dropdown. Make sure it is <b>"Commenter"</b> and NOT "Viewer" or "Editor"<br />
              6. Copy generated link<br />
              7. Paste link in below field and hit Save<br />
            </div>
              <v-col cols="12">
                <v-text-field label="Google Doc Link*" v-model="uploadLink" required></v-text-field>
              </v-col>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
          <v-alert v-if="saveErrorMessage" dense color="error" icon="fas fa-exclamation-triangle" >{{saveErrorMessage}}</v-alert>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="accent" text @click="saveResume()">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-row align="center" justify="center">
        <v-col cols="10">
            <v-card class="resumeframe">
                <p v-if="!resumeLink">Please Upload Resume Using Above Button</p>
                <iframe v-if="resumeLink" class="doc" :src="resumeLink"></iframe>
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
    },
    data:() => ({
        resumeLink: "",
        dialog: "",
        uploadLink: "",
        errorMessage: "",
        saveErrorMessage: ""
    }),
    mounted() {
        this.getResume()
    },
    methods: {
        async saveResume() {
            const response = await this.$store.dispatch('uploadResume',{
                email: this.$store.getters.getCurrentUser,
                resumeLink: this.uploadLink
            })
            if (response.status == 200) {
                this.resumeLink = this.uploadLink
                this.dialog = false
                this.saveErrorMessage = ""
            } else {
                this.saveErrorMessage = "Failed to Upload Resume. Please contact an administrator"
            }
        },
        async getResume() {
            const response = await this.$store.dispatch('getResume')
            if (response.status == 200) {
                this.resumeLink = response.data
            } else {
                this.errorMessage = "Failed to load resume. Please upload a new link."
            }
        }
    }
}
</script>

<style scoped>
.doc {
    width: 100%;
    height: 100%;
}
.feedbackcard {
    margin-top: 10px;
}

.resumeframe {
    margin-top: 10px;
    height: 800px;
}
</style>