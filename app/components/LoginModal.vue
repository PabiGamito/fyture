<template lang="pug">
  modal(name="login", transition="pop-out", :focus-trap="true", width="300px", height="auto")
    .login-modal
      .content
        .header
          | Sign in to Fyture with:
          font-awesome-icon(:icon="['fab', 'font-awesome']")
        .login-options
          .auth-button
            .icon.google(v-on:click="launchAuthWindows('/auth/google')")
              img(src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4NCjwhLS0gR2VuZXJhdG9yOiBBZG9iZSBJbGx1c3RyYXRvciAxOC4wLjAsIFNWRyBFeHBvcnQgUGx1Zy1JbiAuIFNWRyBWZXJzaW9uOiA2LjAwIEJ1aWxkIDApICAtLT4NCjwhRE9DVFlQRSBzdmcgUFVCTElDICItLy9XM0MvL0RURCBTVkcgMS4xLy9FTiIgImh0dHA6Ly93d3cudzMub3JnL0dyYXBoaWNzL1NWRy8xLjEvRFREL3N2ZzExLmR0ZCI+DQo8c3ZnIHZlcnNpb249IjEuMSIgaWQ9IkxheWVyXzEiIHhtbG5zOnNrZXRjaD0iaHR0cDovL3d3dy5ib2hlbWlhbmNvZGluZy5jb20vc2tldGNoL25zIg0KCSB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4PSIwcHgiIHk9IjBweCIgdmlld0JveD0iMCAwIDQwIDQwIg0KCSBlbmFibGUtYmFja2dyb3VuZD0ibmV3IDAgMCA0MCA0MCIgeG1sOnNwYWNlPSJwcmVzZXJ2ZSI+DQo8dGl0bGU+YnRuX2dvb2dsZV9kYXJrX2ZvY3VzX2lvczwvdGl0bGU+DQo8ZGVzYz5DcmVhdGVkIHdpdGggU2tldGNoLjwvZGVzYz4NCjxnIGlkPSJHb29nbGUtQnV0dG9uIiBza2V0Y2g6dHlwZT0iTVNQYWdlIj4NCgk8ZyBpZD0iX3gzOV8tUEFUQ0giIHRyYW5zZm9ybT0idHJhbnNsYXRlKC02NjguMDAwMDAwLCAtMjE5LjAwMDAwMCkiIHNrZXRjaDp0eXBlPSJNU0FydGJvYXJkR3JvdXAiPg0KCTwvZz4NCgk8ZyBpZD0iYnRuX2dvb2dsZV9kYXJrX2ZvY3VzIiB0cmFuc2Zvcm09InRyYW5zbGF0ZSgtMS4wMDAwMDAsIC0xLjAwMDAwMCkiIHNrZXRjaDp0eXBlPSJNU0FydGJvYXJkR3JvdXAiPg0KCQk8ZyBpZD0ibG9nb19nb29nbGVnXzQ4ZHAiIHRyYW5zZm9ybT0idHJhbnNsYXRlKDE1LjAwMDAwMCwgMTUuMDAwMDAwKSIgc2tldGNoOnR5cGU9Ik1TTGF5ZXJHcm91cCI+DQoJCQk8cGF0aCBpZD0iU2hhcGUiIHNrZXRjaDp0eXBlPSJNU1NoYXBlR3JvdXAiIGZpbGw9IiM0Mjg1RjQiIGQ9Ik0yMC43LDYuM2MwLTEuMS0wLjEtMi4xLTAuMy0zLjFINi4zdjUuOGg4LjENCgkJCQljLTAuMywxLjktMS40LDMuNS0zLDQuNXYzLjhoNC44QzE5LjEsMTQuOCwyMC43LDEwLjksMjAuNyw2LjNMMjAuNyw2LjN6Ii8+DQoJCQk8cGF0aCBpZD0iU2hhcGVfMV8iIHNrZXRjaDp0eXBlPSJNU1NoYXBlR3JvdXAiIGZpbGw9IiMzNEE4NTMiIGQ9Ik02LjMsMjFjNC4xLDAsNy40LTEuMyw5LjktMy42bC00LjgtMy44QzEwLDE0LjUsOC4zLDE1LDYuMywxNQ0KCQkJCWMtMy45LDAtNy4yLTIuNi04LjQtNi4yaC01djMuOUMtNC42LDE3LjYsMC40LDIxLDYuMywyMUw2LjMsMjF6Ii8+DQoJCQk8cGF0aCBpZD0iU2hhcGVfMl8iIHNrZXRjaDp0eXBlPSJNU1NoYXBlR3JvdXAiIGZpbGw9IiNGQkJDMDUiIGQ9Ik0tMi4xLDguOUMtMi40LDgtMi42LDctMi42LDZzMC4yLTIsMC41LTIuOXYtMy45aC01DQoJCQkJYy0xLDItMS42LDQuMy0xLjYsNi43czAuNiw0LjcsMS42LDYuN0wtMi4xLDguOUwtMi4xLDguOXoiLz4NCgkJCTxwYXRoIGlkPSJTaGFwZV8zXyIgc2tldGNoOnR5cGU9Ik1TU2hhcGVHcm91cCIgZmlsbD0iI0VBNDMzNSIgZD0iTTYuMy0zYzIuMiwwLDQuMiwwLjgsNS43LDIuMmw0LjMtNC4zYy0yLjYtMi40LTYtMy45LTEwLTMuOQ0KCQkJCUMwLjQtOS00LjYtNS42LTcuMS0wLjdsNSwzLjlDLTAuOS0wLjQsMi40LTMsNi4zLTNMNi4zLTN6Ii8+DQoJCTwvZz4NCgkJPGcgaWQ9ImhhbmRsZXNfc3F1YXJlIiBza2V0Y2g6dHlwZT0iTVNMYXllckdyb3VwIj4NCgkJPC9nPg0KCTwvZz4NCjwvZz4NCjwvc3ZnPg0K")
          
          .auth-button(v-on:click="launchAuthWindows('/auth/facebook')")
            .icon.facebook
              font-awesome-icon(:icon="['fab', 'facebook-f']")

          .auth-button(v-on:click="launchAuthWindows('/auth/twitter')")
            .icon.twitter
              font-awesome-icon(:icon="['fab', 'twitter']")

          .auth-button(v-on:click="launchAuthWindows('/auth/github')")
            .icon.github
              font-awesome-icon(:icon="['fab', 'github']")
</template>

<script>
export default {
  name: "LoginModal",
  methods: {
    launchAuthWindows(authURL) {
      const authWindow = window.open(authURL);

      const timer = setInterval(() => {
        if (authWindow.closed) {
          clearInterval(timer);
          this.$root.$emit("loggedIn", {});
        }
      }, 1000);
    }
  }
};
</script>

<style lang="sass" scoped>
.login-modal
  padding: 25px

  .content
    display: flex
    align-content: center
    align-items: center
    flex-direction: column
    font-family: 'Montserrat', sans-serif

    .login-options
      display: flex
      flex-direction: row

      .auth-button
        padding: 10px
        border-radius: 5px
        cursor: pointer
        margin: 5px

        &:hover
          background-color: #ddd

        .icon
          width: 30px
          height: 30px
          box-sizing: border-box

          &.github
            color: #333
            padding: 2px

          &.twitter
            color: #55acee
            padding: 2px

          &.facebook
            color: #3b5998
            padding: 3px

          svg
            width: 100%
            height: 100%
            margin: auto
</style>