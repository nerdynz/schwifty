<template>
  <div class="list-placeholder u-rel">
    <div v-show="!isStillLoading">
      <div v-if="showing">
        <slot />
      </div>
      <div v-else class="no-data">
        {{ noData }}
      </div>
    </div>
    <footer-bar :buttons="buttons" />
    <b-loading :active="isStillLoading" :is-full-page="true" />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import FooterBar from '~/components/layout/FooterBar'

export default {
  // COMPONENT
  // ______________________________________
  name: 'ListPlaceholder',
  components: {
    FooterBar
  },
  props: {
    noData: {
      type: String,
      default: 'No data'
    },
    showing: Boolean,
    loading: Boolean,
    buttons: {
      type: Array
    }
  },
  data () {
    return {
      hello: true
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
  methods: {
    add () {
      this.$emit('create')
    }
  }
}
</script>

<style lang="scss">
  .list-placeholder {
    min-height: 100%;
  }
</style>
