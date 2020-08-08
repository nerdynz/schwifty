<template>
  <list-placeholder @create="create" :showing="hasData" :loading="isLoading" :buttons="buttons">
    <b-table
      ref="clientList"
      custom-key="ClientID"
      :data="pagedData.records"
      :striped="true"
      :mobile-cards="true"
      :paginated="true"
      :per-page="limit"
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
        <b-table-column field="client_name" label="Name">
          {{ props.row.client_name }}
        </b-table-column>
        <!-- <b-table-column field="Address" label="Address">
          {{ props.row.Address }}
        </b-table-column> -->
        <b-table-column field="Rate" label="Rate">
          {{ props.row.rate }}
        </b-table-column>
        <b-table-column field="date_created" label="Date Created">
          {{ fmtDate(props.row.date_created) }}
        </b-table-column>
        <b-table-column field="date_modified" label="Date Modified">
          {{ fmtDate(props.row.date_modified) }}
        </b-table-column>
      </template>
    </b-table>
  </list-placeholder>
</template>
<script>
import { mapActions } from 'vuex'
import gql from 'graphql-tag'
import ListPlaceholder from '~/components/layout/ListPlaceholder'

export default {
  components: {
    ListPlaceholder
  },
  apollo: {
    pagedData: {
      variables () {
        return {
          siteID: this.siteID,
          offset: this.offset,
          limit: this.limit
        }
      },
      update ({ records, total }) {
        let count = total.aggregate.count || 0
        let result = {
          records,
          total: count
        }
        return result
      },
      query: gql`query jobs($siteID: Int, $offset: Int, $limit: Int){
        records: client(where: { site_id: {_eq: $siteID }}, offset: $offset, limit: $limit) {
          address
          client_name
          contact_name
          date_created
          date_modified
          rate
          site_id
          ulid
        },
        total: client_aggregate(where: { site_id: {_eq: $siteID }}) {
          aggregate {
            count
          }
        }
      }`

    }
  },
  data () {
    return {
      page: 1,
      limit: 3,
      siteID: 1,
      pagedData: {
        records: [],
        total: 0
      }
    }
  },
  computed: {
    isLoading () {
      return this.$apollo.queries.pagedData.loading
    },
    hasData () {
      if (this.pagedData && this.pagedData.records && this.pagedData.records.length > 0) {
        return true
      }
      return false
    },
    offset () {
      return (this.page - 1) * this.limit
    },
    buttons () {
      return [
        { text: 'Add New', alignment: 'left', kind: 'success', click: this.create }
      ]
    }
  },
  created () {
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
      this.page = page
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
