<template>
  <div style="background-color: #dcdde0; font-family: kanit">
    <v-row no-gutters>
      <v-col cols="auto ">
        <v-navigation-drawer permanent expand-on-hover class="rounded-r-3xl">
          <v-list v-if="$auth.loggedIn">
            <v-list-item style="padding-left: 0.5rem">
              <v-list-item-avatar>
                <v-avatar>
                  <img
                    height="500"
                    width="500"
                    v-auth-image="$auth.user.picture"
                  />
                </v-avatar>
              </v-list-item-avatar>
            </v-list-item>
            <v-list-item link>
              <v-list-item-content>
                <v-list-item-title>
                  {{ $auth.user.name }}
                </v-list-item-title>
                <div style="display: none">
                  {{ user }}
                </div>
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
              target=".v-list-item"
              to="/"
              router
              exact
              style="padding-left: 0.4rem"
            >
              <v-list-item-content>
                <v-list-item-title>
                  <lord-icon
                    target=".v-list-item"
                    src="https://cdn.lordicon.com/gmzxduhd.json"
                    trigger="morph"
                    style="width: 2.5rem; height: 2.5rem"
                  >
                  </lord-icon>
                  หน้าหลัก</v-list-item-title
                >
              </v-list-item-content>
            </v-list-item>
            <v-list-item
              v-if="role === 'student'"
              target=".v-list-item"
              to="/booking"
              router
              exact
              style="padding-left: 0.4rem"
            >
              <v-list-item-content>
                <v-list-item-title>
                  <lord-icon
                    target=".v-list-item"
                    src="https://cdn.lordicon.com/puvaffet.json"
                    trigger="morph"
                    style="width: 2.5rem; height: 2.5rem"
                  >
                  </lord-icon>
                  ห้องที่จองไว้</v-list-item-title
                >
              </v-list-item-content>
            </v-list-item>
            <v-list-item
              v-if="role === 'staff'"
              target=".v-list-item"
              to="/approved"
              router
              exact
              style="padding-left: 0.4rem"
            >
              <v-list-item-content>
                <v-list-item-title>
                  <lord-icon
                    target=".v-list-item"
                    src="https://cdn.lordicon.com/yyecauzv.json"
                    trigger="morph"
                    style="width: 2.5rem; height: 2.5rem"
                  >
                  </lord-icon>
                  อนุมัติการจองห้อง</v-list-item-title
                >
              </v-list-item-content>
            </v-list-item>

            <v-list-item
              v-if="role === 'staff'"
              target=".v-list-item"
              to="/manage"
              router
              exact
              style="padding-left: 0.4rem"
            >
              <v-list-item-content>
                <v-list-item-title>
                  <lord-icon
                    target=".v-list-item"
                    src="https://cdn.lordicon.com/huwchbks.json"
                    trigger="morph"
                    style="width: 2.5rem; height: 2.5rem"
                  >
                  </lord-icon>
                  จัดการห้อง</v-list-item-title
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
              to="/login"
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
import gql from "graphql-tag";
// import { h } from "vue";
const GET_ROLE = gql`
  mutation CreateAccount(
    $acc_id: String!
    $fname: String!
    $lname: String!
    $email: String!
  ) {
    createAccount(
      input: {
        account_id: $acc_id
        first_name: $fname
        last_name: $lname
        email: $email
      }
    ) {
      account_id
      first_name
      last_name
      role
      email
      token
    }
  }
`;
export default {
  data() {
    return {
      role: "",
    };
  },
  methods: {
    logout() {
      this.$auth.logout();
      this.$auth.$storage.removeLocalStorage("role");
      this.$auth.$storage.removeLocalStorage("token");
      this.$router.go();
    },
  },
  computed: {
    async user() {
      const role = await this.$apollo.mutate({
        mutation: GET_ROLE,
        variables: {
          acc_id: this.$auth.user.sub,
          fname: this.$auth.user.given_name,
          lname: this.$auth.user.family_name,
          email: this.$auth.user.email,
        },
      });
      this.$auth.$storage.setLocalStorage("role", role.data.createAccount.role);
      this.$auth.$storage.setLocalStorage(
        "token",
        role.data.createAccount.token
      );

      this.role = role.data.createAccount.role;
      // return this.$auth.user.name;
    },
  },
};
</script>

<style></style>
