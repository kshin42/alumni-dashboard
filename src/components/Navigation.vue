<template>
  <div>
    <v-app-bar :clipped-left=true color="tertiary" dark flat app>
      <v-app-bar-nav-icon @click="drawer=true"><v-icon color="secondary">fas fa-bars</v-icon></v-app-bar-nav-icon>
      <v-toolbar-title>Telam</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-dialog v-model="feedbacktoggle" max-width="600px">
        <template v-slot:activator="{ on, attrs }">
          <v-btn v-bind="attrs" v-on="on" color="primary" ><v-icon>fas fa-comment-dots</v-icon></v-btn>
        </template>
        <v-card>
          <v-card-title>
            <span class="headline">We'd love your feedback!</span>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-row style="color:black">
                <v-col cols="12">
                  <v-textarea v-model="feedbackcomment" solo required></v-textarea>
                </v-col>
              </v-row>
            </v-container>
            <v-alert  v-if="successMessage" dense type="success" outlined text dismissible icon="far fa-check-circle">Thanks for your feedback!</v-alert>
            <v-alert  v-if="errorMessage" dense dismissible color="error" icon="fas fa-exclamation-triangle" >{{errorMessage}}</v-alert>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="accent" text @click="submitFeedback()">Submit</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-btn @click="logout()" light color="white">Log out</v-btn>
    </v-app-bar>

    <v-navigation-drawer
      color="white"
      v-model="drawer"
      absolute
      temporary
    >
      <v-list nav>
        <v-list-item
          v-for="item in items"
          :key="item.title"
          :to="item.path"
        >
          <v-list-item-icon>
            <v-icon color="secondary">{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>


<script>
export default {
 data: () => ({
  drawer: false,
  dialog: false,
  feedbacktoggle: false,
  feedbackcomment: "",
  errorMessage: "",
  successMessage: false,
  items: [
    { title: 'Dashboard (WIP)', icon: 'fas fa-th-large', path: '/'},
    { title: 'My Resume', icon: 'far fa-file-alt', path: '/resume'},
    { title: 'Resume Workshop', icon: 'fas fa-screwdriver', path: '/resumeWorkshop'},
    { title: 'Roster', icon: 'fas fa-clipboard-list', path: '/roster' }
   ]
 }),
 methods: {
  async logout() {
    await this.$store.dispatch('destroyToken')
  },
  async submitFeedback() {
    const response = await this.$store.dispatch('submitFeedback',{
      feedbackcomment: this.feedbackcomment,
    })
    if (response.status == 200) {
        this.successMessage = true
    } else {
        this.errorMessage = "Sorry failed to submit feedback. Please reach out to an administrator."
    }
  }
 }
}
</script>

<style scoped>
button {
  margin: 1em;
}
</style>
