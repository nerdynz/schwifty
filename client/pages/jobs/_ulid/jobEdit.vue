<template>
  <edit-placeholder :buttons="buttons" :is-loading="isLoading">
    <edit-form title="Details" v-if="job">
      <div class="columns is-multiline">
        <div class="column is-6">
          <field for="ClientID">
            <lookup
              label="Client"
              v-label.float
              v-model="job.ClientULID"
              name="ClientULID"
              url="/api/v1/client/retrieve"
              value-key="ClientULID"
              value-label="Name"
            />
          </field>
        </div>
        <div class="column is-6" />
        <div class="column is-4">
          <field for="Name">
            <input
              label="Name"
              name="Name"
              class="input"
              type="text"
              v-model="job.Name"
              v-label.float
            >
          </field>
        </div>
        <div class="column is-4">
          <field for="Status">
            <job-status label="Status" name="Status" v-model="job.Status" v-label.float.stuck />
          </field>
        </div>
        <div class="column is-4">
          <field for="DueDate">
            <!-- <datepicker label="Status" name="Status" v-model="job.DueDate" v-label.float.stuck /> -->
          </field>
        </div>
        <div class="column is-12">
          <field for="Notes" label="Notes">
            <notes v-model="job.Notes" />
          </field>
        </div>
      </div>
    </edit-form>
    <edit-form title="Breakdown" v-if="job">
      <template #title>
        <div class="edit-form-header">
          <span class="edit-form-title">
            Breakdown
          </span>
          <div class="u-plr">
            <button class="button is-primary is-small" @click="addMilestone">
              Add
            </button>
          </div>
        </div>
      </template>
      <job-breakdown v-for="(milestone) in job.Milestones" :key="milestone.MilestoneULID" :milestone="milestone" />
    </edit-form>
  </edit-placeholder>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import Lookup from '~/components/Lookup'
import EditPlaceholder from '~/components/layout/EditPlaceholder'
import Notes from '~/components/Notes'
import JobStatus from '~/components/JobStatus'
import JobBreakdown from '~/components/JobBreakdown'

export default {
  components: {
    EditPlaceholder,
    Lookup,
    Notes,
    JobStatus,
    JobBreakdown
  },
  data () {
    return {
      job: null
    }
  },
  computed: {
    ...mapGetters({
      siteID: 'auth/siteID',
      isLoading: 'app/isLoading'
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
    },
    ...mapGetters({
      user: 'app/user',
      siteID: 'auth/siteID'
    })
  },
  watch: {
  },
  beforeCreate () {
  },
  created () {
    if (this.isNewRecord) {
      this.job = this.$service.new('job', this.siteID)
    } else {
      this.load()
    }
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

    addMilestone () {
      this.job.Milestones = this.job.Milestones || []
      let milestone = this.$service.new('Milestone', this.siteID)
      milestone.JobULID = this.job.JobULID
      this.job.Milestones.push(milestone)
    },

    load () {
      this.$service.retrieve('job', this.routeID, 'milestones', true).then((j) => {
        this.job = j
      })
    },

    save () {
      this.$service.save('job', this.job, 'milestones', true).then((record) => {
        this.job = record
        // this.$router.push({ name: 'jobs-ulid-jobEdit', params: { ulid: record.JobULID } })
      })
    },

    back (ev) {
      this.$router.go(-1)
    },

    remove () {
      this.$service.remove('job', this.job).then(() => {
        this.$snackbar.open({ message: 'Job removed', type: 'is-danger' })
        this.back()
      })
    }
  }
}
</script>
