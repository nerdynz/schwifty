<template>
  <list-placeholder @create="create" :showing="hasData" :loading="isLoading">
    <b-table
      ref="boardList"
      datakey="BoardID"
      :data="pagedData.records"
      :striped="true"
      :mobile-cards="true"
      :paginated="true"
      :per-page="pagedData.limit"
      :backend-pagination="true"
      :backend-sorting="true"
      :total="pagedData.total"
      :default-sort="[pagedData.sort, pagedData.direction]"
      @row-drag-dropped="sorting"
      @sort="sort"
      @page-change="pageChange"
    >
      <template slot-scope="props">
        <b-table-column label="">
          <div class="field has-addons">
            <p class="control -u-mb">
              <button class="button is-small is-info" @click="edit(props.row, props.index)">
                View
              </button>
            </p>
            <p class="control -u-mb">
              <button class="button is-small is-success" @click="activate(props.row, props.index)">
                Activate
              </button>
            </p>
            <p class="control -u-mb">
              <button class="button is-small is-danger" @click="(ev) => {remove(props.row, props.index, ev)}">
                Delete
              </button>
            </p>
          </div>
        </b-table-column>

        <b-table-column field="Title" label="Title">
          <strong>{{ props.row.Title }}</strong>
        </b-table-column>

        <b-table-column field="DateModified" label="Last Changed">
          {{ fmtDate(props.row.DateCreated) }}
        </b-table-column>

        <b-table-column field="IsActive" label="Is Active">
          <span class="tag is-small" :class="props.row.IsActive ? 'is-success' : 'is-warning'">{{ props.row.IsActive ? 'Active' : 'Archived' }}</span>
        </b-table-column>
      </template>
    </b-table>
  </list-placeholder>
</template>

<script>
import { mapActions } from 'vuex'
import ListPlaceholder from '~/components/layout/ListPlaceholder'
import { fmtDate } from '~/helpers/format'
import { changeSortReverse } from '~/helpers/filters'

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
    osType () {
      return (window.platform.os.family.includes('OS X') ? 'cmd' : 'ctrl') + ''
    },
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
    this.load('DateModified', 'desc', 20, 1)
    this.setButtons(this.buttons)
  },
  methods: {
    fmtDate,

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
      this.$service.paged('board', sort, direction, limit, pageNum).then((data) => {
        this.pagedData = data
        this.isLoading = false
      })
    },

    create () {
      this.$router.push({ name: 'boards-ID-boardEdit', params: { ID: 0 } })
    },

    edit (record) {
      this.$router.push({ name: 'boards-ID-boardEdit', params: { ID: record.BoardID } })
    },

    remove (record, index, ev) {
      let deleteImediately = ev.metaKey || ev.ctrlKey
      let doDelete = () => {
        this.$service.remove('board', record).then(() => {
          this.$toast.open({ message: 'Board deleted', type: 'is-danger' })
          setTimeout(() => {
            this.load(this.pagedData.sort, this.pagedData.direction, this.pagedData.limit, this.pagedData.pageNum)
          }, 1)
        })
      }
      if (deleteImediately) {
        doDelete()
      } else {
        this.$dialog.confirm({
          title: 'Changes Not Saved',
          message: `You have not saved your changes. <br>Are you sure you want to continue? <br><br> <small> You can press<code class='u-up'>${this.osType}+click</code> when deleting to bypass this confirmation`,
          confirmButtonText: 'Continue',
          cancelButtonText: 'Cancel',
          type: 'is-danger',
          hasIcon: true,
          onConfirm: doDelete
        })
      }
    },

    sorting (from, to) {
      if (from !== to) {
        let newRecords = changeSortReverse(this.pagedData.records, from, to)
        this.pagedData.records = newRecords
        this.$axios.put('/api/v1/board/sort', newRecords)
      }
    }

  }
}
</script>
