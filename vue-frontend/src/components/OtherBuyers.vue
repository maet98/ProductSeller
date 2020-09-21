<template>
    <v-card class="size">
        <v-card-text class="elem">
            <p class="display-1 text--primary">
                Others Buyers
            </p>
          <v-card
            class="mx-auto size"
            v-for="buyer in buyers"
            :key="buyer.uid"
            v-on:click="goToBuyer(buyer.uid)"
          >
            <v-card-text>
              <p class="text--primary">
                {{ buyer.name }}
              </p>
            </v-card-text>
          </v-card>
        </v-card-text>
    </v-card>
</template>

<script>


export default {
    name: "OtherBuyer",
    props: {
        buyers: [Object],
        changeId: Function,
    },
    methods: {
        goToBuyer( id ) {
            console.log(id);
            this.$router.push({ name: 'Buyer_info', params: { id: id} })
            this.changeId( id );
        }
    }
}
</script>


<style>
.size {
    margin: 8px;
    text-align: center;
}

.elem {
    margin: 40;
    margin-top: 50px;
    font-weight: bold;
}

</style>
