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
  props: {
    label: String,
    name: String,
    placeholder: String,
    value: {
      type: [String, Number]
    },
    gql: {
      type: Object,
      required: true
    },
    collectionKey: {
      type: String,
      required: true
    },
    valueKey: {
      type: String,
      required: true
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
    this.$apollo.addSmartQuery('lookup', {
      query: this.gql,
      manual: true,
      result (result) {
        console.log('eh', result.data)
        if (this.collectionKey in result.data) {
          this.options = result.data[this.collectionKey]
          this.isLoaded = true
        } else {
          console.error('Failed to find collection key')
        }
      }
    })

    console.log(this.$apollo)
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
    input (obj) {
      if (!obj) {
        // clearred....
        this.$emit('input', null)
        return
      }
      if (this.isLoaded) {
        if (this.obj) {
          console.log(obj)
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
