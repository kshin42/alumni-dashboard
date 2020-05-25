<template>
  <div>
    <v-row>
    <h1>Roster</h1>
    <v-spacer></v-spacer>
    <v-btn>Import</v-btn>
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

export default {
  data: () => ({
    hello: "world",
    alumni: ""
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
    }
  }
};
</script>
