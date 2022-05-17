<template>
  <v-app>
    <v-container grid-list-xs>
      <h1 class="text-3xl font-bold text-right">หน้าหลัก</h1>

      <!-- {{users}} -->
    </v-container>
  </v-app>
</template>

<script>
import gql from 'graphql-tag'
export default {
  name: "IndexPage",
  data: () => ({
    // posts: [],
  }),
  apollo:{
    
    users: gql` 
      query {  
        users{
          fullname
          _id
          fullname
		      username
          password
          createdAt

          posts{
            title
          }
        }
      }
    `
  }
  
};

</script>
<style>
</style>
