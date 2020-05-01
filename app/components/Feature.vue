<template lang="pug">
.feature
  Content(mainComponent="FeatureThread" sideComponent="RequestFeatureBox")
</template>

<script>
import Content from "./Content";

export default {
  name: "Feature",
  components: {
    Content
  }
};
</script>

<style lang="sass" scoped>
</style>
