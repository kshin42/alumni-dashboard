<template>
    <div>
        <v-parallax
            id="launchimage"
            dark
            src="../assets/networking.png"
            height="700"
        >
        <h1>Build Your Network</h1>
        <h2 class="subheading">Create your organization</h2>
        <v-btn id="startbutton" large color="primary" to="/userImport">Start</v-btn>
    </v-parallax>
    </div>
</template>

<script>
export default {
 data: () => ({}),
}
</script>

<style>
h1 {
  font-size: 80px;
  margin-left: 80px;
  padding-top: 200px;
}

h2 {
  margin-left: 90px;
}

#startbutton {
  width: 100px !important;
  margin-top: 2em;
  margin-left: 90px;
}
</style>