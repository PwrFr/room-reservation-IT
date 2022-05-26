<template>
  <v-app class="my-4 mx-4 rounded-3xl">
    <v-container v-if="$auth.loggedIn">
      <v-col cols="12" class="py-10">
        <v-btn-toggle v-model="text" mandatory group>
          <v-btn value="all">
            <span class="hidden-sm-and-down">All</span>
          </v-btn>

          <v-btn value="pending">
            <span class="hidden-sm-and-down">Pending</span>
          </v-btn>

          <v-btn value="approved">
            <span class="hidden-sm-and-down">Approved</span>
          </v-btn>

          <v-btn value="rejected">
            <span class="hidden-sm-and-down">Rejected</span>
          </v-btn>
        </v-btn-toggle>
      </v-col>
      <!-- {{ $apolloData }} -->
      <v-row
        v-if="$apolloData.loading"
        style="height: 70vh; overflow-y: scroll"
      >
        <v-col v-for="(item, i) in 9" :key="i" md="4">
          <v-skeleton-loader max-width="300" type="article"></v-skeleton-loader
        ></v-col>
      </v-row>
      <v-row v-else style="height: 70vh; overflow-y: scroll">
        <v-col v-for="(item, i) in items" :key="i" md="4">
          <StatusCardVue :item="item" />
        </v-col>
      </v-row>
    </v-container>
    <v-row v-else justify="center" align="center">
      <h1 style="font-size: 5rem">Please LogIn</h1>
    </v-row>
  </v-app>
</template>

<script>
import gql from "graphql-tag";
import StatusCardVue from "../../components/StatusCard.vue";

const REQ_ROOMS = gql`
  query {
    rooms {
      room_name
    }
  }
`;
export default {
  components: {
    StatusCardVue,
  },
  apollo: {
    rooms: {
      query: REQ_ROOMS,
    },
  },
  data: () => ({
    text: "all",
    items: [
      {
        room: "M22",
        reserveDate: "22/5/2022",
        attendee: "60",
        status: "Pending",
      },
      {
        room: "M23",
        reserveDate: "23/5/2022",
        attendee: "61",
        status: "Pending",
      },
      {
        room: "M24",
        reserveDate: "24/5/2022",
        attendee: "62",
        status: "Pending",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Approved",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Approved",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Rejected",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Rejected",
      },
    ],
  }),
};
</script>

<style>
.v-application--wrap {
  min-height: 95vh;
}
</style>
