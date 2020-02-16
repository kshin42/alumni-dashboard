<template>
  <v-app>
    <v-navigation-drawer
      clipped
      permanent
      app
      overflow
    >
    <v-list
        dense
        nav
      >
        <v-list-item
          v-for="item in items"
          :key="item.title"
          link
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar :clipped-left=true color="#730000" app>
      <v-toolbar-title>Alpha Sigma Phi - Alumni Dashboard </v-toolbar-title>
    </v-app-bar>

    <v-content>
      <v-container fluid>
        <v-row
          align="center"
          justify="center"
        >
          <v-col cols="10">
            <v-card>

            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-content>

    <v-footer app>
      <span class="px-4">&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>


<script>
export default {
 data: () => ({
   items: [
    { title: 'Dashboard', icon: 'mdi-view-dashboard' },
    { title: 'Roster', icon: 'mdi-account-box-multiple-outline' },
    { title: 'About', icon: 'mdi-help-box' },
   ]
 }),
 created () {
      this.$vuetify.theme.dark = true
 },
}
</script>
