<template>
  <div class="field" :class="{'has-addons': hasAddons }">
    <label v-if="label" :for="name" class="label thin-label">
      <span v-if="isRequired" class="required-star">*</span>
      {{ label }}
    </label>
    <slot />
  </div>
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
