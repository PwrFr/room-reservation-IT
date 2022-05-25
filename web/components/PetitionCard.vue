<template>
  <v-card style="font-family: kanit">
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
      <v-icon large> mdi-calendar </v-icon>{{ room.reserveDate }}
    </v-card-text>

    <v-card-actions v-if="staff">
      <v-spacer></v-spacer>
      <v-card-text>Reserve By 62070xxx</v-card-text>
      <v-btn
        color="error"
        dark
        rounded
        @click="confirmation(room.room, 'Reject')"
      >
        Reject <v-icon right dark> mdi-close </v-icon></v-btn
      >
      <v-btn
        color="success"
        rounded
        dark
        @click="confirmation(room.room, 'Approve')"
      >
        Approve
        <v-icon right dark> mdi-check </v-icon></v-btn
      >
    </v-card-actions>

    <v-card-actions v-else>
      <v-spacer></v-spacer>
      <v-card-text>Reserve By 62070xxx</v-card-text>
      <v-btn color="green darken-1" text @click="cancel()"> Disagree </v-btn>

      <v-btn color="green darken-1" text @click="reserve()"> Agree </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: ["confirmation", "room", "staff", "cancel", "reserve"],
};
</script>

<style></style>
