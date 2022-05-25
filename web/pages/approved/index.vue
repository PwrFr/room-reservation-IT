<template>
  <v-app class="my-4 mx-4 rounded-3xl">
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

      <v-row style="height: 70vh; overflow-y: scroll">
        <v-col v-for="(item, i) in items" :key="i" md="4" class="mb-5">
          <StatusCardVue :item="item" :readMore="readMore" :staff="true" />
        </v-col>
      </v-row>
    </v-container>

    <v-card-actions style="justify-content: space-evenly">
      <v-dialog v-model="dialog" max-width="500">
        <PetitionCard :confirmation="confirmation" :room="select" />
      </v-dialog>
    </v-card-actions>
  </v-app>
</template>

<script>
import PetitionCard from "../../components/PetitionCard.vue";
import StatusCardVue from "../../components/StatusCard.vue";
export default {
  components: {
    PetitionCard,
    StatusCardVue,
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
        status: "Pending",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Pending",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Pending",
      },
      {
        room: "M25",
        reserveDate: "25/5/2022",
        attendee: "63",
        status: "Pending",
      },
    ],
    select: {},
    mini: true,
    dialog: false,
  }),
  methods: {
    confirmation(msg) {
      if (confirm(`Do you want to ${msg} using Room : blabla`)) {
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
