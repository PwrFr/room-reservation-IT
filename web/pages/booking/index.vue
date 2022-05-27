<template>
  <v-app class="my-4 mx-4 rounded-3xl">
    <v-container v-if="$auth.loggedIn">
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
        v-if="$apolloData.loading"
        style="height: 70vh; overflow-y: scroll"
      >
        <v-col v-for="(item, i) in 9" :key="i" md="4">
          <v-skeleton-loader max-width="300" type="article"></v-skeleton-loader
        ></v-col>
      </v-row>

      <v-row v-else style="height: 70vh; overflow-y: scroll">
        <div style="display: none">
          {{ request }}
        </div>
        <v-col v-for="(item, i) in filterReq" :key="i" md="4">
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

const REQ_ROOMS_BY_ID = gql`
  query requestById($id: String!) {
    requestById(account_id: $id) {
      request_id
      room_name
      request_purpose
      request_attendee
      request_status
      start_datetime
      end_datetime
    }
  }
`;
export default {
  components: {
    StatusCardVue,
  },
  apollo: {
    // requestById: {
    //   query: REQ_ROOMS_BY_ID,
    //   variables: {
    //     id: this.$auth.user.sub,
    //   },
    // },
  },
  data: () => ({
    text: "pending",
    items: [],
  }),
  computed: {
    filterReq() {
      return this.items.filter((req) => {
        if (this.text !== "all") {
          return req.request_status.match(this.text);
        } else {
          return this.items;
        }
      });
    },
    async request() {
      // console.log(this.$auth.user.sub);

      const res = await this.$apollo.query({
        query: REQ_ROOMS_BY_ID,
        variables: {
          id: await this.$auth.user.sub,
        },
      });
      this.items = res.data.requestById;
      // console.log(this.items);

      return res.data.requestById;
    },
  },
};
</script>

<style>
.v-application--wrap {
  min-height: 95vh;
}
</style>
