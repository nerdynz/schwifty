<template>
  <edit-placeholder :buttons="buttons" :is-loading="$service.loading">
    <edit-form title="Settings" v-if="record">
      <field for="LogoPicture" label="Logo Picture">
        <file-picker v-model="record.LogoPicture" />
      </field>
      <field for="PrimaryColor" label="Primary Color">
        <color-picker v-model="record.PrimaryColor" />
      </field>
      <field for="SecondaryColor" label="Secondary Color">
        <color-picker v-model="record.SecondaryColor" />
      </field>
    </edit-form>
  </edit-placeholder>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import EditPlaceholder from '~/components/layout/EditPlaceholder.vue'
import ColorPicker from '~/components/ColorPicker.vue'
import FilePicker from '~/components/FilePicker.vue'

export default {
  components: {
    EditPlaceholder,
    ColorPicker,
    FilePicker
  },
  data () {
    return {
      record: null,
      file: null
    }
  },
  computed: {
    buttons () {
      let buttons = [
        { text: 'Save', alignment: 'left', kind: 'success', click: this.save, isDisabled: false },
        { text: 'Save & Return', alignment: 'left', kind: 'success', click: () => { this.save(true) }, isDisabled: false },
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
  created () {
    console.log(this.$nuxt.$loading)
    this.load()
  },
  methods: {
    ...mapActions({
      siteID: 'auth/siteID'
    }),

    load () {
      this.$service.retrieve('setting').then((setting) => {
        this.record = setting
      })
    },

    save () {
      let promise = new Promise((resolve, reject) => {
        // eslint-disable-next-line
        if (true) { // todo: validation here
          this.$service.save('setting', this.record)
            .then((newRecord) => {
              this.record = newRecord
              this.$store.commit('app/SET_SETTINGS', {
                ...newRecord // shallow copy that mofo
              })
              this.$snackbar.open({ position: 'is-top', message: 'Setting saved successfully', type: 'is-success' })
              resolve(newRecord)
            })
        } else {
          this.$snackbar.open({ position: 'is-top', title: 'Failed to save!', message: 'Please double check the fields highlighted.', type: 'is-danger' })
        }
      })
      return promise
    },

    back (ev) {
      this.$router.go(-1)
    },

    remove () {
      this.$service.remove('setting', this.record).then(() => {
        this.$toast.open({ message: 'Setting removed', type: 'is-danger' })
        this.back()
      })
    }
  }
}
</script>
