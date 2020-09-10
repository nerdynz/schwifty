<template>
  <div class="edit-form-header">
    <h2 class="edit-form-title" :class="{ 'is-small': size === 'small' }">
      {{ title }}
    </h2>
  </div>
</template>

<script>
// ... imports

export default {
  // COMPONENT
  // ______________________________________
  components: {},
  props: {
    title: {
      type: String,
      required: true
    },
    size: String
  },
  data () {
    return {
    }
  },
  computed: {
  },
  watch: {
  },
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
