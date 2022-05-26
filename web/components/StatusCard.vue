<template>
  <v-hover v-slot="{ hover }">
    <v-card class="rounded-lg" :elevation="hover ? 4 : 0" outlined>
      <v-row no-gutters align="center">
        <v-col md="4" class="py-6 text-3xl text-center">
          <span style="color: #e91e63">{{ item.room.room_name }} </span>
        </v-col>

        <v-col md="8" class="py-6">
          <span>Date : </span>

          {{ item.start_datetime }} ~ {{ item.end_datetime }} <br />

          <span> Attendee : </span>
          {{ item.request_attendee }}<br />

          <span> Status : </span>
          <v-chip
            label
            outlined
            class=""
            :color="
              item.request_status == 'approved'
                ? '#4CAF50'
                : item.request_status == 'pending'
                ? '#FFC107'
                : '#F44336'
            "
          >
            {{ item.request_status }} </v-chip
          ><br />
        </v-col>
        <v-row v-if="approve" justify="center">
          <v-btn
            color="warning"
            class="px-10"
            rounded
            dark
            @click="readMore(item, index)"
          >
            More
          </v-btn>
        </v-row>
      </v-row>
    </v-card>
  </v-hover>
</template>

<script>
export default {
  props: ["item", "readMore", "staff", "approve", "index"],
};
</script>

<style></style>
