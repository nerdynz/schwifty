<template>
  <list-placeholder @create="create" :showing="hasData" :buttons="buttons">
    <b-table
      ref="jobList"
      custom-key="JobID"
      :data="pagedData.records"
      :striped="true"
      :mobile-cards="true"
      :paginated="true"
      :per-page="pagedData.limit"
      :backend-pagination="true"
      :backend-sorting="true"
      :total="pagedData.total"
      :default-sort="[pagedData.sort, pagedData.direction]"
      @sort="sort"
      @page-change="pageChange"
    >
      <template slot-scope="props">
        <b-table-column label="">
          <div class="field has-addons">
            <p class="control -u-mb">
              <button class="button is-small" @click="edit(props.row, props.index)">
                Edit
              </button>
            </p>
          </div>
        </b-table-column>

        <b-table-column field="Name" label="Name">
          {{ props.row.Name }}
        </b-table-column>
        <b-table-column field="QuoteNotes" label="Quote Notes">
          {{ props.row.QuoteNotes }}
        </b-table-column>
        <b-table-column field="Notes" label="Notes">
          {{ props.row.Notes }}
        </b-table-column>
        <b-table-column field="Status" label="Status">
          {{ props.row.Status }}
        </b-table-column>
        <b-table-column field="DateCreated" label="Date Created">
          {{ fmtDate(props.row.DateCreated) }}
        </b-table-column>
        <b-table-column field="DateModified" label="Date Modified">
          {{ fmtDate(props.row.DateModified) }}
        </b-table-column>
      </template>
    </b-table>
  </list-placeholder>
</template>
<script>
import { mapActions } from 'vuex'
import ListPlaceholder from '~/components/layout/ListPlaceholder'

export default {
  components: {
    ListPlaceholder
  },
  data () {
    return {
      pagedData: {
        sort: '',
        direction: 'desc',
        records: [],
        total: 0,
        pageNum: 1,
        limit: 50
      }
    }
  },
  computed: {
    hasData () {
      if (this.pagedData && this.pagedData.records && this.pagedData.records.length > 0) {
        return true
      }
      return false
    },
    buttons () {
      return [
        { text: 'Add New', alignment: 'left', kind: 'success', click: this.create }
      ]
    }
  },
  created () {
    this.load('DateModified', 'desc', 50, 1)
  },
  methods: {
    ...mapActions({
      setButtons: 'app/setButtons'
    }),

    sort (field, direction) {
      let pagedData = this.pagedData
      this.load(field, direction, pagedData.limit, pagedData.pageNum)
    },

    pageChange (page) {
      let pagedData = this.pagedData
      this.load(pagedData.sort, pagedData.direction, pagedData.limit, page)
    },

    load (sort, direction, limit, pageNum) {
      this.isLoading = true
      this.$service.paged('job', sort, direction, limit, pageNum).then((data) => {
        this.pagedData = data
        this.isLoading = false
      })
    },

    create () {
      this.$router.push({ name: 'jobs-ulid-jobEdit', params: { ulid: 'new' } })
    },

    edit (record) {
      this.$router.push({ name: 'jobs-ulid-jobEdit', params: { ulid: record.JobULID } })
    }
  }
}

/* ROUTE DEFINITION
  'jobs': {
   title: 'Job',
   icon: 'fa-circle-o',
   children: {
     'jobs-ID-jobEdit': {
       title: function (instance) {
         return instance.$route.params.ID === 0 ? 'Create Job' : 'Edit Job'
       }
     }
   }
  }
*/
</script>
