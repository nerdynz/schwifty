<template>
  <boards />
</template>

<script>
import { mapGetters } from 'vuex'
import Boards from './boards.vue'
export default {
  components: {
    Boards
  },
  computed: {
    ...mapGetters({
      currentBoardULID: 'general/currentBoardULID'
    })
  },
  created () {
    if (this.currentBoardULID) {
      this.$router.push({ name: 'boards-ulid', params: { ulid: this.currentBoardULID } })
    }
  }
}
</script>
