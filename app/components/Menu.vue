<template lang="pug">
.menu
  .search
    .text-input(v-on:click="focusInput")
      label S
      input(type="text", name="search", placeholder="Search...", autocomplete="off")
    
  .feature-menu
    .tabs
      .tab(v-for="tab in tabs", v-bind:class="{active : tab.active}", v-on:click="setActiveTab(tab)")
        | {{ tab.name }}
</template>

<script>
export default {
  name: "Menu",
  props: ["items", "activeItem", "onTabChange"],
  data: function() {
    const tabs = [];
    for (const item of this.items) {
      tabs.push({ name: item, active: this.activeItem == item });
    }

    return {
      tabs: tabs
    };
  },
  methods: {
    focusInput: function(e) {
      const input = e.currentTarget.querySelector("input, textarea");
      input.focus();
    },
    setActiveTab: function(tab) {
      if (tab.active) {
        return;
      }

      for (const tab of this.tabs) {
        tab.active = false;
      }
      tab.active = true;

      this.onTabChange(tab.name);
    }
  }
};
</script>

<style lang="sass" scoped>
.menu
  display: flex
  align-self: center
  align-items: center
  justify-content: center
  font-family: 'Montserrat', sans-serif
  font-weight: 600
  font-size: 13px
  background: white
  // border-bottom: 2px rgb(235, 235, 235) solid

  .feature-menu
    width: 630px // TODO: Needs to match content width

    .tabs
      padding: 0 30px // TODO: Needs to match content padding
      display: flex

      .tab
        margin: 8px
        color: #aaa
        padding: 8px 15px
        border-radius: 3px
        cursor: pointer

        &:first-child
          margin-left: 0

        &.active
          color: #666
          background: rgb(237, 237, 237)

  .search
    margin: 15px 0
    width: 300px // TODO: Needs to match sidebar width

    .text-input
      border: 1px solid #efefef
      background: #fff
      cursor: text
      box-sizing: border-box
      border-radius: 5px
      display: flex
      flex-direction: row
      align-items: center

      label
        color: #999
        font-size: 11px
        font-weight: 600
        line-height: 15px
        text-transform: uppercase
        display: inline-block
        padding: 10px 15px
        color: #bbb
        cursor: pointer

      input
        width: 100%
        border: 0
        margin: 0
        padding: 0
        resize: none
        outline: none
        color: #333
        background: none
        font-size: 14px
        font-weight: 500
        text-align: left
        font-family: 'Montserrat', sans-serif

        &::placeholder
          color: #999
</style>
