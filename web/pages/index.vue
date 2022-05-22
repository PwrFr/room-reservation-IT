<template>
  <v-app>
    <v-container grid-list-xs>
      <h1 class="text-3xl font-bold text-right">หน้าหลัก</h1>
      {{$apolloData.loading?"Loading":$apolloData.data}}


    </v-container>
  </v-app>
</template>

<script>
// import user from '~/apollo/queries/user.gql'
import gql from 'graphql-tag';

const ALL_ROOMS = gql`
query rooms{
  rooms{
    id
    name
    status
    capacity
    type
  }
}
`
export default {
  name: "IndexPage",
   
  data: () => ({
  }),
  apollo: {
    rooms: {
      query: ALL_ROOMS,
    },
  },
};

</script>
<style>
</style>
