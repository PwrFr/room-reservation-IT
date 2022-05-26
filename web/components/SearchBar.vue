<template>
  <v-form class="pt-11">
    <v-container>
      <v-row justify="center">
        <v-col cols="12" md="6">
          <v-text-field
            v-model="name"
            label="Room name"
            required
            outlined
            shaped
          ></v-text-field>
        </v-col>

        <v-col cols="12" md="3">
          <v-menu
            ref="menu"
            v-model="menu"
            :close-on-content-click="false"
            :return-value.sync="dates"
            transition="scale-transition"
            offset-y
            min-width="auto"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-text-field
                label="Date"
                v-model="dateRangeText"
                prepend-icon="mdi-calendar"
                readonly
                outlined
                shaped
                v-bind="attrs"
                v-on="on"
              ></v-text-field>
            </template>
            <v-date-picker v-model="dates" range no-title scrollable>
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="menu = false"> Cancel </v-btn>
              <v-btn text color="primary" @click="$refs.menu.save(dates)">
                OK
              </v-btn>
            </v-date-picker>
          </v-menu>
        </v-col>

        <v-col cols="12" md="1">
          <v-text-field
            v-model="capacity"
            label="Capacity"
            type="number"
            required
            outlined
            shaped
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="2">
          <v-btn
            class="mt-1"
            color="secondary"
            block
            x-large
            outlined
            style="border-top-left-radius: 1rem; border-top-right-radius: 1rem"
            @click="filter(name, dates, capacity)"
          >
            Filter
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<script>
export default {
  props: ["filter"],
  data() {
    return {
      dates: [],
      menu: false,
      name: "",
      capacity: "",
    };
  },

  computed: {
    dateRangeText() {
      return this.dates.join(" ~ ");
    },
  },
};
</script>

<style></style>
