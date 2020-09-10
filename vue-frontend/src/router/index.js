import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Buyer from "../views/Buyer.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/buyer/:id",
    name: "Buyer_info",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Buyer,
      props:true
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
