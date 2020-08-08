<template>
  <div>
    <label v-if="label" :for="label" class="label is-thin">
      {{ label }}
    </label>
    <div class="u-rel">
      <input
        @click.stop="click"
        @select.stop="focus"
        ref="input"
        :name="label"
        class="input title-input is-editable"
        type="text"
        :value="value"
        @input="input"
        @blur="blur"
        :style="{'width': width + 'px' }"
      >
      <input disabled ref="ghost" class="title-input ghost" :value="value">
    </div>
  </div>
</template>

<script>

export default {
  name: 'TitleInput',
  components: {

  },
  props: {
    label: String,
    value: String
  },
  data () {
    return {
      width: 250
    }
  },
  mounted () {
    this.adjustTitleWidth()
  },
  methods: {
    input (ev) {
      this.adjustTitleWidth(ev.target)
      this.$emit('input', ev.target.value)
    },
    click () {
    },
    blur (ev) {
      this.$emit('blur', this.value)
    },
    focus () {
      this.$refs.input.focus()
    },
    select (ev) {
      this.$refs.input.select()
    },
    adjustTitleWidth () {
      let input = this.$refs.input
      if (input) {
        input.style.width = 0
        let oldWidth = this.width
        let newWidth = input.scrollWidth + 10
        if (newWidth > oldWidth) {
          input.style.width = newWidth + 'px'
          console.log('newWidth', newWidth)
          this.width = newWidth
        } else {
          // this.$nextTick(() => {
          console.log('newWidth', newWidth)
          this.width = newWidth
          // })
        }
      }
    }
  }
}
</script>

<style lang="scss">
  .title-input {
    min-width: 200px !important;
    max-width: 750px !important;
    position: relative;
    z-index: 1;
    &.ghost {
      position: absolute;
      opacity: 0;
      top:0;
      left:0;
      z-index: 0;
      position: absolute;
    }
  }
</style>
