<template lang="pug">
.feature-thread
  .feature
    .header
      .title 
        | {{ feature.title }} 
        span.id {{ '#' + id }}
    .details {{ feature.details }}
  .thread
    input(placeholder="Leave a comment")
</template>

<script>
export default {
  name: "FeatureThread",
  data() {
    return {
      id: this.$route.params.id,
      feature: {}
    };
  },
  async created() {
    const response = await fetch(`/api/feature/${this.id}`, {
      method: "GET",
      headers: {
        Accept: "application/json"
      }
    });

    if (response.ok) {
      const data = await response.json();
      this.feature = data.feature;
    }

    // TODO: Handle failure
  }
};
</script>

<style lang="sass" scoped>
.feature-thread
  font-family: 'Montserrat', sans-serif
  padding: 30px

  .title
    font-size: 20px
    color: #333
    font-weight: 500

    .id
      color: #888
      font-weight: 400

  .details
    font-size: 14px
    color: #333
</style>
