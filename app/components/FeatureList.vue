<template lang="pug">
.features
  .feature(v-for="(feature, index) in features")
    FeatureListItem(v-bind:title="feature.title", v-bind:upvotes="feature.upvotes", v-bind:id="feature.id", v-bind:details="feature.details")
</template>

<script>
import FeatureListItem from "./FeatureListItem";

export default {
  name: "FeatureList",
  props: ["sort"],
  data() {
    return { features: [] };
  },
  watch: {
    async sort(newVal, oldVal) {
      if (newVal != oldVal) {
        this.features = await this.fetchFeatures(newVal);
      }
    }
  },
  async created() {
    this.features = await this.fetchFeatures(this.sort);
  },
  components: { FeatureListItem },
  methods: {
    async fetchFeatures(sort, limit = 10, offset = 0) {
      const response = await fetch(
        `/api/features?sort=${sort}&limit=${limit}&offset=${offset}`,
        {
          method: "GET",
          headers: {
            Accept: "application/json"
          }
        }
      );
      if (!response.ok) {
        // TODO: Handle
        return [];
      }

      const json = await response.json();
      return json.features;
    }
  }
};
</script>

<style lang="sass" scoped>
.features
  padding: 0 30px
</style>
