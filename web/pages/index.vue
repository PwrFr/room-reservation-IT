<template>
  <v-app>
    <v-container grid-list-xs>
      <h1 class="text-3xl font-bold text-right">หน้าหลัก</h1>

      {{user}}
    </v-container>
  </v-app>
</template>

<script>
import gql from 'graphql-tag'
export default {
  name: "IndexPage",
  data: () => ({
    user: []
  }),

  apollo: {
    user: {
      prefetch: true,
      query: ``
    }
  }
  
};

</script>
<style>
</style>
