<template>
<v-row no-gutters>
  <v-col >
  <v-app class="m-4 rounded-3xl">
  <v-form class="pt-11">
    <v-container>
      <v-row  justify="center">
        <v-col
          cols="12"
          md="6"
        >
          <v-text-field
            v-model="name"
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
        :return-value.sync="dates"
        transition="scale-transition"
        offset-y
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
           
            label="Date"
            v-model="dateRangeText"
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
    
<v-row v-if="$apolloData.loading">

  <v-col cols="4"  v-for="i in 6" :key="i">
<v-skeleton-loader 
      class="m-4"
      max-width="300"
      type="article, actions"
    ></v-skeleton-loader>
     </v-col>
</v-row>

<v-row v-else-if="$apolloData.data.rooms == ''">
No Rooms
</v-row>


<v-row v-else>
  <v-col cols="4"  v-for="(room,i) in $apolloData.data.rooms" :key="i">
  <v-hover v-slot="{ hover }">
  <v-card 
  :elevation="hover ? 2 : 0"
    class=" p-3 rounded-lg"
    outlined
  >
    <v-list-item three-line style="align-items:flex-start">
      <v-list-item-content>
              <v-list-item-title class=" mb-1" style="font-size: 1.5rem;">
          {{room.name}}
        </v-list-item-title>
        <v-list-item-subtitle>Type : {{room.type}}</v-list-item-subtitle>
        <v-list-item-subtitle>Capacity : {{room.capacity}}</v-list-item-subtitle>
      </v-list-item-content>
    <v-chip
          label
          outlined
          class="mt-2"
          :color="room.status == 'Avalible'?'green':'red'"
    >
      {{room.status}}
    </v-chip>
     
    </v-list-item>

    <v-card-actions style="justify-content: end;">
      <v-btn
          color="pink"
          dark
          @click.stop="selected(room)"
        >
          Book
        </v-btn>
    </v-card-actions>
  </v-card>
  </v-hover>
</v-col>
</v-row>

    </v-container>
       </div>
  </v-app>
</v-col>

<v-col cols="auto" class="my-4">
      <v-navigation-drawer 
      class="rounded-l-3xl"
      mini-variant-width="0"
      :mini-variant.sync="mini"
      right
      permanent
    >
    <!-- {{select}} -->
 <div class="p-5">
    <v-list-item three-line style="align-items:flex-start">
      <v-list-item-content>
        <v-list-item-title  style="font-size: 1.5rem;">
          {{select.name}}
        </v-list-item-title>
        <v-list-item-subtitle class=" mt-3">Type : {{select.type}}</v-list-item-subtitle>
        <v-list-item-subtitle>Capacity : {{select.capacity}}</v-list-item-subtitle>

        <v-list-item-title class=" mt-3">
          Facility
          </v-list-item-title>
          <v-row no-gutters>
            <v-col sm="6" v-for="i in  7" :key="i">
            <v-list-item-subtitle>
             <li > {{i}} </li>
             </v-list-item-subtitle>
             </v-col>
          </v-row>
           
         
        
      </v-list-item-content>
      <v-chip
          label
          x-small
          outlined
          class="mt-2"
          :color="select.status == 'Avalible'?'#4CAF50':'#F44336'"
    >
      {{select.status}}
    </v-chip>

    </v-list-item>
    <v-form v-model="valid" ref="form">
  <v-textarea
          label="Purpose"
          v-model="purpose"
            :rules="Rules"

          rows="3"
          row-height="20"
        ></v-textarea>

 <v-menu
        ref="menu2"
        v-model="menu2"
        :close-on-content-click="false"
        :return-value.sync="dates2"
        transition="scale-transition"
        offset-y
        :nudge-left="100"

        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            label="Date"
            v-model="dateRangeText2"
            readonly
            :rules="Rules"
                   v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker
          v-model="dates2"
          range
          no-title
          scrollable
        >
          <v-spacer></v-spacer>
          <v-btn
            text
            color="primary"
            @click="menu2 = false"
          >
            Cancel
          </v-btn>
          <v-btn
            text
            color="primary"
            @click="$refs.menu2.save(dates2)"
          >
            OK
          </v-btn>
        </v-date-picker>
      </v-menu>
       <v-text-field
            label="Attendance"
            v-model="attendance"
            type="number"
            :rules="Rules"
            required
          ></v-text-field>
</v-form>

    <v-card-actions style="justify-content: space-evenly;">

      <v-btn
          text
          @click.stop="mini = !mini,
          select=  {}
          dates2= []"
        >
          Cancel
        </v-btn>
         <v-btn
         text
          color="#2196F3"
          @click.stop="check()"
        >
          Reserve
        </v-btn>
<v-dialog
      v-model="dialog"
      max-width="500"
      persistent
    >
    <v-card style="font-family: kanit;">
        <v-card-title style="font-size: 2rem;">
          Petition For {{select.name}}

        </v-card-title>

        <v-card-text>
            <span style="font-size: 1.2rem;"> Purpose</span> 
          <v-divider class="mb-5"></v-divider>
        <v-icon
      large
    >
      mdi-file-document
    </v-icon> {{purpose}}
        </v-card-text>
        <v-card-text>
            <span style="font-size: 1.2rem;"> Attendance</span> 
          <v-divider class="mb-5"></v-divider>
          <v-icon
      large
    >
      mdi-account
    </v-icon>{{attendance}}
        </v-card-text>
        <v-card-text>
            <span style="font-size: 1.2rem;"> Date</span> 
          <v-divider class="mb-5"></v-divider>
          <v-icon
      large
    >
      mdi-calendar
    </v-icon>{{dateRangeText2}}
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-card-text>Reserve By 62070xxx</v-card-text>
          <v-btn
            color="green darken-1"
            text
            @click="dialog = false"
          >
            Disagree
          </v-btn>

          <v-btn
            color="green darken-1"
            text
            @click="reserve()"
          >
            Agree
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    </v-card-actions>
    </div>
    </v-navigation-drawer>
    </v-col>
  </v-row>
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
    dates: [],
    menu: false,
    attendance:null,
    valid: false,
    name:'',
    Rules: [
        v => !!v || 'Plaese Enter',
      ],
    capacity: '',
    dialog: false,
    purpose:'', 
    select:{},


    menu2:false,
    dates2: [],
    mini: true,


  }),
  apollo: {
    rooms: {
      query: ALL_ROOMS,
    },
  },
  methods: {
    selected(room) {
      this.mini = false
      this.select = room
    },
    check(){
      this.$refs.form.validate()
      if(this.valid){
        this.dialog = true
      }
    },
    reserve(){
      this.dialog = false,
      this.mini = !this.mini
      this.purpose = ''
      this.dates2 = []
      this.attendance = ''
    },
  
  },
   computed: {
      dateRangeText () {
        return this.dates.join(' ~ ')
      },
      dateRangeText2 () {
        return this.dates2.join(' ~ ')
      },
    },
};

</script>
<style>
.v-application--wrap {
    min-height: 95vh
}
</style>
