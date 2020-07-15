<template>
<div>
    <Navigation />
    <v-data-table
    :headers="headers"
    :items="resumes"
    :single-expand="singleExpand"
    :expanded.sync="expanded"
    item-key="name"
    show-expand
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
      <td :colspan="headers.length">More info about {{ item.name }}</td>
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
        singleExpand: false,
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
        resumes: [
          {
            name: 'John Smith',
            major: 'CS',
            class: 'First Year',
            feedNum: 24,
            protein: 4.0,
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
        
    }
}
</script>

<style scoped>

</style>