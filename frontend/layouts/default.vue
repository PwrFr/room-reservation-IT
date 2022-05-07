

<template>
  <div class="px-3 py-5" style="background-color: #dcdde0; font-family: kanit">
    <v-row no-gutters>
      <v-col cols="auto">
        <v-navigation-drawer fixed permanent expand-on-hover>
          <v-list>
            <v-list-item style="padding-left: 0.5rem">
              <v-list-item-avatar>
                <lord-icon
                  target=".v-list"
                  src="https://cdn.lordicon.com/nobciafz.json"
                  trigger="morph"
                  style="width: 3rem; height: 3rem"
                >
                </lord-icon>
              </v-list-item-avatar>
            </v-list-item>

            <v-list-item link>
              <v-list-item-content>
                <v-list-item-title> Pawaris Wongsaied </v-list-item-title>
                <v-list-item-subtitle>pwrfr@gmail.co.th</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-list>

          <v-divider></v-divider>

          <v-list nav style="padding-left: 0rem">
            <NuxtLink to="/">
              <v-list-item link>
                <v-list-item-title class="icon">
                  <lord-icon
                    target=".v-list-item"
                    src="https://cdn.lordicon.com/gmzxduhd.json"
                    trigger="morph"
                    style="width: 2.5rem; height: 2.5rem"
                  >
                  </lord-icon>
                  หน้าหลัก
                </v-list-item-title>
              </v-list-item>
            </NuxtLink>
            <NuxtLink to="/booking">
              <v-list-item link>
                <v-list-item-title
                  ><lord-icon
                    target=".v-list-item"
                    src="https://cdn.lordicon.com/nocovwne.json"
                    trigger="morph"
                    state="hover-2"
                    style="width: 2.5rem; height: 2.5rem"
                  >
                  </lord-icon
                  >จองห้อง</v-list-item-title
                >
              </v-list-item>
            </NuxtLink>
            <v-list-item link>
              <v-list-item-title
                ><lord-icon
                  target=".v-list-item"
                  src="https://cdn.lordicon.com/puvaffet.json"
                  trigger="morph"
                  style="width: 2.5rem; height: 2.5rem"
                >
                </lord-icon
                >ห้องที่จองไว้</v-list-item-title
              >
            </v-list-item>
            <v-list-item link>
              <v-list-item-title
                ><lord-icon
                  target=".v-list-item"
                  src="https://cdn.lordicon.com/yyecauzv.json"
                  trigger="morph"
                  style="width: 2.5rem; height: 2.5rem"
                >
                </lord-icon
                >อนุมัติการจองห้อง</v-list-item-title
              >
            </v-list-item>
          </v-list>
        </v-navigation-drawer>
      </v-col>
      <v-col>
        <Nuxt class="bg-white ml-16 p-6 rounded-md" />
      </v-col>
    </v-row>
  </div>
</template>

sc