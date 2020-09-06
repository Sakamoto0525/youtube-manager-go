<template>
  <div id="app">
    <TheHeader />

    <div class="columns">
      <TheSidebar />
      <nuxt />
    </div>
  </div>
</template>

<script>
import TheHeader from "~/components/TheHeader";
import TheSidebar from "~/components/TheSidebar";

export default {
  components: {
    TheHeader,
    TheSidebar,
  },
};
</script>

<style>
#app {
  min-height: 100vh;
  background: #fafafa;
}

.columns {
  min-height: 100vh;
}
</style>
