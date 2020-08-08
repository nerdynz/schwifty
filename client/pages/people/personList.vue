<template>
  <list-placeholder @create="create" :showing="hasData" :loading="isLoading">
    <b-table
      ref="personList"
      custom-key="ulid"
      :data="personData.people"
      :striped="true"
      :mobile-cards="true"
      :paginated="true"
      :per-page="limit"
      :backend-pagination="true"
      :backend-sorting="true"
      :total="personData.total"
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
        <b-table-column field="name" label="Name">
          {{ props.row.name }}
        </b-table-column>
        <b-table-column field="email" label="Email">
          {{ props.row.email }}
        </b-table-column>
        <b-table-column field="username" label="Username">
          {{ props.row.username }}
        </b-table-column>
        <b-table-column field="initials" label="Initials">
          {{ props.row.initials }}
        </b-table-column>
        <b-table-column field="date_created" label="Date Created">
          {{ fmtDate(props.row.date_created) }}
        </b-table-column>
        <b-table-column field="date_modified" label="Date Modified">
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
import gql from 'graphql-tag'
import { mapActions } from 'vuex'
import ListPlaceholder from '~/components/layout/ListPlaceholder'

export default {
  apollo: {
    personData: {
      variables () {
        return {
          siteID: this.siteID,
          offset: this.offset,
          limit: this.limit
        }
      },
      update ({ people, total }) {
        let count = total.aggregate.count || 0
        let result = {
          people,
          total: count
        }
        return result
      },
      query: gql`query people($siteID: Int, $offset: Int, $limit: Int){
        people: person(where: { site_id: {_eq: $siteID }}, offset: $offset, limit: $limit) {
          date_created
          date_modified
          email
          initials
          name
          password
          phone
          picture
          role
          site_id
          ulid
        },
        total: person_aggregate(where: { site_id: {_eq: $siteID }}) {
          aggregate {
            count
          }
        }
      }`

    }
  },
  components: {
    ListPlaceholder
  },
  data () {
    return {
      page: 1,
      limit: 3,
      siteID: 1,
      personData: {
        people: [],
        total: 0
      }
    }
  },
  computed: {
    isLoading () {
      return this.$apollo.queries.personData.loading
    },
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
    this.setButtons(this.buttons)
  },
  methods: {
    ...mapActions({
      setButtons: 'app/setButtons'
    }),

    // sort (field, direction) {
    //   let pagedData = this.pagedData
    //   this.load(field, direction, pagedData.limit, pagedData.pageNum)
    // },

    pageChange (page) {
      this.page = page
    },

    create () {
      this.$router.push({ name: 'person-ID-personEdit', params: { ID: 0 } })
    },

    edit (record) {
      this.$router.push({ name: 'person-ID-personEdit', params: { ID: record.PersonID } })
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
