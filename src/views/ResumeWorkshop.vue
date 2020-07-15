<template>
<div>
    <Navigation />
    <v-data-table
    :headers="headers"
    :items="resumes"
    :expanded="expanded"
    item-key="name"
    show-expand
    @click:row="clicked"
    class="elevation-1"
    style="margin:10px;"
  >
    <template v-slot:top>
      <v-toolbar flat>
        <v-toolbar-title>Resumes</v-toolbar-title>
        <v-spacer></v-spacer>
      </v-toolbar>
    </template>
    <template v-slot:expanded-item="{ headers, item }">
        <td :colspan="headers.length">{{ item.name }}
            <v-row>
                <v-col cols="6">
                    <v-card>
                        <iframe class="doc" src=""></iframe>
                    </v-card>
                </v-col>
                <v-col cols="6">
                    <h3>Feedback Card</h3>
                    <v-divider style="margin-bottom:20px"></v-divider>
                    <v-card class="feedbackcard">
                        <v-card-title>Please Enter Your Feedback</v-card-title>
                         <v-textarea style="margin:10px"
                        :auto-grow="true">
                      </v-textarea>
                        <v-card-actions>
                            <v-btn text color="accent">Submit</v-btn>
                        </v-card-actions>
                    </v-card>
                    <h3 style="margin-top:50px">Prior Feedback Given</h3>
                    <v-divider style="margin-bottom:20px"></v-divider>
                     <v-data-table
                        :headers="subheaders"
                        :items="feedback"
                        :items-per-page="3"
                        item-key="feedback"
                        class="elevation-1"
                        style="margin:10px;" dark >
                    </v-data-table>
                </v-col>
            </v-row>
        </td>
    </template>
  </v-data-table>
</div>
</template>

<script>
import Navigation from '../components/Navigation.vue'
export default {
    components: {
        Navigation,
    },
    data:() => ({
        expanded: [],
        headers: [
          {
            text: 'Name',
            align: 'start',
            sortable: false,
            value: 'name',
          },
          { text: 'Major', sortable: true, value: 'major' },
          { text: 'Class', sortable: true, value: 'class' },
          { text: '# Feedback Received', sortable: true, value: 'feedNum' },
        ],
        subheaders: [
            {
                text: 'Feedback',
                align: 'start',
                value: 'feedback',
            }
        ],
        feedback: [
            {feedback: "this was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guess"},
            {feedback: "this wa alright I guess"},
            {feedback: "this was alrigt I guess"},
            {feedback: "this was alright I guehis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guess"},
            {feedback: "this walright I guess"},
            {feedback: "this was alright I ess"},
            {feedback: "this was alright I this was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guessthis was alright I guess"},
            {feedback: "this was aght I guess"},
            {feedback: "this was alright I gs"},
        ],
        resumes: [
          {
            name: 'John Smith',
            major: 'CS',
            class: 'First Year',
            feedNum: 24,
            iron: '1%',

          },
          {
            name: 'Adam Knight',
            major: 'CE',
            class: 'Second Year',
            feedNum: 37,
            protein: 4.3,
            iron: '1%',
          },
          {
            name: 'Julio Jones',
            major: 'ME',
            class: 'Third Year',
            feedNum: 23,
            protein: 6.0,
            iron: '7%',
          }
        ],
    }),
    mounted() {
        // this.$store.dispatch('getResumes')
        // .then(response => {
        //     this.resumeLink = response
        // })
    },
    methods: {
        clicked(value) {
            if (this.expanded.includes(value)) {
                var i = this.expanded.indexOf(value)
                this.expanded.splice(i, 1)
            } else {
                this.expanded.push(value)
            }
        }
    }
}
</script>

<style scoped>
.doc {
    margin-top: 10px;
    height: 800px;
    width: 100%;
}
</style>