]<template>
  <div v-if="record">
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <p class="modal-card-title">
          {{ type }} Contact
        </p>
      </header>
      <section class="modal-card-body">
        <div class="columns is-multiline">
          <div class="column is-12">
            <control label="Name" name="Name" type="text" v-model="record.Name" v-validate="'required'" />
          </div>
          <div class="column is-12">
            <control label="Email" name="Email" type="text" v-model="record.Email" v-validate="'required|email'" />
          </div>
          <div class="column is-12">
            <control label="Phone" name="Phone" type="text" v-model="record.Phone" v-validate="'required'" />
          </div>
        </div>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-primary" @click="change">
          Add
        </button>
        <button class="button" type="button" @click="$parent.close()">
          Cancel
        </button>
      </footer>
    </div>
  </div>
</template>

<script>
export default {
  components: {
  },
  inject: {
    $validator: '$validator'
  },
  props: {
    record: Object,
    type: String
  },
  data () {
    return {
    }
  },
  computed: {
    isNewRecord () {
      if (this.record.ContactID <= 0) {
        return true
      }
      return false
    }
  },
  methods: {
    change () {
      this.$validator.validateAll(['Name', 'Email', 'Phone']).then((valid) => {
        if (valid) {
          this.$emit('change', this.record)
          this.$parent.close()
        }
      })
    }
  }
}
</script>
