<template>
  <footer class="footer">
    <div class="footer-buttons level">
      <div class="level-left">
        <div v-for="(item, i) in leftButtons" :key="i" class="level-item">
          <button class="button" :class="'is-' + item.kind" :type="item.kind" :disabled="item.isDisabled || !item.click" @click="(ev) => { item.click(ev) }">
            {{ item.text }}
          </button>
        </div>
      </div>
      <div v-for="(item, i) in centerButtons" :key="i" class="level-item">
        <button class="button" :class="'is-' + item.kind" :type="item.kind" :disabled="item.isDisabled || !item.click" @click="(ev) => { item.click(ev) }">
          {{ item.text }}
        </button>
      </div>
      <div class="level-right">
        <div v-for="(item, i) in rightButtons" :key="i" class="level-item">
          <button class="button" :class="'is-' + item.kind" :type="item.kind" :disabled="item.isDisabled || !item.click" @click="(ev) => { item.click(ev) }">
            {{ item.text }}
          </button>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  props: {
    buttons: {
      type: Array,
      default () {
        return []
      }
    }
  },
  data () {
    return {
    }
  },
  computed: {
    leftButtons (state) {
      return this.buttons
        .filter(button => button.alignment === 'left' && (typeof (button.role) === 'undefined' || button.role === state.login.details.Person.Role))
    },
    rightButtons (state) {
      return this.buttons
        .filter(button => button.alignment === 'right' && (typeof (button.role) === 'undefined' || button.role === state.login.details.Person.Role))
    },
    centerButtons (state) {
      return this.buttons
        .filter(button => button.alignment === 'center' && (typeof (button.role) === 'undefined' || button.role === state.login.details.Person.Role))
    },
    specialButtons (state) {
      return this.buttons
        .filter(button => button.alignment === 'special' && (typeof (button.role) === 'undefined' || button.role === state.login.details.Person.Role))
    }
  }
}
</script>
