<template>
  <div>
    <v-app-bar :clipped-left=true color="tertiary" dark flat app>
      <v-app-bar-nav-icon @click="drawer=true"><v-icon color="secondary">fas fa-bars</v-icon></v-app-bar-nav-icon>
      <v-toolbar-title>Telam</v-toolbar-title>
      <v-spacer></v-spacer>
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
  items: [
    { title: 'Dashboard (WIP)', icon: 'fas fa-th-large', path: '/'},
    { title: 'Resume', icon: 'far fa-file-alt', path: '/resume'},
    { title: 'Roster', icon: 'fas fa-clipboard-list', path: '/roster' }
   ]
 }),
 methods: {
   async logout() {
      await this.$store.dispatch('destroyToken')
    }
 }
}
</script>

<style scoped>

</style>
