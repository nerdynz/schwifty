<template>
  <div>
    <multiselect
      v-if="options"
      :disabled="disabled"
      :value="selectedVal"
      :options="options"
      @input="input"
      :track-by="valueKey"
      :label="valueLabel"
      :placeholder="placeholder"
    />
    <div class="loader" v-else />
  </div>
</template>

<script>
import { byField } from '~/helpers/filters'

export default {
  // COMPONENT
  // ______________________________________
  components: {
  },
  props: {
    label: String,
    name: String,
    placeholder: String,
    value: {
      type: [String, Number]
    },
    url: {
      type: String,
      required: true
    },
    optionsKey: {
      type: String
    },
    valueKey: {
      type: String
    },
    valueLabel: {
      type: String
    },
    disabled: Boolean,
    obj: Boolean,
    validation: Object
  },
  data () {
    return {
      options: [],
      isLoaded: false
    }
  },
  computed: {
    selectedVal () {
      return byField(this.options, this.valueKey, this.value)
    }
  },
  // LIFECYCLE METHODS
  // ______________________________________
  beforeCreate () {
  },
  created () {
    this.load()
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
  methods: {
    load () {
      if (this.url) {
        this.$axios.get(this.url).then(({ data }) => {
          if (this.optionsKey && Object.prototype.hasOwnProperty.call(data, this.optionsKey)) {
            this.options = data[this.optionsKey]
          } else {
            // console.log(data)
            this.options = data
          }
          this.$nextTick(() => {
            this.isLoaded = true
          })
        })
      }
    },
    input (obj) {
      if (!obj) {
        // clearred....
        this.$emit('input', null)
        return
      }
      if (this.isLoaded) {
        if (this.obj) {
          this.$emit('input', obj)
        } else if (Object.prototype.hasOwnProperty.call(obj, this.valueKey)) {
          this.$emit('input', obj[this.valueKey])
        } else {
          throw new Error('Couldn\'t find those options. Bad value-key')
        }
      }
    }
  }
}
</script>
