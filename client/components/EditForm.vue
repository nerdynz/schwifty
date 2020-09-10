<template>
  <div class="edit-form" :class="{'is-small': size == 'small'}">
    <slot name="title">
      <div class="edit-form-header">
        <h2 class="edit-form-title" v-if="title">
          {{ title }}
        </h2>
      </div>
    </slot>
    <div class="edit-form-body"><slot /></div>
  </div>
</template>

<script>
// ... imports

export default {
  // COMPONENT
  // ______________________________________
  components: {},
  props: {
    title: String,
    size: String,
    noMargin: Boolean
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

<style lang="scss">
@import "~assets/variables";

.edit-form {
  .edit-form-header {
    // margin: 1.5rem 0 1rem;
    padding: 0.75rem 1rem 0.65rem;
    background-color: $input-border-color;
    display: flex;
    align-items: center;
    &.has-no-top-margin {
      margin-top: 0;
    }
  }

  .edit-form-title {
    font-size: $size-large;
    font-weight: $weight-bold;

    &.is-small {
    font-size: $size-small;
    }
  }

  &:first-of-type {
    .edit-form-header {
      margin-top: 0;
    }
  }

  &.is-small {
    .edit-form-title {
      font-size: $size-small;
    }
  }
}
</style>
