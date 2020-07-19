<template>
<div>
    <Navigation />
     <v-alert v-if="errorMessage" dense color="error" icon="fas fa-exclamation-triangle" >{{errorMessage}}</v-alert>
    <v-data-table
    :headers="headers"
    :items="resumes"
    :expanded="expanded"
    item-key="Name"
    show-expand
    @click:row="clicked"
    class="elevation-1"
    style="margin:20px;"
  >
    <template v-slot:top>
      <v-toolbar flat>
        <v-toolbar-title>Resumes</v-toolbar-title>
        <v-spacer></v-spacer>
      </v-toolbar>
    </template>
    <template v-slot:expanded-item="{ headers, item }">
        <td :colspan="headers.length">
            <v-row align="center" justify="center">
                <v-col cols="10">
                    <iframe class="doc" :src=item.Link></iframe>
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
        errorMessage: "",
        expanded: [],
        headers: [
          {
            text: 'Name',
            align: 'start',
            sortable: false,
            value: 'Name',
          },
          { text: 'Major', sortable: true, value: 'major' },
          { text: 'Year', sortable: true, value: 'year' },
          { text: '# Feedback Received', sortable: true, value: 'feedNum' },
        ],
        subheaders: [
            {
                text: 'Feedback',
                align: 'start',
                value: 'feedback',
            }
        ],
        resumes: [],
    }),
    mounted() {
        this.getResumes();
    },
    methods: {
        clicked(value) {
            if (this.expanded.includes(value)) {
                var i = this.expanded.indexOf(value)
                this.expanded.splice(i, 1)
            } else {
                this.expanded.push(value)
            }
        },
        async getResumes() {
            const response = await this.$store.dispatch('getResumes')
            if (response.status == 200) {
                console.log(response.data)
                this.resumes = response.data
            } else {
                this.errorMessage = "Failed to load resumes. Please contact an administrator."
            }
        }
    }
}
</script>

<style scoped>
.doc {
    margin-top: 10px;
    height: 900px;
    width: 100%;
}
</style>