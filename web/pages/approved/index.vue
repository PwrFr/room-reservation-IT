<template>
  <v-app class="my-4 mx-4 rounded-3xl">
    <div
      v-if="
        $auth.loggedIn && $auth.$storage.getLocalStorage('role') === 'staff'
      "
    >
      <v-container>
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
        <v-row
          style="height: 70vh; overflow-y: scroll"
          v-if="$apolloData.loading"
        >
          <v-col v-for="(item, i) in 6" :key="i" md="4">
            <v-skeleton-loader
              max-width="300"
              type="list-item-three-line, actions"
            ></v-skeleton-loader
          ></v-col>
        </v-row>
        <v-row v-else style="height: 70vh; overflow-y: scroll">
          <v-col v-for="(item, i) in items" :key="i" md="4" class="mb-5">
            <StatusCardVue :item="item" :readMore="readMore" :staff="true" />
          </v-col>
        </v-row>
      </v-container>

      <v-card-actions style="justify-content: space-evenly">
        <v-dialog v-model="dialog" max-width="500">
          <PetitionCard
            :confirmation="confirmation"
            :room="select"
            :staff="true"
          />
        </v-dialog>
      </v-card-actions>
    </div>
    <v-row v-else justify="center" align="center">
      <h1 style="font-size: 5rem">You're not Staff GTFO</h1>
    </v-row>
  </v-app>
</template>

<script>
import gql from "graphql-tag";
import PetitionCard from "../../components/PetitionCard.vue";
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
    PetitionCard,
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
        purpose: "Want to study",
      },
      {
        room: "M23",
        reserveDate: "23/5/2022",
        attendee: "61",
        status: "Pending",
        purpose: "Want to study",
      },
      {
        room: "M24",
        reserveDate: "24/5/2022",
        attendee: "62",
        status: "Pending",
        purpose: "Want to study",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Pending",
        purpose: "Want to study",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Pending",
        purpose: "Want to study",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Pending",
        purpose: "Want to study",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Pending",
        purpose: "Want to study",
      },
    ],
    select: {},
    mini: true,
    dialog: false,
  }),
  methods: {
    confirmation(room, msg) {
      if (confirm(`Do you want to ${msg} using Room : ${room}`)) {
        this.dialog = !this.dialog;
      } else {
      }
    },
    readMore(item) {
      this.dialog = true;
      this.select = item;
    },
  },
};
</script>
