<template>
  <div class="p-5">
    <v-list-item three-line style="align-items: flex-start">
      <v-list-item-content>
        <v-list-item-title style="font-size: 1.5rem">
          {{ select.name }}
        </v-list-item-title>
        <v-list-item-subtitle class="mt-3"
          >Type : {{ select.type }}</v-list-item-subtitle
        >
        <v-list-item-subtitle
          >Capacity : {{ select.capacity }}</v-list-item-subtitle
        >

        <v-list-item-title class="mt-3"> Facility </v-list-item-title>
        <v-row no-gutters>
          <v-col sm="6" v-for="i in 7" :key="i">
            <v-list-item-subtitle>
              <li>{{ i }}</li>
            </v-list-item-subtitle>
          </v-col>
        </v-row>
      </v-list-item-content>
      <v-chip
        label
        x-small
        outlined
        class="mt-2"
        :color="select.status == 'Avalible' ? '#4CAF50' : '#F44336'"
      >
        {{ select.status }}
      </v-chip>
    </v-list-item>

    <v-form v-model="valid" ref="form">
      <v-textarea
        label="Purpose"
        v-model="purpose"
        :rules="Rules"
        rows="3"
        row-height="20"
      ></v-textarea>

      <v-menu
        ref="menu2"
        v-model="menu2"
        :close-on-content-click="false"
        :return-value.sync="dates2"
        transition="scale-transition"
        offset-y
        :nudge-left="100"
        min-width="auto"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            label="Date"
            v-model="dateRangeText2"
            readonly
            :rules="Rules"
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker v-model="dates2" range no-title scrollable>
          <v-spacer></v-spacer>
          <v-btn text color="primary" @click="menu2 = false"> Cancel </v-btn>
          <v-btn text color="primary" @click="$refs.menu2.save(dates2)">
            OK
          </v-btn>
        </v-date-picker>
      </v-menu>
      <v-text-field
        label="Attendance"
        v-model="attendance"
        type="number"
        :rules="Rules"
        required
      ></v-text-field>
    </v-form>

    <v-card-actions style="justify-content: space-evenly">
      <v-btn
        text
        @click.stop="
          cancelForm();
          $refs.form.resetValidation();
        "
      >
        Cancel
      </v-btn>
      <v-btn
        text
        color="#2196F3"
        @click.stop="validate(purpose, dateRangeText2, attendance)"
      >
        Reserve
      </v-btn>
    </v-card-actions>
  </div>
</template>

<script>
export default {
  props: ["select", "cancelForm", "check"],
  data() {
    return {
      attendance: null,
      valid: false,
      Rules: [(v) => !!v || "Plaese Enter"],
      purpose: "",
      menu2: false,
      dates2: [],
    };
  },
  methods: {
    validate(purpose, dateRangeText2, attendance) {
      if (this.$auth.loggedIn) {
        this.$refs.form.validate();
        if (this.valid) {
          this.check(purpose, dateRangeText2, attendance);
        }
      } else {
        alert("You need to Login First!");
      }
    },
    resetForm() {
      this.$refs.form.reset();
      this.dates2 = [];
    },
  },
  computed: {
    dateRangeText2() {
      return this.dates2.join(" ~ ");
    },
  },
};
</script>

<style></style>
