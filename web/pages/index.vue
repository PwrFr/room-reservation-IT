<template>
  <v-row no-gutters>
    <v-col>
      <v-app class="m-4 rounded-3xl">
        <SearchBarVue :filter="filter" />
        <div class="rounded-3xl">
          <v-container
            class="p-10"
            grid-list-xs
            style="height: 74.1vh; overflow-y: scroll"
          >
            <v-row v-if="$apolloData.loading">
              <v-col cols="4" v-for="i in 6" :key="i">
                <v-skeleton-loader
                  class="m-4"
                  max-width="300"
                  type="article, actions"
                ></v-skeleton-loader>
              </v-col>
            </v-row>

            <v-row v-else-if="$apolloData.data.rooms == null"> No Rooms </v-row>

            <v-row v-else>
              <v-col
                cols="4"
                v-for="(room, i) in $apolloData.data.rooms"
                :key="i"
              >
                <RoomCardVue :room="room" :selected="selected" :check="check" />
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
        <ReserveFormVue
          ref="ReserveFormVue"
          :select="select"
          :cancelForm="cancelForm"
          :check="check"
        />
      </v-navigation-drawer>
    </v-col>
    <v-dialog v-model="dialog" max-width="500" persistent>
      <PetitionCardVue :cancel="cancel" :reserve="reserve" :room="detail" />
    </v-dialog>
  </v-row>
</template>

<script>
import gql from "graphql-tag";
import PetitionCardVue from "../components/PetitionCard.vue";
import RoomCardVue from "../components/RoomCard.vue";
import ReserveFormVue from "../components/ReserveForm.vue";
import SearchBarVue from "../components/SearchBar.vue";

const ALL_ROOMS = gql`
  query rooms {
    rooms {
      id
      name
      status
      capacity
      type
    }
  }
`;
export default {
  name: "IndexPage",
  components: {
    PetitionCardVue,
    RoomCardVue,
    ReserveFormVue,
    SearchBarVue,
  },
  data: () => ({
    dialog: false,
    mini: true,
    attendance: null,
    purpose: "",
    select: {},
    dates2: [],
  }),
  apollo: {
    rooms: {
      query: ALL_ROOMS,
    },
  },
  methods: {
    selected(room) {
      this.mini = false;
      this.select = room;
    },

    //PetitionCardVue
    reserve() {
      if (confirm("Are you sure?")) {
        (this.dialog = false), (this.mini = !this.mini);
        this.$refs.ReserveFormVue.resetForm();
      }
    },
    //PetitionCardVue
    cancel() {
      this.dialog = false;
    },
    //ReserveFormVue
    cancelForm() {
      this.$refs.ReserveFormVue.resetForm();
      (this.mini = !this.mini), (this.select = {});
      this.purpose = "";
      this.dates2 = [];
      this.attendance = "";
    },
    //ReserveFormVue
    check(p, d, a) {
      this.dialog = true;
      this.purpose = p;
      this.dateRangeText2 = d;

      this.attendance = a;
    },
    filter(n, d, c) {
      console.log(n, d, c);
    },
  },
  computed: {
    detail() {
      return {
        room: this.select.name,
        purpose: this.purpose,
        attendee: this.attendance,
        reserveDate: this.dateRangeText2,
      };
    },
  },
};
</script>
<style>
.v-application--wrap {
  min-height: 95vh;
}
</style>
