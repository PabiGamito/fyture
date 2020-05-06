<template lang="pug">
.app
  LoginModal
  Header(appName="My App")
  .content
    transition(name="fade" mode="out-in" :duration="{ enter: 300, leave: 300 }" appear)
      router-view

</template>

<script>
import Vue from "vue";
import VueRouter from "vue-router";
import Sticky from "vue-sticky-directive";
import TextareaAutosize from "vue-textarea-autosize";

import Header from "./Header.vue";
import Home from "./Home.vue";
import Feature from "./Feature.vue";
import LoginModal from "./LoginModal.vue";

Vue.use(VueRouter);
Vue.use(Sticky);
Vue.use(TextareaAutosize);

import FeatureList from "./FeatureList";
import FeatureThread from "./FeatureThread";
import RequestFeatureBox from "./RequestFeatureBox";

Vue.component("FeatureList", FeatureList);
Vue.component("FeatureThread", FeatureThread);
Vue.component("RequestFeatureBox", RequestFeatureBox);

// Routes
const routes = [
  {
    path: "/",
    component: Home
  },
  {
    path: "/feature/:id",
    component: Feature,
    name: "feature"
  },
  // not found
  {
    path: "*",
    redirect: "/"
  }
];

export default {
  name: "App",
  router: new VueRouter({ routes, mode: "history" }),
  components: { Header, LoginModal }
};
</script>
