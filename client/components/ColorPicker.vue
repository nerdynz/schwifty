<template>
  <div class="custom-color" :style="simple ? 'display:inline-block' : ''">
    <span v-if="simple" class="swatch" :style="{'background-color': color.hex }" @click="show" />
    <div v-else class="input" @click="show">
      <span class="swatch" :style="{'background-color': color.hex}" />
      &nbsp;{{ color.hex }}
    </div>
    <div class="color-picker-outer" v-show="isShowing">
      <div class="is-overlay is-fixed" @click="hide" />
      <color-picker disable-alpha :value="color" @input="changeColor" />
    </div>
  </div>
</template>

<script>
import { Chrome } from 'vue-color'

export default {
  components: {
    'color-picker': Chrome
  },
  props: {
    value: String,
    simple: Boolean,
    default: {
      type: String,
      default: '#FFFFFF'
    }
  },
  data () {
    return {
      isShowing: false,
      eventListener: null
    }
  },
  computed: {
    color () {
      if (this.value) {
        return {
          hex: this.value
        }
      }
      return {
        hex: this.default
      }
    }
  },
  // LIFECYCLE METHODS
  // ______________________________________
  beforeCreate () {
  },
  created () {
    let self = this
    //  eslint-disable-next-line
    this.eventListener = document.addEventListener('keyup', function (evt) {
      if (evt.keyCode === 27) {
        self.hide()
      }
    })
  },
  destroyed () {
    document.removeEventListener('keyup', this.eventListener)
  },
  beforeMount () {
  },
  mounted () {
    if (!this.value) {
      this.$emit('input', this.default)
    }
  },
  beforeUpdate () {
  },
  updated () {
  },
  beforeDestroy () {
  },
  methods: {
    show () {
      this.isShowing = !this.isShowing
    },
    hide (ev) {
      if (
        ev.target.className.includes('is-overlay') ||
        ev.target.className.includes('vc-saturation-circle') ||
        ev.target.className.includes('vc-saturation--black') ||
        ev.target.className.includes('vc-saturation--white')
      ) {
        this.isShowing = false
      }
      this.$emit('blur')
    },
    changeColor (colors) {
      // clearTimeout(this.hideTimer)
      // this.hideTimer = setTimeout(() => {
      //   this.isShowing = false
      // }, 800)
      this.$emit('input', colors.hex)
    }
  }
}
</script>

<style lang="scss">
@import "~assets/variables";

.custom-color {
  position: relative;
  min-width: 120px;
  .color-picker-outer {
    position: absolute;
    .is-overlay {
      z-index: 5000;
    }
  }
  .vc-chrome {
    position: relative;
    background: $grey-lighter;
    margin-top: 2px;
    position: relative;
    box-shadow: none;
    z-index: 5001;

    &:after, &:before {
      margin-bottom: 1px;
      bottom: 100%;
      left: 15px;
      border: solid transparent;
      content: " ";
      height: 0;
      width: 0;
      position: absolute;
      pointer-events: none;
    }

    &:after {
      border-color: rgba(136, 183, 213, 0);
      border-bottom-color: $white;
      border-width: 5px;
      margin-left: -5px;
    }

    &:before {
      border-color: rgba(194, 225, 245, 0);
      border-bottom-color:  $grey-light;
      border-width: 6px;
      margin-left: -6px;
    }

    .vc-chrome-body {
      background: $grey-lighter;
      border: 1px solid $grey-light;
    }
    .vc-chrome-fields-wrap {
      .vc-chrome-toggle-btn {
        display: none;
      }
    }
  }
  .swatch {
    width: 16px;
    height: 16px;
    border: 1px solid #000;
    border-radius: 2px;
    display: inline-block;
  }
}

</style>
