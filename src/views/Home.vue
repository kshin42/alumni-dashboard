<template>
  <div class="bg">
    <v-content>
    <v-container class="fill-height" fluid>
      <v-row align="center" justify="center">
        <v-icon dark x-large class="homeLogo">fab fa-connectdevelop</v-icon>
        <h1>Telam</h1>
      </v-row>
      <v-row align="center" justify="center">
        <h2 class="subheader">Networking Made Easy</h2>
      </v-row>
      <v-row align="center" justify="center">
        <v-btn class="homeBtn" to="/signIn">Sign In</v-btn>
        <v-btn class="homeBtn" color="primary" to="/createProfile">Sign Up</v-btn>
      </v-row>
    </v-container>
    </v-content>
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
  padding-top: 200px;
  color: white;
}

h2 {
  color: white;
  margin-bottom: 1em;
}

.homeLogo {
  padding-top: 200px;
  margin-right: 15px;
}

.homeBtn {
  margin: 2em;
}

.header h1 {
  font-size: 40px;
}

.bg {
  background: #4281A4;
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.container {
  width: 100%;
  height: 100%;
}

</style>