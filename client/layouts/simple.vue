<template>
  <main class="simple main-body">
    <nuxt />
  </main>
</template>

<script>
export default {

}
</script>

<style lang="scss" >
  @import "~assets/login";
  @import "~nerdutils/nerdutils";

  .simple.main-body {
    background: $sidebar-color;
    .logo {
      padding: 2rem 4rem;
    }
  }
</style>
