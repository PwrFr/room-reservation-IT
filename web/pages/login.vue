<template>
  <div style="margin: 100px">
    <div v-if="$auth.loggedIn">
      {{ $auth.user }}
      <v-btn @click="$auth.logout()">Logout</v-btn>
    </div>
    <div v-else>
      <v-btn @click="login">Login</v-btn>
    </div>
  </div>
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
