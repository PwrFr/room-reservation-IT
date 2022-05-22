<template>
  <v-app>
    <v-container grid-list-xs>
      <h1 class="text-3xl font-bold text-right">หน้าหลัก</h1>
      {{$apolloData.loading?"Loading":$apolloData.data}}


      <!-- {{item}} -->
    </v-container>
  </v-app>
</template>

<script>
// import user from '~/apollo/queries/user.gql'
import gql from 'graphql-tag';

const ALL_USER = gql`
query user{
	user{
    id
    name
  }

}`
export default {
  name: "IndexPage",
   apollo: {
    user: {
      query: ALL_USER,
      prefetch: true,
    },
  },
  data: () => ({
    // user: []
  }),


 
  
};

</script>
<style>
</style>
