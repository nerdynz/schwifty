<template>
  <div id="task-detail" class="modal-card">
    <header class="modal-card-head">
      <title-input v-model="task.Title" />
    </header>
    <section class="modal-card-body">
      <div class="columns is-gapless u-fh-cascade">
        <div class="column is-8 ">
          <div class="is-flex-column">
            <edit-form title="Description" size="small"><textarea class="description" v-model="task.Notes" rows="5" /></edit-form>
            <comments :task="task" />
          </div>
        </div>
      </div>
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
import Comments from '~/components/Comments'

export default {
  // COMPONENT
  // ______________________________________
  components: {
    TitleInput,
    Comments
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
    this.$service.retrieve('~Task', this.taskUlid).then((task) => {
      this.task = task
    })
  },
  beforeUpdate () {
  },
  updated () {
  },
  beforeDestroy () {
    this.$service.save('~Task', this.task)
  },
  destroyed () {
  },
  // METHODS
  // ______________________________________
  methods: {
  }
}
</script>
