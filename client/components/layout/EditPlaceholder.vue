<template>
  <div class="edit-page">
    <slot />
    <div class="footer-push" />
    <footer-bar :buttons="buttons" />
    <b-loading :is-full-page="true" :active="isStillLoading" />
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
import FooterBar from '~/components/layout/FooterBar'

export default {
  // COMPONENT
  // ______________________________________
  components: {
    FooterBar
  },
  props: {
    loading: Boolean,
    buttons: {
      type: Array
    }
  },
  data () {
    return {
    }
  },
  computed: {
    ...mapGetters({
      isLoading: 'app/isLoading'
    }),
    isStillLoading () {
      if (this.loading) {
        return true
      }
      if (this.isLoading) {
        return true
      }
      return false
    }
  },
  watch: {},
  // LIFECYCLE METHODS
  // ______________________________________
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
  // METHODS
  // ______________________________________
  methods: {
  }
}
</script>
