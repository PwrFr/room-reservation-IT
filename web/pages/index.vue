<template>
  <v-app class="my-4 ">
  <v-form v-model="valid" class="pt-11">
    <v-container>
      <v-row  >
        <v-col
          cols="12"
          md="6"
        >
          <v-text-field
            v-model="firstname"
            
            :counter="10"
            label="Room name"
            required
             outlined
            shaped
          ></v-text-field>
        </v-col>

        
    <v-col
      cols="12"
      md="3"
    >
      <v-menu
        ref="menu"
        v-model="menu"
        :close-on-content-click="false"
        :return-value.sync="date"
        transition="scale-transition"
        offset-y
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="dates"
            label="Date"
            prepend-icon="mdi-calendar"
            readonly
             outlined
            shaped
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker
          v-model="dates"
          range
          no-title
          scrollable
        >
          <v-spacer></v-spacer>
          <v-btn
            text
            color="primary"
            @click="menu = false"
          >
            Cancel
          </v-btn>
          <v-btn
            text
            color="primary"
            @click="$refs.menu.save(dates)"
          >
            OK
          </v-btn>
        </v-date-picker>
      </v-menu>
    </v-col>


        <v-col
          cols="12"
          md="1"
        >
          <v-text-field
            v-model="capacity"
            label="Capacity"
            type="number"
            required
             outlined
            shaped
          ></v-text-field>
        </v-col>
        <v-col
          cols="12"
          md="2"
        >
          <v-btn class="mt-1" color="secondary" block x-large outlined style="border-top-left-radius: 1rem;border-top-right-radius: 1rem;">
    Filter
  </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
    <div class="rounded-3xl ">
    <v-container class="p-10 " grid-list-xs style="height: 74.1vh; overflow-y: scroll;">
      <!-- {{$apolloData.loading?"Loading":$apolloData.data}} -->
<v-row v-if="$apolloData.loading">

  <v-col cols="4"  v-for="i in 6" :key="i">
<v-skeleton-loader 
      class="m-4"
      max-width="300"
      type="article, actions"
    ></v-skeleton-loader>
     </v-col>
</v-row>
<v-row v-else-if="!$apolloData.loading">
  <v-col cols="4"  v-for="(room,i) in $apolloData.data.rooms" :key="i">
  <v-card 
    class=" p-3 rounded-lg"
    outlined
  >
    <v-list-item three-line>
      <v-list-item-content>
              <v-list-item-title class=" mb-1" style="font-size: 1.5rem;">
          {{room.name}}
        </v-list-item-title>
        <v-list-item-subtitle>Type : {{room.type}}</v-list-item-subtitle>
        <v-list-item-subtitle>Capacity : {{room.capacity}}</v-list-item-subtitle>
        <v-list-item-subtitle>Status : {{room.status}}</v-list-item-subtitle>


      </v-list-item-content>

      <v-list-item-avatar
        tile
        size="80"
        color="grey"
      ></v-list-item-avatar>
    </v-list-item>

    <v-card-actions style="justify-content: end;">
      <v-btn
        outlined
        text
        color="primary"
      >
        Book
      </v-btn>
    </v-card-actions>
  </v-card>
</v-col>
</v-row>
    </v-container>
       </div>
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
     dates: ['2019-09-10', '2019-09-20'],
     valid: false,
      firstname: '',
      lastname: '',
      nameRules: [
        v => !!v || 'Name is required',
        v => v.length <= 10 || 'Name must be less than 10 characters',
      ],
      capacity: '',
      
      menu: false,


  }),
  apollo: {
    rooms: {
      query: ALL_ROOMS,
    },
  },
   computed: {
      dateRangeText () {
        return this.dates.join(' ~ ')
      },
    },
};

</script>
<style>
.v-application--wrap {
    min-height: 95vh
}
</style>
