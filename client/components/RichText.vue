<template>
  <div>
    <editor-menu-bar :editor="editor" v-slot="{ commands, isActive }">
      <div>
        <button class="button is-small" :class="{ 'is-active': isActive.bold() }" @click="commands.bold">
          <b-icon icon="far fa-bold" />
        </button>
      </div>
    </editor-menu-bar>
    <editor-content :editor="editor" class="textarea" />
  </div>
</template>

<script>
import { EditorMenuBar, Editor, EditorContent } from 'tiptap'
import {
  Bold
} from 'tiptap-extensions'

export default {
  // COMPONENT
  // ______________________________________
  components: {
    EditorMenuBar,
    EditorContent
  },
  props: {
  },
  data () {
    return {
      editor: new Editor({
        extensions: [
          new Bold()
        ],
        content: `
          <h1>Yay Headlines!</h1>
          <p>All these <strong>cool tags</strong> are working now.</p>
        `
      })
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
  },
  beforeUpdate () {
  },
  updated () {
  },
  beforeDestroy () {
    this.editor.destroy()
  },
  destroyed () {
  },
  // METHODS
  // ______________________________________
  methods: {
  }
}
</script>
