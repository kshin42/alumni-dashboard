<template>
  <div>
    <v-spacer></v-spacer>
    <v-btn @click="logout()" color="white">Log out</v-btn>
    <v-row>
    <h1>Roster</h1>
    <v-spacer></v-spacer>
    <v-btn color="primary" to="/userImport">Import Users</v-btn>
  </v-row>
    <v-container>
      <v-row>
        <v-col cols="auto" v-for="(test, index) in alumni" v-bind:key="index">
          <v-card hover>
            <v-card-title>{{ test.firstName }} {{ test.lastName }}</v-card-title>
            <v-card-subtitle>{{ test.email }}</v-card-subtitle>
            <v-card-actions>
              <v-btn color="red" text @click="openMailDialog()">
                <v-icon>mdi-email</v-icon>
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import RosterService from "../services/RosterService";
import Navigation from '../components/Navigation';

export default {
  components: {
    Navigation
  },
  data: () => ({
    alumni: "",
    importDialogOpen: false,
  }),
  mounted() {
    this.getAlumni();
  },
  methods: {
    async getAlumni() {
      this.alumni = await RosterService.getRoster();
      
    },
    openMailDialog(recipient) {
      window.location.href = `mailto:${recipient}?subject=Subject Placeholder`;
    },
    async logout() {
      await this.$store.dispatch('destroyToken')
        .then(response => {
            this.$router.push('/signIn')
        })
    }
  }
};
</script>
