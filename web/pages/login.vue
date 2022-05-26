<template>
  <v-app class="my-4 mx-4 rounded-3xl">
    <v-row align="center">
      <v-col class="text-right" style="border-right: 2px solid black">
        <h1 style="font-size: 3rem">
          Welcome to<br />
          Reservation@IT
        </h1>
      </v-col>
      <v-col v-if="$auth.loggedIn">
        <v-btn to="/" class="mx-5" x-large outlined color="success"
          ><v-icon left dark x-large> mdi-home </v-icon>Go to Home</v-btn
        >
      </v-col>
      <v-col v-else>
        <v-btn @click="login()" class="mx-5" x-large outlined color="success"
          ><v-icon left dark x-large> mdi-google </v-icon>Sign In With
          Google</v-btn
        >
      </v-col>
    </v-row>
  </v-app>
</template>

<script>
export default {
  methods: {
    login() {
      this.$auth.loginWith("google");
    },
  },
};
</script>

<style></style>
