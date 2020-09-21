<template>
  <div class="home">
    <v-card-title class="justify-center">
        <h1>Buyers</h1>
    </v-card-title>
    <div class="text-center">
        <v-pagination
          :length="length"
          v-on:next="updateVisibles()"
          v-on:prev="updateVisibles()"
          v-on:input="updateVisibles()"
          v-model="current_page"
          circle
        ></v-pagination>
      </div>
    <v-card
      class="d-flex  align-content-space-around justify-center flex-wrap"
      color="grey lighten-2"
      flat
      tile
      min-height="400"
      v-if="!this.loading"
    >
        <v-card
          class="pa-2 margin"
          max-width="344"
          v-for="item in visibles"
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
    <div v-else class="text-center">
        <v-progress-circular
          indeterminate
          color="primary"
        >loading</v-progress-circular>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Home",
  data() {
    return {
        info: null,
        loading: false,
        current_page: 1,
        page_size: 15,
        visibles: [],
        length: 0
    };
  },
  mounted() {
      this.loading = true;
    axios
      .get("http://localhost:3000/buyers")
      .then(response => {
          this.loading = false;
          this.info = response.data.Buyers;
          this.length = Math.ceil(this.info.length / this.page_size);
          this.updateVisibles();
      })
      .catch(err =>{
          this.loading = false;
          console.log(err)
      });
  },
    methods: {
        updateVisibles() {
            this.visibles = this.info.slice((this.current_page-1) * this.page_size, this.current_page * this.page_size);
        }
    }
};
</script>

<style> 
    .margin {
        margin: 16px;
    }
</style>
