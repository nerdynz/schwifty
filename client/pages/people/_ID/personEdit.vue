<template>
  <div v-if="record">
    <div class="columns is-multiline">
      <div class="column is-12">
        <control label="Person Id" name="PersonID" type="number" v-model="record.PersonID" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="First Name" name="FirstName" type="text" v-model="record.FirstName" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="Last Name" name="LastName" type="text" v-model="record.LastName" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="Email" name="Email" type="text" v-model="record.Email" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="Password" name="Password" type="text" v-model="record.Password" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="Username" name="Username" type="text" v-model="record.Username" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="Initials" name="Initials" type="text" v-model="record.Initials" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="Date Created" name="DateCreated" type="datetime" v-model="record.DateCreated" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="Date Modified" name="DateModified" type="datetime" v-model="record.DateModified" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="Ulid" name="Ulid" type="text" v-model="record.Ulid" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="Site Id" name="SiteID" type="number" v-model="record.SiteID" x-validate="'required'" />
      </div>
      <div class="column is-12">
        <control label="Role" name="Role" type="text" v-model="record.Role" x-validate="'required'" />
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  components: {
  },
  inject: {
    $validator: '$validator'
  },
  data () {
    return {
      record: null
    }
  },
  computed: {
    isNewRecord () {
      if (this.routeID <= 0) {
        return true
      }
      return false
    },
    routeID () {
      return this.$route.params.ID ? parseInt(this.$route.params.ID) : 0
    },
    buttons () {
      let buttons = [
        { text: 'Save', alignment: 'left', kind: 'success', click: this.save, isDisabled: this.$validator.errors.any() },
        { text: 'Save & Return', alignment: 'left', kind: 'success', click: () => { this.save(true) }, isDisabled: this.$validator.errors.any() },
        { text: 'Back', alignment: 'left', kind: '', click: this.back }
      ]
      if (this.isNewRecord) {

      } else {
        // cant delete a new record
        buttons.push({ text: 'Delete', alignment: 'right', kind: 'danger', click: this.remove })
      }
      return buttons
    },
    ...mapGetters({
      user: 'app/user'
    })
  },
  watch: {
    'errors.items' () {
      this.setButtons(this.buttons)
    }
  },
  beforeCreate () {
  },
  created () {
    this.loadRecord()
    this.setButtons(this.buttons)
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
    ...mapActions({
      setButtons: 'app/setButtons'
    }),

    loadRecord () {
      let loader = this.$loading.open()
      if (this.isNewRecord) {
        this.record = this.$service.new('person')
        loader.close()
      } else {
        this.$service.retrieve('person', this.routeID).then((newRecord) => {
          this.record = newRecord
          loader.close()
        })
      }
    },

    save (goBack) {
      let promise = new Promise((resolve, reject) => {
        this.$validator.validateAll().then((valid) => {
          if (valid) {
            this.$service.save('person', this.record).then((newRecord) => {
              this.record = newRecord
              this.$snackbar.open({ position: 'is-top', message: 'Person saved successfully', type: 'is-success' })
              if (this.routeID === 0) {
                this.$router.replace({ name: 'people-ID-personEdit', params: { ID: newRecord.PersonID } })
              }
              resolve(newRecord)
            })
          } else {
            this.$snackbar.open({ position: 'is-top', title: 'Failed to save!', message: 'Please double check the fields highlighted.', type: 'is-danger' })
          }
        })
      })
      return promise
    },

    back (ev) {
      this.$router.go(-1)
    },

    remove () {
      this.$service.remove('person', this.record).then(() => {
        this.$toast.open({ message: 'Person removed', type: 'is-danger' })
        this.back()
      })
    }
  }
}
</script>
