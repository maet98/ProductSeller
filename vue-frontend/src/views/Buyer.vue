<template>
<v-container class="grey lighten-5">
    <v-row
      class="mb-6"
      no-gutters
    >
        <v-col>
            columns
        </v-col>

        <v-col cols="6">
            <History :history="this.buyer_info.Compras" />
            {{ this.buyer_info }} 
        </v-col>

        <v-col>
            columns
        </v-col>
    </v-row>
  </v-container>
</template>

<script>
    import axios from "axios";
    import History from "../components/History";

    export default {
        components: {
            History
        },
        name: "buyer",
        props: ['id'],
        data() {
            return {
                buyer_info: null,
                recommendation: null,
                others_buyers: null
            }
        },
        mounted() {
            axios.get("http://localhost:3000/buyers/" + this.$props.id)
                .then(Response => {
                    this.buyer_info = Response.data.buyer[0];
                    this.others_buyers = Response.data.others;
                    this.recommendation = Response.data.recommendation;
                    console.log(this.buyer_info);
            })
            .catch(()=> console.log("error"));
        }
    }

</script>
