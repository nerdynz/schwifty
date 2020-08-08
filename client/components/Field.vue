<template>
  <b-field :horizontal="horizontal" :label="label" class="thin-label">
    <slot />
  </b-field>
</template>

<script>
// ... imports

export default {
  // COMPONENT
  // ______________________________________
  name: 'Field',
  components: {},
  props: {
    hasAddons: Boolean,
    horizontal: {
      type: Boolean,
      default: false
    },
    for: {
      type: String,
      required: true
    },
    label: String
  },
  data () {
    return {
      isRequired: false
    }
  },
  computed: {
    name () {
      return this.for
    }
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
