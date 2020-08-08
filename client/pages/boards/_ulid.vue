<template>
  <boards />
</template>
<script>
import Boards from './boards.vue'
export default {
  components: {
    Boards
  },
  computed: {
  },
  beforeCreate () {
    this.$store.commit('general/CURRENT_SELECTED_BOARD', this.$route.params.ulid)
  }
}
</script>
