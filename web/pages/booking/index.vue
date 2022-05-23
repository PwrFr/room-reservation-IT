<template>
    <v-app>
        <v-container grid-list-xs>
            <h1 class="text-3xl font-bold text-right mb-8">หน้าการจอง</h1>
        
            <v-col cols="12" class="py-2">
                
                <v-btn-toggle v-model="text" color="deep-purple accent-3" group>
                    <v-btn value="all">
                        <span class="hidden-sm-and-down">All</span>
                    </v-btn>

                    <v-btn value="pending">
                        <span class="hidden-sm-and-down">Pending</span>
                    </v-btn>

                    <v-btn value="approved">
                        <span class="hidden-sm-and-down">Approved</span>
                    </v-btn>

                    <v-btn value="rejected">
                        <span class="hidden-sm-and-down">Rejected</span>
                    </v-btn> 
                    {{text}}
                </v-btn-toggle>
            </v-col>

            <template v-for="(item, i) in items">
                <v-col :key="i" cols="12" md="12">
                    <v-hover v-slot="{ hover }">
                        <v-card :elevation="hover ? 12 : 2" :class="{ 'on-hover': hover }">
                            <v-row no-gutters align="center">
                                <v-col cols="12" sm="4" md="4" class="py-6 font-bold text-3xl text-center">
                                    <span>
                                        {{ item.room }}
                                    </span>
                                </v-col>

                                <v-col cols="12" sm="3" md="3" class="py-6 pl-3">
                                    <span class="font-bold">
                                        Reserve Date
                                    </span><br>
                                    {{ item.reserveDate }}
                                </v-col>

                                <v-col cols="12" sm="3" md="3" class="py-6 pl-3">
                                    <span class="font-bold">
                                        Attendee
                                    </span><br>
                                    {{ item.attendee }}
                                </v-col>

                                <v-col cols="12" sm="2" md="2" class="py-6 pl-3">
                                    <span class="font-bold">
                                        Status
                                    </span><br>
                                    {{ item.status }}
                                </v-col>
                            </v-row>
                        </v-card>
                    </v-hover>
                </v-col>
            </template>
        </v-container>
    </v-app>
</template>



<script>
export default {
    data: () => ({
         text: 'all',
        items: [
            {
                room: "M22",
                reserveDate: "22/5/2022 10:15:20",
                attendee: "60",
                status: "Pending..."

            },
            {
                room: "M23",
                reserveDate: "23/5/2022 10:15:20",
                attendee: "61",
                status: "Pending..."

            },
            {
                room: "M24",
                reserveDate: "24/5/2022 10:15:20",
                attendee: "62",
                status: "Pending..."

            },
            {
                room: "M25",
                reserveDate: "25/5/2022 10:15:20",
                attendee: "63",
                status: "Approved..."

            }
        ],
    }),
};
</script>


<style scoped>
.v-card {
    transition: opacity .2s ease-in-out;
}

.v-card:not(.on-hover) {
    opacity: 0.7;
}

.v-application--wrap {
    min-height: 94vh
}
</style>