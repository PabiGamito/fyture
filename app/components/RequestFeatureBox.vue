<template lang="pug">
div
  .feature-request-box
    .title Suggest a feature
    .description Let us know what you want to see in our app.

    form(@submit="submitRequest")
      .text-input(v-on:click="focusInput")
        label Title
        input(type="text", name="title", placeholder="Short, descriptive title", required=true, autocomplete="off", v-model="title").title
      .textarea(v-on:click="focusInput")
        label Details
        textarea-autosize(name="details", placeholder="Any additional details", required=true, autocomplete="off", v-model="details", :min-height="60", :max-height="200").details

      button(type="submit") Suggest Feature
  
  .powered-by
    a(href="" target="_blank") Powered by Fyture
</template>

<script>
export default {
  name: "RequestFeatureBox",
  data() {
    return { title: null, details: null };
  },
  methods: {
    focusInput(e) {
      const input = e.currentTarget.querySelector("input, textarea");
      input.focus();
    },
    async submitRequest(e) {
      e.preventDefault();

      const response = await fetch(`/api/feature`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          title: this.title,
          details: this.details
        })
      });

      if (response.ok) {
        const data = await response.json();
        this.$router.push({ path: `/feature/${data.id}` });
      } else {
        if (response.status == 401) {
          // Unauthorized, promt user to login
          this.$modal.show("login");

          const callback = () => {
            this.$root.$off("loggedIn", callback);
            this.$modal.hide("login");

            // Resubmit
            this.submitRequest(e);
          };

          this.$root.$on("loggedIn", callback);
        }
      }

      // TODO: Handle failure
    }
  }
};
</script>

<style lang="sass" scoped>
.feature-request-box
  margin-top: 30px
  font-family: 'Montserrat', sans-serif
  padding: 10px 15px 15px 15px
  background-color: rgb(250, 250, 250)
  border-radius: 10px
  border: 1px solid #efefef

  .title
    font-weight: 500
    color: #333
    font-size: 16px
    margin: 15px 0
    text-align: center

  .description
    margin: 0 0 15px
    font-weight: 300
    font-size: 14px
    text-align: center

  form
    margin: 0

  .text-input, .textarea
    border: 1px solid #efefef
    margin: 0 0 15px
    background: #fff
    cursor: text
    box-sizing: border-box
    border-radius: 5px
    display: flex
    flex-direction: column
    padding: 8px 12px

    label
      margin: 0 0 5px
      color: #999
      font-size: 11px
      font-weight: 600
      line-height: 15px
      text-transform: uppercase
      cursor: text

    input, textarea
      width: 100%
      border: 0
      margin: 0
      padding: 0
      resize: none
      outline: none
      color: #333
      background: none
      font-size: 14px
      font-weight: 400
      text-align: left
      font-family: 'Montserrat', sans-serif

      &::placeholder
        color: #999

    textarea
      height: 60px

  button
    font-family: 'Montserrat', sans-serif
    background: rgb(255, 223, 176)
    cursor: pointer
    margin: 0
    padding: 12px 15px
    border: none
    outline: none
    // text-transform: uppercase
    font-weight: 600
    border-radius: 5px
    color: rgba(0, 0, 0, 0.75)
    font-size: 11px

    &:hover
      background: rgb(255, 223, 176, 0.8)

.powered-by
  font-family: 'Montserrat', sans-serif
  color: #888
  padding: 15px
  text-align: center
  font-size: 11px

  a
    text-decoration: none
    color: #888

    &:hover
      color: #666
</style>
