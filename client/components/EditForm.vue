<template>
  <div class="edit-form">
    <slot name="title">
      <div class="edit-form-header">
        <h2 class="edit-form-title" v-if="title">
          {{ title }}
        </h2>
      </div>
    </slot>
    <div class="">
      <slot />
    </div>
  </div>
</template>

<script>
// ... imports

export default {
  // COMPONENT
  // ______________________________________
  components: {},
  props: {
    title: String
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
