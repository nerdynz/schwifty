<template>
  <list-placeholder :buttons="buttons" @create="create" :showing="hasData">
    <b-table
      ref="personList"
      custom-key="ulid"
      :data="pagedData.records"
      :striped="true"
      :mobile-cards="true"
      :paginated="true"
      :per-page="limit"
      :backend-pagination="true"
      :backend-sorting="true"
      :total="pagedData.total"
      sort="sort"
      @page-change="pageChange"
      class="is-vertical-centered"
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
        <b-table-column field="Email" label="Email">
          {{ props.row.Email }}
        </b-table-column>
        <b-table-column field="Username" label="Username">
          {{ props.row.Username }}
        </b-table-column>
        <b-table-column field="Initials" label="Initials">
          {{ props.row.Initials }}
        </b-table-column>
        <b-table-column field="DateCreated" label="Date Created">
          {{ fmtDate(props.row.DateCreated) }}
        </b-table-column>
        <b-table-column field="DateModified" label="Date Modified">
          {{ fmtDate(props.row.DateModified) }}
        </b-table-column>
        <b-table-column field="role" label="Role">
          {{ props.row.role }}
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
    hasData () {
      if (this.personData && this.personData.people && this.personData.people.length > 0) {
        return true
      }
      return true
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
    this.load()
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
      this.$service.paged('person', sort, direction, limit, pageNum).then((data) => {
        this.pagedData = data
      })
    },

    create () {
      this.$router.push({ name: 'person-ulid-personEdit', params: { ulid: 'new' } })
    },

    edit (record) {
      this.$router.push({ name: 'person-ulid-personEdit', params: { ulid: record.PersonULID } })
    }
  }
}

/* ROUTE DEFINITION
  'people': {
   title: 'Person',
   icon: 'fa-circle-o',
   children: {
     'people-ID-personEdit': {
       title: function (instance) {
         return instance.$route.params.ID === 0 ? 'Create Person' : 'Edit Person'
       }
     }
   }
  }
*/
</script>
