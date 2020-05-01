// TODO: If user already voted for it disable voting and color the upvote arrow

<template lang="pug">
.feature
  .body
    router-link(:to="{ name: 'feature', params: { id: id }}")
      .title
        span {{ title }}
        span.id {{ " #" + id }}
      .details 
        v-clamp(autoresize, :max-lines="2") {{ details }}
  .votes
    .upvotes(v-on:click="upvote")
      .triangle
      span {{ upvotes }}
</template>

<script>
import VClamp from "vue-clamp";

export default {
  name: "FeatureListItem",
  props: ["title", "upvotes", "id", "details"],
  components: {
    VClamp
  },
  data: () => {
    return { upvoted: false };
  },
  methods: {
    async upvote(e) {
      if (this.upvoted) {
        return;
      }

      this.upvoted = true;
      this.upvotes++;

      const response = await fetch(`/api/feature/upvote/${this.id}`, {
        method: "POST"
      });

      if (!response.ok) {
        this.upvoted = false;
        this.upvotes--;
      }
    }
  }
};
</script>

<style lang="sass" scoped>
.feature
  font-family: 'Montserrat', sans-serif
  display: flex
  padding: 18px 0

  .body
    width: 100%

    a
      color: inherit
      text-decoration: inherit

    .title
      font-family: 'Montserrat', sans-serif
      font-weight: 600
      font-size: 15px
      color: #333
      margin-bottom: 8px

      span.id
        color: #888
        font-weight: 400
        font-size: 14px

    .details
      font-family: 'Montserrat', sans-serif
      font-size: 14px
      color: #666
      width: 100%
      text-align: justify

  .upvotes
    display: flex
    align-items: center
    margin: 5px 0 5px 30px
    border: 2px rgb(232, 232, 232) solid
    border-radius: 5px
    padding: 8px 20px
    cursor: pointer
    box-shadow: 0px 1px 5px rgba(0,0,0,0.01)

    .triangle,
    .triangle:before,
    .triangle:after
      width: 4px
      height: 4px
      border-top-right-radius: 50%

    .triangle
      position: relative
      background-color: rgb(125, 125, 125)
      text-align: left
      margin-right: 8px
      margin-bottom: 1px
      transform: rotate(-60deg) skewX(-30deg) scale(1,.866)
      &:before
        transform: rotate(-135deg) skewX(-45deg) scale(1.414,.707) translate(0,-50%)
      &:after
        transform: rotate(135deg) skewY(-45deg) scale(.707,1.414) translate(50%)
      &:before, &:after
        content: ''
        position: absolute
        background-color: inherit

    span
      font-family: 'Roboto', sans-serif
      font-size: 13px
      color: #333
      font-weight: 600

    &:hover
      background: rgb(248, 251, 255)
      box-shadow: 0px 1px 5px rgba(0,0,0,0.05)

      .triangle
        background-color: #333
</style>
