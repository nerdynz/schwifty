<template>
  <div class="select">
    <select :value="value" :name="name" @input="onInput" @change="onChange">
      <optgroup label="Lead">
        <option value="New">
          New
        </option>
        <option value="Hot">
          Hot
        </option>
        <option value="On Hold">
          On Hold
        </option>
      </optgroup>
      <optgroup label="Quoting">
        <option value="Quoted">
          Quoted
        </option>
        <option value="Ballpark">
          Ballpark
        </option>
        <option value="Accepted">
          Accepted
        </option>
        <option value="Rejected">
          Rejected
        </option>
      </optgroup>
      <optgroup label="Working">
        <option value="In Progress">
          In Progress
        </option>
        <option value="Delayed">
          Delayed
        </option>
        <option value="With Client">
          With Client
        </option>
      </optgroup>
      <optgroup label="Finished">
        <option value="Needs Review">
          Needs Review
        </option>
        <option value="To Bill">
          To Bill
        </option>
        <option value="Invoiced">
          Invoiced
        </option>
      </optgroup>
      <optgroup label="Archive">
        <option value="Archived">
          Archived
        </option>
        <option value="No Charge">
          No Charge
        </option>
      </optgroup>
    </select>
  </div>
</template>

<script>
// ... imports

export default {
  // COMPONENT
  // ______________________________________
  name: 'JobStatus',
  components: {},
  props: {
    name: String,
    value: String
  },
  data () {
    return {}
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
    onInput (ev) {
      let val = ev.target.value
      this.$emit('input', val)
    },
    onChange (ev) {
      let val = ev.target.value
      this.$emit('change', val)
    }
  }
}
</script>
