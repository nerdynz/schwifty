<template>
  <list-placeholder @create="create" :showing="hasData" :loading="isLoading" :buttons="buttons">
    <b-table
      ref="clientList"
      custom-key="ClientID"
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
        <!-- <b-table-column field="Address" label="Address">
          {{ props.row.Address }}
        </b-table-column> -->
        <b-table-column field="Rate" label="Rate">
          {{ props.row.Rate }}
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
import ListPlaceholder from '~/components/layout/ListPlaceholder'

export default {
  components: {
    ListPlaceholder
  },
  data () {
    return {
      isLoading: true,
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
      this.$service.paged('client', sort, direction, limit, pageNum).then((data) => {
        this.pagedData = data
        this.isLoading = false
      })
    },

    create () {
      this.$router.push({ name: 'clients-ulid-clientEdit', params: { ulid: 'new' } })
    },

    edit (record) {
      this.$router.push({ name: 'clients-ulid-clientEdit', params: { ulid: record.ClientULID } })
    }
  }
}

</script>
