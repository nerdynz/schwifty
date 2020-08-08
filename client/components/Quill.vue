<template>
  <div class="editor-wrap" :class="{ 'is-focused': focused }">
    <div ref="editor" class="quill-editor" />
  </div>
</template>

<script>
import Quill from 'quill'
export default {
  props: {
    value: {
      type: String
    },
    options: {
      type: Object,
      default: () => ({})
    }
  },
  data () {
    return {
      focused: false,
      editor: null
    }
  },
  watch: {
    value () {
      if (!this.editor.hasFocus()) {
        this.editor.root.innerHTML = this.value
      }
    }
  },
  mounted () {
    let rootEl = this.$refs.editor
    if (!this.editor) {
      this.editor = new Quill(rootEl, this.options)
      this.editor.on('text-change', () => {
        if (this.editor.root.innerHTML !== this.value) {
          this.$emit('input', this.editor.root.innerHTML)
        }
      })
      this.editor.on('editor-change', (ev, a, b, c) => {
        if (ev === 'selection-change') {
          let newRange = a
          let oldRange = b
          if (newRange && !oldRange) {
            this.focused = true
          }
          if (!newRange && oldRange) {
            this.focused = false
          }
        }
      })
      this.editor.root.innerHTML = this.value
    }
  }
}
</script>

<style lang="styl">
@import '~quill/assets/core';
</style>
