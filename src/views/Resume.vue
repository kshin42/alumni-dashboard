<template>
<div>
    <Navigation />
    <v-dialog v-model="dialog" max-width="600px">
      <template v-slot:activator="{ on, attrs }">
        <v-btn v-bind="attrs" v-on="on" style="margin:10px"><v-icon style="margin-right:10px" class="fillin">far fa-file-alt</v-icon>Resume Link</v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="headline">How to Link Your Resume</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              1. Upload your resume to google docs<br />
              2. Hit the share button in the top right corner<br />
              3. Select the "Get link" option<br />
              4. Below the link is a dropdown with access options. Make sure it is "Anyone with the link" and NOT "Restricted"<br />
              5. Copy generated<br />
              6. Paste link in below field and hit Save<br />
              <v-col cols="12">
                <v-text-field label="Google Doc Link*" required></v-text-field>
              </v-col>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="accent" text @click="dialog=false,saveResume()">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-row>
        <v-col cols="6">
            <v-card class="resumeframe">
                <p v-if="!resumeLink">Please Upload Resume Using Above Button</p>
                <iframe v-if="resumeLink" class="doc" :src="resumeLink"></iframe>
            </v-card>
        </v-col>
        <v-col cols="6">
            <h3 style="color:white;">Feedback Cards</h3>
            <v-divider></v-divider>
            <v-card class="feedbackcard">
                <v-card-title>Kevin Shin</v-card-title>
                <v-card-text>So I think you have a lot of things right here
                    but I think that you can be way more descriptive in your 
                    work items and make sure to always put the results of 
                    your work first.
                </v-card-text>
                <v-card-actions>
                    <v-btn text color="accent">Reply</v-btn>
                    <v-btn text color="error">Archive</v-btn>
                </v-card-actions>
            </v-card>
             <v-card class="feedbackcard">
                <v-card-title>Kevin Shin</v-card-title>
                <v-card-text>So I think you have a lot of things right here
                    but I think that you can be way more descriptive in your 
                    work items and make sure to always put the results of 
                    your work first.
                </v-card-text>
                <v-card-actions>
                    <v-btn text color="accent">Reply</v-btn>
                    <v-btn text color="error">Archive</v-btn>
                </v-card-actions>
            </v-card>
             <v-card class="feedbackcard" >
                <v-card-title>Kevin Shin</v-card-title>
                <v-card-text>So I think you have a lot of things right here
                    but I think that you can be way more descriptive in your 
                    work items and make sure to always put the results of 
                    your work first.
                </v-card-text>
                <v-card-actions>
                    <v-btn text color="accent">Reply</v-btn>
                    <v-btn text color="error">Archive</v-btn>
                </v-card-actions>
            </v-card>
             <v-card class="feedbackcard">
                <v-card-title>Kevin Shin</v-card-title>
                <v-card-text>So I think you have a lot of things right here
                    but I think that you can be way more descriptive in your 
                    work items and make sure to always put the results of 
                    your work first.
                </v-card-text>
                <v-card-actions>
                    <v-btn text color="accent">Reply</v-btn>
                    <v-btn text color="error">Archive</v-btn>
                </v-card-actions>
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
    }),
    mounted() {
        this.$store.dispatch('getResume')
        .then(response => {
            this.resumeLink = response
        })
    },
    methods: {
        saveResume() {
            this.resumeLink=""
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