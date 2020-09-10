<template>
    <div>
        <h1>Historial de Compras</h1>
        <v-row justify="center" v-for="purchases in history" :key="purchases.uid">
            <v-expansion-panels popout>
              <v-expansion-panel
                v-for="(item,i) in 5"
                :key="i"
              >
              <v-expansion-panel-header>{{purchases.ip + " " + purchases.device}} </v-expansion-panel-header>
            <v-expansion-panel-content>
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
                        <tr v-for="item in purchases.Products" :key="item.uid">
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
      {{this.history }}
    </div>
</template>

<script>
export default {
    name: "History",
    props: {
        history: Object
    }
}
</script>

