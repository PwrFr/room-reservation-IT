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
              :index="i"
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
            :index="index"
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
      room {
        room_name
      }
      request_purpose
      request_attendee
      request_status
      start_datetime
      end_datetime
      request_by
    }
  }
`;

const UPD_REQ = gql`
  mutation updateRequest(
    $staff: String!
    $reqId: Int!
    $status: String!
    $room: Int!
  ) {
    updateRequest(
      input: {
        approve_by: $staff
        request_id: $reqId
        request_status: $status
        room_id: $room
      }
    ) {
      request_id
      request_status
      approve_by
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
    index: null,
  }),
  methods: {
    async confirmation(req, room, msg, i) {
      // alert(i);
      if (confirm(`Do you want to ${msg} using Room : ${room}`)) {
        const update = await this.$apollo
          .mutate({
            mutation: UPD_REQ,
            variables: {
              staff: this.$auth.user.sub,
              reqId: parseInt(req),
              status: msg,
              room: parseInt(room),
            },
          })
          .then((res) => console.log(res))
          .finally(() => this.$router.go());
        // console.log(update);
        this.dialog = !this.dialog;
      } else {
      }
    },
    readMore(item, i) {
      this.dialog = true;
      this.select = item;
      this.index = i;
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
  // mounted() {
  //   console.log(this.$apollo.queries);
  // },
};
</script>
