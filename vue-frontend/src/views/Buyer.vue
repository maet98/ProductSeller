<template>
    <div>
      <v-card class="overflow-hidden" v-if="!this.loading">
        <v-app-bar
          absolute
          elevate-on-scroll
          scroll-target="#scrolling-techniques-7"
        >
          <v-app-bar-nav-icon></v-app-bar-nav-icon>

          <v-spacer></v-spacer>
          <v-toolbar-title>{{this.buyer_info.name}}</v-toolbar-title>
          <v-spacer></v-spacer>

          <router-link :to="{ name: 'Home'}">
            <v-icon>mdi-home</v-icon>
          </router-link>
        </v-app-bar>
        <v-container>
            <v-row
              class="mb-6"
              no-gutters
            >
                <v-col cols="3">
                    <OtherBuyers :changeId="this.changeId" :buyers="this.others_buyers" />
                </v-col>

                <v-col cols="6">
                    <History :purchases="this.buyer_info.Purchases" />
                    <RecommendedProducts :Recommendations="this.recommendation" />
                </v-col>
            </v-row>
          </v-container>
      </v-card>
      <div class="text-center" v-else>
          <v-progress-circular
          indeterminate
          color="primary"
        >loading</v-progress-circular>
      </div>
    </div>
</template>

<script>
    import axios from "axios";
    import History from "../components/History";
    import OtherBuyers from "../components/OtherBuyers";
    import RecommendedProducts from "../components/recommendationProducts";

    export default {
        components: {
            History,
            OtherBuyers,
            RecommendedProducts
        },
        name: "buyer",
        props: ['id'],
        data() {
            return {
                loading: false,
                uid: this.$props.id,
                buyer_info: null,
                recommendation: null,
                others_buyers: null
            }
        },
        mounted() {
            this.fetchAll();
        },
        methods: {
            fetchAll() {
                this.loading = true;
                axios.get("http://localhost:3000/buyers/" + this.uid)
                    .then(Response => {
                        this.loading = false;
                        this.buyer_info = Response.data.buyer[0];
                        this.others_buyers = Response.data.Others;
                        this.recommendation = Response.data.recommendations;
                        console.log(this.others_buyers);
                })
                .catch((err)=>{
                    console.log(err)
                    this.loading = false;
                });
            },
            changeId( id ) {
                console.log( id );
                this.uid = id;
                this.fetchAll();
            }
        }

    }

</script>
