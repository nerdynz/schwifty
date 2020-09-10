<template>
  <edit-placeholder v-if="record">
    <edit-form title="Client">
      <div class="columns">
        <div class="column is-10">
          <field for="Name" label="Name">
            <b-input label="Name" name="Name" type="text" v-model="record.Name" x-validate="'required'" />
          </field>
        </div>
        <div class="column is-2">
          <field for="Rate" label="Default Rate">
            <b-input
              label="Rate"
              name="Rate"
              type="number"
              :value="record.Rate"
              @input="(val) => { record.Rate = Number(val) }"
              x-validate="'required'"
            />
          </field>
        </div>
      </div>
      <field for="Address" label="Address">
        <b-input label="Address" name="Address" type="text" v-model="record.Address" x-validate="'required'" />
      </field>
      <field for="Logo" label="Logo">
        <file-picker v-model="record.Logo" />
      </field>
    </edit-form>
    <edit-form title="Contacts">
      <b-table v-if="record.Contacts" datakey="ULID" :data="filteredContacts">
        <template slot-scope="props">
          <b-table-column field="Name" label="Name">
            {{ props.row.Name }}
          </b-table-column>
          <b-table-column field="Email" label="Email">
            {{ props.row.Email }}
          </b-table-column>
          <b-table-column field="Phone" label="Phone">
            {{ props.row.Phone }}
          </b-table-column>
          <b-table-column field="IsPrimary" label="Primary Contact">
            <span class="tag" :class="props.row.IsPrimary ? 'is-success' : 'is-warning'">
              <b-checkbox :value="props.row.IsPrimary" @click.native.prevent="uncheckOthers(props.row)" />
              {{ props.row.IsPrimary ? 'Yes' : 'No' }}
            </span>
          </b-table-column>
          <b-table-column field="IsDeleted" label="Deleted">
            <button class="button is-danger is-small" @click="props.row.IsDeleted = true"><i class="far fa-trash" /></button>
          </b-table-column>
        </template>
      </b-table>
      <button class="button is-small" @click.prevent="addNewContact">
        Add New
      </button>
    </edit-form>
    <b-modal :active.sync="isContactModalOpen" has-modal-card scroll="keep">
      <contact-modal type="Customer" :record="currentEditContact" @change="updateContact" />
    </b-modal>
  </edit-placeholder>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import { indexByField } from '~/helpers/filters'
import FilePicker from '~/components/FilePicker.vue'
import EditPlaceholder from '~/components/layout/EditPlaceholder.vue'

export default {
  components: {
    FilePicker,
    EditPlaceholder
  },
  data () {
    return {
      record: null
    }
  },
  computed: {
    ...mapGetters({
      user: 'app/user'
    }),
    ...mapGetters({
      siteID: 'auth/siteID'
    }),
    isNewRecord () {
      if (this.routeID === 'new') {
        return true
      }
      return false
    },
    routeID () {
      return this.$route.params.ulid || 'new'
    },
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
    }
  },
  watch: {
  },
  beforeCreate () {
  },
  created () {
    this.loadRecord()
    // this.setButtons(this.buttons)
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
      // let loader = this.$loader.open()
      if (this.isNewRecord) {
        this.record = this.$service.new('client', this.siteID)
        // loader.close()
      } else {
        this.$service.retrieve('client', this.routeID).then((newRecord) => {
          this.record = newRecord
          // loader.close()
        })
      }
    },

    addNewContact () {
      this.openContactModal()
    },

    openContactModal () {
      let contact = this.$service.new('Contact')
      this.currentEditContact = contact
      this.isContactModalOpen = true
    },

    updateContact (contact) {
      let contacts = this.record.Contacts || []
      let index = indexByField(contacts, 'ContactULID', contact.ContactULID)
      if (index >= 0) {
        contacts[index] = contact
      } else {
        contacts.push(contact)
      }
      this.record.Contacts = contacts
    },

    save () {
      let promise = new Promise((resolve, reject) => {
        // eslint-disable-next-line
        if (true) { // todo: validation here
          this.$service.save('client', this.record).then((newRecord) => {
            this.record = newRecord
            this.$snackbar.open({ position: 'is-top', message: 'Client saved successfully', type: 'is-success' })
            if (this.routeID === 0) {
              this.$router.replace({ name: 'clients-ID-clientEdit', params: { ID: newRecord.ClientID } })
            }
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
      this.$service.remove('client', this.record).then(() => {
        this.$toast.open({ message: 'Client removed', type: 'is-danger' })
        this.back()
      })
    }
  }
}
</script>
