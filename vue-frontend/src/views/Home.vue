<template>
  <div class="home">
    <v-card-title class="justify-center">
        <h1>Buyers</h1>
    </v-card-title>
    <v-card
      class="d-flex  align-content-space-around justify-center flex-wrap"
      color="grey lighten-2"
      flat
      tile
      min-height="400"
    >
        <v-card
          class="pa-2 margin"
          max-width="344"
          v-for="item in info"
          :key="item.uid"
        >
          <v-card-text>
            <p class="display-1 text--primary">
                {{ item.name }}
            </p>
            <p>Age</p>
            <div class="text--primary">
                {{ item.age }}
            </div>
          </v-card-text>
          <v-card-actions>
              <router-link :to="{ name: 'Buyer_info', params: { id : item.uid }}">
                <v-btn text color="deep-purple accent-4">
                  Learn More
                </v-btn>
            </router-link>
          </v-card-actions>
        </v-card>
    </v-card>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from "axios";

export default {
  name: "Home",
  data() {
    return {
      info: null
    };
  },
  mounted() {
    axios
      .get("http://localhost:3000/buyers")
      .then(response => {
        console.log(response);
        this.info = response.data.Buyers;
      })
      .catch(err => console.log(err));
  }
};
</script>

<style> 
    .margin {
        margin: 16px;
    }
</style>
