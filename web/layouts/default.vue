<template>
  <div style="background-color: #dcdde0; font-family: kanit">
    <v-row no-gutters>
      <v-col cols="auto ">
        <v-navigation-drawer permanent expand-on-hover class="rounded-r-3xl">
          <v-list v-if="$auth.loggedIn">
            <v-list-item style="padding-left: 0.5rem">
              <v-list-item-avatar>
                <v-avatar>
                  <img v-auth-image="$auth.user.picture" />
                </v-avatar>
              </v-list-item-avatar>
            </v-list-item>
            <v-list-item link>
              <v-list-item-content>
                <v-list-item-title>
                  {{ $auth.user.name }}
                </v-list-item-title>

                <v-list-item-subtitle>{{
                  $auth.user.email
                }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-list>
          <v-list v-else>
            <v-list-item style="padding-left: 0.5rem">
              <v-list-item-avatar>
                <lord-icon
                  target=".v-list"
                  src="https://cdn.lordicon.com/dxjqoygy.json"
                  trigger="morph"
                  style="width: 3rem; height: 3rem"
                >
                </lord-icon>
              </v-list-item-avatar>
            </v-list-item>
            <v-list-item link>
              <v-list-item-content>
                <v-list-item-title> Anonymous </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>

          <v-divider></v-divider>
          <v-list shaped>
            <v-list-item
              v-for="(item, i) in items"
              :key="i"
              :to="item.to"
              router
              exact
              style="padding-left: 0.4rem"
            >
              <v-list-item-content>
                <v-list-item-title>
                  <lord-icon
                    target=".v-list-item"
                    :src="item.src"
                    trigger="morph"
                    style="width: 2.5rem; height: 2.5rem"
                  >
                  </lord-icon>
                  {{ item.title }}</v-list-item-title
                >
              </v-list-item-content>
            </v-list-item>

            <v-list-item
              v-if="$auth.loggedIn"
              @click="logout"
              exact
              router
              style="padding-left: 0.33rem"
            >
              <v-list-item-content>
                <v-list-item-title>
                  <lord-icon
                    target=".v-list-item"
                    src="https://cdn.lordicon.com/iiueiwdd.json"
                    trigger="morph"
                    style="width: 2.5rem; height: 2.5rem"
                  >
                  </lord-icon>
                  Log Out</v-list-item-title
                >
              </v-list-item-content>
            </v-list-item>
            <v-list-item
              v-else
              @click="login"
              router
              exact
              style="padding-left: 0.4rem"
            >
              <v-list-item-content>
                <v-list-item-title>
                  <lord-icon
                    target=".v-list-item"
                    src="https://cdn.lordicon.com/gkditgni.json"
                    trigger="morph"
                    style="width: 2.5rem; height: 2.5rem"
                  >
                  </lord-icon>
                  Login</v-list-item-title
                >
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-navigation-drawer>
      </v-col>
      <v-col>
        <Nuxt class="mx-5" style="font-family: kanit" />
      </v-col>
    </v-row>
  </div>
</template>

<script>
// import { h } from "vue";

export default {
  methods: {
    login() {
      this.$auth.loginWith("google");
    },
    logout() {
      this.$auth.logout();
      location.reload();
    },
  },
  data() {
    return {
      items: [
        {
          title: "หน้าหลัก",
          src: "https://cdn.lordicon.com/gmzxduhd.json",
          to: "/",
        },

        {
          title: "ห้องที่จองไว้",
          src: "https://cdn.lordicon.com/puvaffet.json",
          to: "/booking",
        },
        {
          title: "อนุมัติการจองห้อง",
          src: "https://cdn.lordicon.com/yyecauzv.json",
          to: "/approved",
        },
      ],
    };
  },
};
</script>

<style></style>
