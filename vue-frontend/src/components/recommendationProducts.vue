<template>
    <v-card class="ma-4">
        <v-card-text> 
             <p class="display-1 text--primary">
                A few Recommendations
             </p>
        </v-card-text>
        
        <v-card v-for="product in this.Recommendations" :key="product.uid" class="ma-4">
            <v-card-text class="products"> 
                <p>
                    {{ product.name }} at {{ product.price }}$
                </p>
            </v-card-text>
        </v-card>
    </v-card>
</template>


<script>
export default {
    name: "RecommendedProducts",
    props: {
        Recommendations: [Object]
    },
    mounted() {
        console.log(this.Recommendations);
    }
}
</script>

<style scope> 

.products {
    text-align: center;
    font-weight: bold;
}

</style>
