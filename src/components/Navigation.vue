<template>
  <v-app>
    <v-app-bar :clipped-left=true color="secondary" dark app class="hidden-sm-and-down">
      <v-icon style="margin-right: 10px">fab fa-connectdevelop</v-icon>
      <v-toolbar-title>Telam</v-toolbar-title>
    </v-app-bar>

    <v-navigation-drawer
      color="secondary"
      permanent
      dark
      expand-on-hover
      absolute
      class="hidden-sm-and-up"
    >
        <v-list-item
          v-for="item in items"
          :key="item.title"
          :to="item.path"
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
    </v-navigation-drawer>
  </v-app>
</template>


<script>
export default {
 data: () => ({
  dialog: false,
  items: [
    { title: 'Dashboard', icon: 'fas fa-th-large', path: '/'},
    { title: 'Portal Management', icon: 'fas fa-lock', path: '/manage'},
    { title: 'Roster', icon: 'fas fa-clipboard-list', path: '/roster' },
    { title: 'User Import', icon: 'fas fa-user-plus', path: '/userImport' },
   ]
 })

}
</script>
