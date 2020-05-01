<template lang="pug">
.content
  .side-container
    div(v-sticky="true", sticky-offset="{top: 0}")
      component(v-bind:is="sideComponent", v-bind="sideComponentProps")
  .main-container
    component(v-bind:is="mainComponent", v-bind="mainComponentProps")
</template>

<script>
export default {
  name: "Content",
  props: [
    "sideComponent",
    "mainComponent",
    "sideComponentProps",
    "mainComponentProps"
  ]
};
</script>

<style lang="sass" scoped>
.content
  display: flex
  align-self: center
  justify-content: center

  .side-container
    width: 300px

  .main-container
    width: 630px
</style>
