<template>
  <div id="task-detail" class="modal-card">
    <header class="modal-card-head">
      <title-input v-model="task.Title" />
    </header>
    <section class="modal-card-body">
      asdf
    </section>
    <footer class="modal-card-foot">
      <button class="button is-info">
        asdf
      </button>
    </footer>
  </div>
</template>
<script>
import TitleInput from '~/components/TitleInput'
export default {
  // COMPONENT
  // ______________________________________
  components: {
    TitleInput
  },
  props: {
    taskUlid: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      task: {}
    }
  },
  computed: {
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
    this.$service.retrieve('~task', this.taskUlid).then((task) => {
      this.task = task
    })
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
