<template>
  <div>
    <Navigation />
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
    this.$store.dispatch('getRoster')
      .then(response => {
          this.alumni = response
      })
  },
  methods: {
    openMailDialog(recipient) {
      window.location.href = `mailto:${recipient}?subject=Subject Placeholder`;
    },
    async logout() {
      await this.$store.dispatch('destroyToken')
    }
  }
};
</script>
