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
          <v-hover v-slot="{ hover }">
            <v-card class="rounded-lg" :elevation="hover ? 4 : 0" outlined>
              <v-row no-gutters align="center">
                <v-col md="4" class="py-6 font-bold text-3xl text-center">
                  <span style="color: #e91e63">
                    {{ item.room }}
                  </span>
                </v-col>

                <v-col md="8" class="py-6">
                  <span class="font-bold"> Reserve Date : </span>
                  {{ item.reserveDate }}<br />

                  <span class="font-bold"> Attendee : </span>
                  {{ item.attendee }}<br />

                  <span class="font-bold"> Status : </span>
                  <v-chip
                    label
                    outlined
                    class=""
                    :color="
                      item.status == 'Approved'
                        ? '#4CAF50'
                        : item.status == 'Pending'
                        ? '#FFC107'
                        : '#F44336'
                    "
                  >
                    {{ item.status }} </v-chip
                  ><br />
                </v-col>
                <v-row justify="space-around">
                  <v-btn
                    color="error"
                    dark
                    rounded
                    @click="confirmation(item, 'Reject')"
                  >
                    Reject <v-icon right dark> mdi-close </v-icon></v-btn
                  >
                  <v-btn
                    color="success"
                    rounded
                    dark
                    @click="confirmation(item, 'Approve')"
                  >
                    Approve
                    <v-icon right dark> mdi-check </v-icon></v-btn
                  >
                </v-row>
              </v-row>
            </v-card>
          </v-hover>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
export default {
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
  }),
  methods: {
    confirmation(room, msg) {
      confirm(`Do you want to ${msg} using Room : ${room.room}`)
    },
  },
};
</script>
