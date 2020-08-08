<template>
  <div v-if="record" class="u-scr ">
    <div v-if="!record.IsActive" class="notification is-warning u-m u-up">
      <strong>This board is archived</strong>
      <button class="button u-fr is-small" @click="activate">
        Activate
      </button>
    </div>
    <board :board="record" @board-save="save" />
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import Board from '~/components/Board.vue'
export default {
  components: {
    Board
  },
  inject: {
    $validator: '$validator'
  },
  data () {
    return {
      record: null
    }
  },
  computed: {
    isNewRecord () {
      if (this.routeID <= 0) {
        return true
      }
      return false
    },
    routeID () {
      return this.$route.params.ID ? parseInt(this.$route.params.ID) : 0
    },
    buttons () {
      return [
        { text: 'Activate', alignment: 'left', kind: 'success', click: this.activate },
        { text: 'Back', alignment: 'left', kind: 'text', click: this.back },
        { text: 'Delete', alignment: 'right', kind: 'danger', click: this.remove }
      ]
    },
    ...mapGetters({
      user: 'app/user'
    })
  },
  watch: {
    'errors.items' () {
      this.setButtons(this.buttons)
    }
  },
  beforeCreate () {
  },
  created () {
    this.loadRecord()
    this.setButtons(this.buttons)
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
    ...mapActions({
      setButtons: 'app/setButtons'
    }),

    loadRecord () {
      let loader = this.$loading.open()
      if (this.isNewRecord) {
        this.record = this.$service.new('board')
        loader.close()
      } else {
        this.$service.retrieve('board', this.routeID).then((newRecord) => {
          this.record = newRecord
          loader.close()
        })
      }
    },

    save (goBack) {
      if (this.errors.any()) {
        this.$toast.open({ title: 'Failed to save!', message: 'Please double check the fields highlighted.', type: 'is-danger' })
        return
      }

      this.setButtons(this.buttons)
      this.$service.save('board', this.record).then((newRecord) => {
        this.record = newRecord
        this.$toast.open({ message: 'Board saved successfully', type: 'is-success' })
        if (goBack === true) {
          this.back()
          return
        } else if (this.routeID === 0) {
          this.$router.replace({ name: 'board-boardID', params: { boardID: newRecord.BoardID } })
        }
        this.setButtons(this.buttons)
      })
    },

    activate () {
      this.record.IsActive = true
      this.save()
    },

    back (ev) {
      this.$router.go(-1)
    },

    remove () {
      this.$service.remove('board', this.record).then(() => {
        this.$toast.open({ message: 'Board removed', type: 'is-danger' })
        this.back()
      })
    }
  }
}
</script>
