<template>
  <div class="notes-control control content" :class="className">
    <label :for="name" class="label thin-label">{{ label }}</label>
    <quill :name="name" :id="name" :value="value" :options="{ theme: 'snow' }" @input="change" />
    <span class="help is-danger" v-if="error">{{ error }}</span>
  </div>
</template>

<script>
import Quill from './Quill'

export default {
  components: {
    Quill
  },
  props: {
    name: String,
    className: String,
    value: {
      type: String,
      default: ''
    },
    label: String,
    error: String
  },
  data () {
    return {
      isSelected: false,
      password: ''
    }
  },
  computed: {

  },
  methods: {
    change (val) {
      this.$emit('input', val)
    }
  }
}
</script>

<style lang="styl">
@import "~quill/assets/snow";

.ql-toolbar {
  border-top-left-radius: 3px;
  border-top-right-radius: 3px;
}

.quill-editor{
  border-bottom-left-radius: 3px;
  border-bottom-right-radius: 3px;
}

.quill-editor.ql-container .ql-editor {
  padding: .6rem;
}

.notes-control{
  position: relative;
  z-index: 10;
}
</style>
