<template>
    <v-card class="ma-4 mt-8 py-3">
        <v-card-text class="ma-4">
            <p class="display-1 text--primary">
                Historial de Compras:
            </p>
        </v-card-text>

        <v-row justify="center" v-for="(purchase, index) in purchases" :key="purchase.uid" class="ma-2">
            <v-expansion-panels popout>
                <v-expansion-panel>
                  <v-expansion-panel-header>{{index + ". " + formatDate(purchase.Date)}} </v-expansion-panel-header>
            <v-expansion-panel-content>
                <p><strong>Device:</strong> {{purchase.device }} </p>
                <p><strong>Ip:</strong> {{purchase.ip }} </p>
                <h2> Products: </h2>
                <v-simple-table>
                    <template v-slot:default>
                      <thead>
                        <tr>
                          <th class="text-left">Name</th>
                          <th class="text-left">Price</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="item in purchase.Products" :key="item.uid">
                          <td>{{ item.name }}</td>
                          <td>{{ item.price }}</td>
                        </tr>
                      </tbody>
                    </template>
                  </v-simple-table>
            </v-expansion-panel-content>
          </v-expansion-panel>
        </v-expansion-panels>
      </v-row>
    </v-card>
</template>

<script>
export default {
    name: "History",
    props: {
        purchases: Object
    },
    methods: {
        getTotal( purchases ) {
            let Total = 0;
            for(let i = 0; i < purchases.length;i++){
                Total += purchases[i].price;
            }
            return Total;
        },
        formatDate( current ) {
            console.log(current);
            const date = new Date( current );
            return date.getDate()+'-' + (date.getMonth()+1) + '-'+date.getFullYear();
        }
    }
}
</script>

<style>
.main-text {
    text-align: center;
    font-weight: bold;
    margin-top: 50px;
}
</style>
