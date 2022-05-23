<template>
    <v-app>
        <v-container grid-list-xs>
            <h1 class="text-3xl font-bold text-right mb-8">หน้าการจอง</h1>
            <v-card elevation="5" outlined>
                <v-row no-gutters align="center" justify="center">
                    <v-col cols="12" sm="6" md="4" class="py-6 font-bold text-3xl text-center">
                        <span>
                            M 22
                        </span>
                    </v-col>

                    <v-col cols="12" sm="6" md="3" class="py-6">
                        <span class="font-bold">
                            Reserve Date
                        </span><br>
                        23 / 5 /2022 10 : 52 :12
                    </v-col>

                    <v-col cols="12" sm="6" md="3" class="py-6">
                        <span class="font-bold">
                            Attendee
                        </span><br>
                        60
                    </v-col>

                    <v-col cols="12" sm="6" md="2" class="py-6">
                        <span class="font-bold">
                            Status
                        </span><br>
                        Pending...
                    </v-col>
                </v-row>
            </v-card>




 <v-row >
        <v-col
          v-for="(item, i) in items"
          cols="12"
          md="4"
          :key="i"
        >
          <v-hover v-slot="{ hover }">
            <v-card
              :elevation="hover ? 12 : 2"
              :class="{ 'on-hover': hover }"
              
            >
              <v-img
                :src="item.img"
                height="225px"
              >
                <v-card-title class="text-h6 white--text">
                  <v-row
                    class="fill-height flex-column"
                    justify="space-between"
                  >
                    <p class="mt-4 subheading text-left">
                      {{ item.title }}
                    </p>

                    <div>
                      <p class="ma-0 text-body-1 font-weight-bold font-italic text-left">
                        {{ item.text }}
                      </p>
                      <p class="text-caption font-weight-medium font-italic text-left">
                        {{ item.subtext }}
                      </p>
                    </div>
                  </v-row>
                </v-card-title>
              </v-img>
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
      items: [
        { 
          room : "M22",
          reserveDate: "23/5/2022 10:15:20",
          attendee: "60",
          status: "Pending..."   , 
          title: 'New Releases',
          text: `It's New Release Friday`,
          subtext: 'Newly released songs. Updated daily.',
          img: 'https://images.unsplash.com/photo-1429514513361-8fa32282fd5f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=3264&q=80',
        },
        {
          title: 'Rock',
          text: 'Greatest Rock Hits',
          subtext: 'Lose yourself in rock tunes.',
          img: 'https://images.unsplash.com/photo-1498038432885-c6f3f1b912ee?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2100&q=80',
        },
        {
          title: 'Mellow Moods',
          text: 'Ambient Bass',
          subtext: 'Chill beats to mellow you out.',
          img: 'https://images.unsplash.com/photo-1542320868-f4d80389e1c4?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=3750&q=80',
        },
      ],
      transparent: 'rgba(255, 255, 255, 0)',
    }),

    
};
</script>


<style scoped>
.v-card {
  transition: opacity .4s ease-in-out;
}

.v-card:not(.on-hover) {
  opacity: 0.6;
 }

.v-application--wrap {
    min-height: 94vh
}

</style>