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
                    <iframe class="doc" :src=item.link></iframe>
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
        resumes: [
          {
            name: 'John Smith',
            major: 'CS',
            class: 'First Year',
            feedNum: 24,
            link: 'https://docs.google.com/document/d/1868X2oXFgknj6-Al1y8ldhvJt4qpu38ZVyAMTEjO674/edit?usp=sharing',

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
    height: 900px;
    width: 100%;
}
</style>