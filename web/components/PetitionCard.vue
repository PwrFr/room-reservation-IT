<template>
  <v-card style="font-family: kanit" v-if="staff">
    <v-card-title style="font-size: 2rem">
      Petition For room {{ room.room_name }}
    </v-card-title>

    <v-card-text>
      <span style="font-size: 1.2rem"> Purpose</span>
      <v-divider class="mb-5"></v-divider>
      <v-icon large> mdi-file-document </v-icon> {{ room.request_purpose }}
    </v-card-text>
    <v-card-text>
      <span style="font-size: 1.2rem"> Attendance</span>
      <v-divider class="mb-5"></v-divider>
      <v-icon large> mdi-account </v-icon>{{ room.request_attendee }}
    </v-card-text>
    <v-card-text>
      <span style="font-size: 1.2rem"> Date</span>
      <v-divider class="mb-5"></v-divider>
      <v-icon large> mdi-calendar </v-icon>
      {{ room.start_datetime }} ~ {{ room.end_datetime }}
    </v-card-text>

    <v-card-actions>
      <v-spacer></v-spacer>
      <v-card-text v-if="$auth.loggedIn"
        >Request By {{ room.email }}</v-card-text
      >
      <v-btn
        color="error"
        dark
        rounded
        @click="confirmation(room.request_id, room.room_id, 'rejected', index)"
      >
        Reject <v-icon right dark> mdi-close </v-icon></v-btn
      >
      <v-btn
        color="success"
        rounded
        dark
        @click="confirmation(room.request_id, room.room_id, 'approved', index)"
      >
        Approve
        <v-icon right dark> mdi-check </v-icon></v-btn
      >
    </v-card-actions>
  </v-card>
  <v-card style="font-family: kanit" v-else>
    <v-card-title style="font-size: 2rem">
      Petition For {{ room.room }}
    </v-card-title>

    <v-card-text>
      <span style="font-size: 1.2rem"> Purpose</span>
      <v-divider class="mb-5"></v-divider>
      <v-icon large> mdi-file-document </v-icon> {{ room.purpose }}
    </v-card-text>
    <v-card-text>
      <span style="font-size: 1.2rem"> Attendance</span>
      <v-divider class="mb-5"></v-divider>
      <v-icon large> mdi-account </v-icon>{{ room.attendee }}
    </v-card-text>
    <v-card-text>
      <span style="font-size: 1.2rem"> Date</span>
      <v-divider class="mb-5"></v-divider>
      <v-icon large> mdi-calendar </v-icon>
      {{ dateJoin }}
    </v-card-text>

    <v-card-actions>
      <v-spacer></v-spacer>
      <v-card-text v-if="$auth.loggedIn"
        >Reserve By {{ $auth.user.email }}</v-card-text
      >
      <v-btn color="green darken-1" text @click="cancel()"> Disagree </v-btn>

      <v-btn
        color="green darken-1"
        text
        @click="reserve(room.purpose, room.attendee, room.reserveDate)"
      >
        Agree
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: ["confirmation", "room", "staff", "cancel", "reserve", "index"],
  computed: {
    dateJoin() {
      return this.room.reserveDate.join(" ~ ");
    },
  },
};
</script>

<style></style>
