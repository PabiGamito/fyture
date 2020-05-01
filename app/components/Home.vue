<template lang="pug">
.home
  Menu(v-bind:items="tabs", v-bind:onTabChange="handleSortChange", v-bind:activeItem="activeTab")
  Content(mainComponent="FeatureList", v-bind:mainComponentProps="{'sort': sortBy}", sideComponent="RequestFeatureBox")
</template>

<script>
import Content from "./Content";
import Menu from "./Menu";

export default {
  name: "Home",
  components: {
    Content,
    Menu
  },
  data() {
    const sortKeyToTabNames = {
      trending: "Trending",
      top: "Most votes",
      new: "Newest"
    };

    const tabNameToSortKey = {};
    for (const key in sortKeyToTabNames) {
      if (sortKeyToTabNames.hasOwnProperty(key)) {
        tabNameToSortKey[sortKeyToTabNames[key]] = key;
      }
    }

    const tabs = Object.keys(tabNameToSortKey);
    const sortBy = this.$route.query.sort || tabNameToSortKey[tabs[0]];

    const activeTab = sortKeyToTabNames[sortBy];

    return {
      tabNameToSortKey: tabNameToSortKey,
      tabs: tabs,
      activeTab: activeTab,
      sortBy: sortBy
    };
  },
  methods: {
    handleSortChange(newTab) {
      this.sortBy = this.tabNameToSortKey[newTab];
    }
  }
};
</script>

<style lang="sass" scoped>
</style>
