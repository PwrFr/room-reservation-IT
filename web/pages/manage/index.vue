<template>
  <v-row no-gutters>
    <v-col>
      <v-app class="my-4 mx-4 rounded-3xl">
        <v-container v-if="$auth.loggedIn">
          <v-col cols="12" class="py-10">
            <v-btn-toggle v-model="text" mandatory group>
              <v-btn value="all">
                <span class="hidden-sm-and-down">All</span>
              </v-btn>

              <v-btn value="Available" color="success" text>
                <span class="hidden-sm-and-down">Available</span>
              </v-btn>

              <v-btn value="Not available" color="error" text>
                <span class="hidden-sm-and-down">Not available</span>
              </v-btn>

              <v-btn value="Maintenance" color="primary" text>
                <span class="hidden-sm-and-down">Maintenance</span>
              </v-btn>
            </v-btn-toggle>
          </v-col>
          <v-row
            v-if="$apolloData.loading"
            style="height: 70vh; overflow-y: scroll"
          >
            <v-col v-for="(item, i) in 9" :key="i" md="4">
              <v-skeleton-loader
                max-width="300"
                type="article"
              ></v-skeleton-loader
            ></v-col>
          </v-row>

          <v-row v-else style="height: 70vh; overflow-y: scroll">
            <!-- <div style="display: none">
          {{ request }}
        </div> -->
            <v-col v-for="(item, i) in filterStatus" :key="i" md="4">
              <v-hover v-slot="{ hover }">
                <v-card
                  :elevation="hover ? 2 : 0"
                  class="p-3 rounded-lg"
                  outlined
                >
                  <v-list-item three-line style="align-items: flex-start">
                    <v-list-item-content>
                      <v-list-item-title class="mb-1" style="font-size: 1.5rem">
                        {{ item.room_name }}
                      </v-list-item-title>
                      <v-list-item-subtitle v-if="item.request">
                        Start :
                        {{
                          item.request[item.request.length - 1].start_datetime
                        }}
                        <br />
                        End :
                        {{ item.request[item.request.length - 1].end_datetime }}
                      </v-list-item-subtitle>
                    </v-list-item-content>
                    <v-chip
                      label
                      outlined
                      class="mt-2"
                      :color="
                        item.room_status == 'Available'
                          ? 'green'
                          : item.room_status == 'Not available'
                          ? 'red'
                          : 'primary'
                      "
                    >
                      {{ item.room_status }}
                    </v-chip>
                  </v-list-item>

                  <v-card-actions
                    v-if="filter(item.request)"
                    style="justify-content: end"
                  >
                    <v-btn
                      color="indigo"
                      dark
                      @click="
                        mini = false;
                        select = item;
                      "
                      >Manage</v-btn
                    >
                  </v-card-actions>
                </v-card>
              </v-hover>
              <!-- <StatusCardVue :item="item" /> -->
            </v-col>
          </v-row>
        </v-container>
        <v-row v-if="!$auth.loggedIn" justify="center" align="center">
          <h1 style="font-size: 5rem">Please LogIn</h1>
        </v-row>
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
        <v-list-item three-line style="align-items: flex-start">
          <v-list-item-content>
            <v-list-item-subtitle style="font-size: 1.4rem; color: black">
              {{ select.room_name }}
            </v-list-item-subtitle>
            <v-list-item-subtitle v-if="select.room_type" class="mt-3">
              Type : {{ select.room_type.type_name }}
            </v-list-item-subtitle>
            <v-list-item-subtitle
              >Capacity : {{ select.room_capacity }}</v-list-item-subtitle
            >
          </v-list-item-content>
          <v-chip
            label
            x-small
            outlined
            class="mt-2"
            :color="
              select.room_status == 'Available'
                ? '#4CAF50'
                : select.room_status == 'Not available'
                ? '#F44336'
                : '#2196F3'
            "
          >
            {{ select.room_status }}
          </v-chip>
        </v-list-item>

        <v-list-item>
          <v-select
            class="mx-4"
            v-model="status"
            :items="items"
            label="Status"
          ></v-select>
        </v-list-item>
        <v-card-actions style="justify-content: space-evenly">
          <v-btn
            text
            @click="
              mini = true;
              select = {};
              status = '';
            "
            >Cancel</v-btn
          >
          <v-btn color="#2196F3" text @click="update(select.room_id)"
            >Change</v-btn
          >
        </v-card-actions>
      </v-navigation-drawer>
    </v-col>
  </v-row>
</template>

<script>
import gql from "graphql-tag";
const REQ_ROOMS = gql`
  query roomWithRequest {
    roomWithRequest {
      room_id
      room_name
      room_capacity
      room_status
      room_type {
        type_name
      }
      type_id
      # room_type {
      #   type_id
      #   type_name
      # }
      # room_facility {
      #   quantity
      #   facility {
      #     facility_name
      #   }
      # }
      request {
        request_id
        request_status
        start_datetime
        end_datetime
        request_datetime
      }
    }
  }
`;

const UPDATE_ROOM = gql`
  mutation updateRoom($id: Int!, $status: String!, $token: String!) {
    updateRoom(room_id: $id, status: $status, token: $token)
  }
`;
export default {
  data() {
    return {
      text: "all",
      mini: true,
      select: {},
      items: ["Available", "Not available", "Maintenance"],
      status: "",
    };
  },
  apollo: {
    roomWithRequest: {
      query: REQ_ROOMS,
    },
  },
  methods: {
    async update(id) {
      if (confirm("Are you sure?")) {
        await this.$apollo
          .mutate({
            mutation: UPDATE_ROOM,
            variables: {
              id: id,
              status: this.status,
              token: this.$auth.$storage.getLocalStorage("token"),
            },
          })
          .finally(() => this.$router.replace({ path: "/" }));
        this.mini = true;
      }
    },
  },
  computed: {
    filter() {
      return (req) => {
        if (!req) {
          return true;
        }

        const filterReqStatus = req.filter((r) => {
          return r.request_status == "approved";
        });

        var max = filterReqStatus[0];
        for (let i = 0; i < filterReqStatus.length; i++) {
          if (filterReqStatus[i].end_datetime > max.end_datetime) {
            max = filterReqStatus[i];
          }
        }
        if (max) {
          if (new Date(max.end_datetime) > new Date()) {
            return false;
          }
        }
        return true;
      };
    },
    filterStatus() {
      return this.roomWithRequest.filter((room) => {
        if (this.text !== "all") {
          return room.room_status.match(this.text);
        }
        return this.roomWithRequest;
      });
    },
  },
};
</script>

<style></style>
