<template>
  <div class="application">
    <navbar />
    <div class="navbar-push" />
    <main class="main">
      <nuxt />
    </main>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Navbar from '~/components/layout/Navbar'

export default {
  components: {
    Navbar
  },
  props: {
  },
  data () {
    return {
      // settings: []
    }
  },
  computed: {
    ...mapGetters({
      user: 'app/user',
      siteID: 'auth/siteID'
    })
  },
  watch: {},
  beforeCreate () {
  },
  created () {
  },
  beforeMount () {
  },
  mounted () {
  },
  beforeUpdate () {
  },
  updated () {
  },
  beforeDestroy () {
  },
  destroyed () {
  },
  methods: {
  }
}
</script>

<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>
<style lang="scss">
  @import "~assets/main";
</style>
