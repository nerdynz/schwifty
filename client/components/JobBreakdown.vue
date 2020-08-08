<template>
  <div class="job-breakdown" :class="depthClass">
    <div class="field is-vertical is-combined u-mb">
      <div class="field has-addons -u-m">
        <div class="control u-fw">
          <input v-label.float label="Name / Description" type="text" v-model="milestone.Title" class="input is-top u-b">
        </div>
        <div class="control">
          <button :disabled="milestone.Depth === 0" class="button is-primary has-icon" @click="downDepth">
            <b-icon icon="chevron-left" size="is-small" />
          </button>
        </div>
        <div class="control">
          <button :disabled="milestone.Depth === 2" class="button is-primary has-icon" @click="upDepth">
            <b-icon icon="chevron-right" size="is-small" />
          </button>
        </div>
      </div>
      <textarea rows="4" v-model="milestone.Description" class="textarea is-bottom" />
    </div>
  </div>
</template>
<script>
export default {
  // COMPONENT
  // ______________________________________
  components: {
  },
  props: {
    milestone: Object
  },
  data () {
    return {
    }
  },
  computed: {
    depthClass () {
      return 'depth-' + this.milestone.Depth + ''
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
    upDepth () {
      if (this.milestone.Depth < 2) {
        this.milestone.Depth = this.milestone.Depth + 1
      }
    },
    downDepth () {
      if (this.milestone.Depth > 0) {
        this.milestone.Depth = this.milestone.Depth - 1
      }
    }
  }
}
</script>
