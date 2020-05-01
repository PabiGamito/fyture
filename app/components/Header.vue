<template lang="pug">
.header
  .content
    h1.appName {{ appName }}
    p.description Let us know how we can improve.<br>
      | Vote on existing ideas or suggest new ones.
</template>

<script>
export default {
  name: "Header",
  props: ["appName"]
};
</script>

<style lang="sass" scoped>
.header
  background: rgb(254, 230, 190)
  background: linear-gradient(to right, rgb(254, 230, 190), rgb(254, 204, 129))
  display: flex
  align-items: center
  align-content: center
  flex-direction: column
  border-bottom: 1px rgb(235, 235, 235) solid

  .content
    font-family: 'Montserrat', sans-serif
    padding: 50px
    width: 930px // TODO: Side bar + content width

    h1.appName
      font-weight: 300
      color: #000
      margin: 0
      padding: 0

    p.description
      color: rgba(60, 60, 60, 0.75)
</style>
