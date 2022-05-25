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
        <v-col v-for="(item, i) in items" :key="i" md="4">
          <StatusCardVue :item="item" />
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
import StatusCardVue from "../../components/StatusCard.vue";

export default {
  components: {
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
