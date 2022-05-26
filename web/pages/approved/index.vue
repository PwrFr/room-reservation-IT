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

            <v-btn value="pending" color="warning" text>
              <span class="hidden-sm-and-down">Pending</span>
            </v-btn>

            <v-btn value="approved" color="success" text>
              <span class="hidden-sm-and-down">Approved</span>
            </v-btn>

            <v-btn value="rejected" color="error" text>
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
          <v-col v-for="(item, i) in filterReq" :key="i" md="4" class="mb-5">
            <StatusCardVue
              :item="item"
              :readMore="readMore"
              :staff="true"
              :approve="item.request_status == 'pending' ? true : false"
            />
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
      <h1 style="font-size: 5rem">You're not Staff</h1>
    </v-row>
  </v-app>
</template>

<script>
import gql from "graphql-tag";
import PetitionCard from "../../components/PetitionCard.vue";
import StatusCardVue from "../../components/StatusCard.vue";
const REQ_ROOMS = gql`
  query request {
    request {
      request_id
      room_id
      request_purpose
      request_attendee
      request_status
      start_datetime
      end_datetime
      request_by
    }
  }
`;
export default {
  components: {
    PetitionCard,
    StatusCardVue,
  },
  apollo: {
    request: {
      query: REQ_ROOMS,
    },
  },
  data: () => ({
    text: "pending",
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
  computed: {
    filterReq() {
      return this.$apolloData.data.request.filter((req) => {
        if (this.text !== "all") {
          return req.request_status.match(this.text);
        } else {
          return this.$apolloData.data.request;
        }
      });
    },
  },
};
</script>
