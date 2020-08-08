<template>
  <div class="boards">
    <div class="board-tabs">
      <drop
        v-for="(board, index) in boards"
        class="board-tab"
        :class="{'is-active': currentBoardULID === board.BoardULID}"
        :key="index"
        @dragover="tabOver"
        @drop="(xfer, ev) => {tabDrop(xfer, ev, board)}"
      >
        <drag
          class="tab-item-custom"
          @drag="tabDragging"
          @dragend="dragCleanUp"
          :transfer-data="{task: board, index: index, type: 'board', status: ''}"
          :title="board.Title"
          :class="{'has-hover': board.ulid === currentHoverULID, 'is-being-dragged': board.ulid === currentDraggingULID}"
          :ulid="board.BoardULID"
        >
          <div class="board-color-tab" :style="{'background': board.Color}" />
          <a class="is-overlay" @click="changeBoard(board.BoardULID)" />
          <div class="inner">
            {{ fmtMax(board.Title, 20, '...') }}
          </div>
          <!-- <span v-if="currentBoardIndex == index" class="tag is-rounded  is-small" :class="{'is-success': currentBoardTaskCount(board) < 5 , 'is-warning': currentBoardTaskCount(board) >= 5 && currentBoardTaskCount(board) < 10,'is-danger': board.TaskCount > 10 }">
            {{ currentBoardTaskCount(board) }}
          </span>
          <span v-else class="tag is-rounded  is-small" :class="{'is-success': board.TaskCount < 5 , 'is-warning': board.TaskCount >= 5 && board.TaskCount <= 10,'is-danger': board.TaskCount > 10 }">
            {{ board.TaskCount ? board.TaskCount : 0 }}
          </span> -->
        </drag>
      </drop>
      <div class="board-tab spacer-tab">
        <div class="tab-item-custom">
          <div class="inner">
            &nbsp;
          </div>
        </div>
      </div>
    </div>
    <board
      :board="currentBoard"
      :selected-task-ulid="selectedTaskULID"
      @add-task="addTask"
      @board-loaded="boardLoaded"
      @board-updated="updateBoards"
      @board-save="saveBoard"
      class="board u-fh"
      :key="currentBoardULID"
    />
    <b-loading :active="isBoardLoading" />
    <footer-bar :buttons="buttons" />
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import FooterBar from '~/components/layout/FooterBar'
import Board from '~/components/Board.vue'
import { fmtMax, fmtDate } from '~/helpers/format.js'
import { randomHexColor } from '~/helpers/helpers.js'
import { changeSortReverse, getMaxSort, indexByField } from '~/helpers/filters.js'

// import {showMessage, showNotification} from '../../helpers/notification.js'
export default {
  components: {
    FooterBar,
    Board
  },
  props: {
  },
  data () {
    return {
      isBoardLoading: false,
      currentDraggingULID: '',
      currentHoverULID: '',
      tabsPageIndex: 0,
      xpos: 0,
      ypos: 0,
      isComponentModalActive: false,
      selectedTaskULID: '',
      boards: [],
      currentBoard: {
        Tasks: []
      }
    }
  },
  computed: {
    ...mapGetters({
      userID: 'auth/userID',
      siteID: 'auth/siteID',
      usertoken: 'auth/usertoken',
      currentBoardULID: 'general/currentBoardULID'
    }),
    // currentBoard () {
    //   if (this.boards && this.currentBoardULID) {
    //     let boardIndex = indexByField(this.boards, 'BoardULID', this.currentBoardULID)
    //     if (boardIndex === -1) {
    //       return this.boards[boardIndex]
    //     }
    //     return {}
    //   }
    //   return {}
    // },
    buttons () {
      return [
        { text: 'Add New Task', alignment: 'left', kind: 'link', click: this.addTask },
        { text: 'Add New Board', alignment: 'left', kind: 'success', click: this.addBoard },
        { text: 'Archive Board', alignment: 'right', kind: 'warning', click: () => { this.archive() } },
        { text: 'Print Board', alignment: 'right', kind: '', click: () => { window.open(`/board/print/${this.currentBoard.BoardID}?authtoken=${this.usertoken}`, 'print_board', '', true) } }
      ]
    }
  },
  watch: {

  },
  beforeCreate () {
  },
  created () {
    this.load()
    this.loadBoard(this.currentBoardULID)
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
    fmtMax,

    loadBoard (ulid) {
      this.$service.retrieve('board', ulid, 'tasks', true).then((b) => {
        this.currentBoard = b
        this.boardLoaded()
        // this.$emit('board-loaded', b)
      })
    },

    tabDragging (na, ev) {
      if (ev.target && ev.target.children && ev.target.children.length > 0) {
        let blockEl = ev.target.children[0]
        // if (blockEl.className.indexOf('is-being-dragged') < 0) {
        //   blockEl.classList += ' is-being-dragged'
        // }
        let ulid = blockEl.getAttribute('ulid')
        this.currentDraggingULID = ulid
      }
    },

    tabOver (na, ev) {
      if (ev.target) {
        let ulid = ev.target.getAttribute('ulid')
        if (ulid) {
          this.xpos = ev.clientX
          this.ypos = ev.clientY
          this.currentHoverULID = ulid
        }
      }
    },

    tabDrop (xfer, ev, board) {
      // console.log(xfer, ev, board)
      if (xfer && xfer.type && xfer.type === 'task') {
        let task = xfer.task
        task.BoardID = board.BoardID
        task.SortPosition = 10000
        xfer.status = 'moved'
        this.$service.save('task', task)
      } else if (ev.target) {
        let ulid = this.currentHoverULID
        let from = indexByField(this.boardData.records, 'ULID', this.currentDraggingULID)
        let to = indexByField(this.boardData.records, 'ULID', ulid)
        if (from !== -1 && to !== -1) {
          // if (from === (to + 1)) {
          //   console.log('x', from, to)
          // } else if (from > to) {
          //   console.log('y', from, to)
          //   to = to + 1
          // }
          this.sortBoards(from, to)
        }
      }
      this.dragCleanUp()
    },

    dragCleanUp () {
      this.xpos = 0
      this.ypos = 0
      this.currentHoverULID = ''
      this.currentDraggingULID = ''
    },

    load () {
      this.$service.retrieve('board').then((data) => {
        this.boards = data || []
      })
    },

    addBoard () {
      let board = this.$service.new('Board', this.siteID)
      board.Title = 'New board'
      board.Color = randomHexColor()
      this.$service.save('board', board).then(() => {
        this.changeBoard(board.BoardULID)
      })
      // console.log('new board', board)

      // board.forceUpdate = true
      // let max = 0
      // this.boardData.records.forEach((record) => {
      //   if (record.SortPosition > max) {
      //     max = record.SortPosition
      //   }
      // })
      // max += 50
      // board.SortPosition = max
      // board.Color = randomHexColor()
      // this.boardData.records.splice(0, 0, board)
      // this.$nextTick(() => {
      //   this.changeBoard(board.ulid)
      // })
    },

    addTask ({ status, title }) {
      let task = this.$service.new('Task', this.siteID)
      console.log(task)
      task._IsNew = true
      task.BoardULID = this.currentBoard.BoardULID
      if (typeof (status) === 'string') {
        task.Status = status
      } else {
        task.Status = 'Todo'
      }
      if (title) {
        task.Title = title
      } else {
        task.Title = 'A task to do'
      }
      this.currentBoard.Tasks = this.currentBoard.Tasks || []
      this.currentBoard.Tasks.push(task)
      // this.updateBoards(board, this.currentBoardIndex)
      this.selectedTaskULID = task.ulid
      task.SortPosition = getMaxSort(this.currentBoard.Tasks)
    },

    changeBoard (ulid) {
      if (ulid === this.currentBoardULID) {
        return
      }
      // this.$store.commit('app/LOADING_GENERAL', true)
      this.isBoardLoading = true
      console.log('board changed ', ulid)
      // this.currentBoardIndex = index
      // let newULID = this.boardData.records[index].ulid
      // this.$router.push({ name: 'boards-ulid', params: { ulid } })
      this.loadBoard(ulid)
      this.$store.commit('general/CURRENT_SELECTED_BOARD', ulid)
      window.history.replaceState(null, '', `/boards/${ulid}`)
    },

    saveBoard (board) {
      this.$service.save('board', board).then((newBoard) => {
        this.updateBoards(newBoard)
      })
    },

    sortBoards (from, to) {
      let currentULID = this.currentBoard.ulid
      if (from !== to) {
        let newRecords = changeSortReverse(this.boardData.records, from, to)
        this.boardData.records = newRecords
        setTimeout(() => {
          this.$axios.put('/api/v1/board/sort', newRecords)
        })
        this.changeBoard(currentULID)
      }
    },

    boardLoaded () {
      this.isBoardLoading = false
    },

    updateBoards (board) {
      let index = indexByField(this.boards, 'BoardULID', board.BoardULID)
      if (board) {
        // update
        this.$set(this.boards, index, board)
      } else {
        // // remove
        // this.boards.splice(tabindex, 1)
        // this.changeBoard(tabindex)
      }
    },

    fmtDate,
    ...mapActions({
      setButtons: 'app/setButtons'
      // hideFooter: 'app/hideFooter'
    }),

    archive (board) {
      let boardID = -1
      if (board) {
        boardID = board.BoardID
      } else {
        boardID = this.currentBoard.BoardID
      }
      let index = indexByField(this.boardData.records, 'BoardID', boardID)
      this.$service.remove('board', boardID).then(() => {
        this.updateBoards(null, index)
      })
    }
  }
}
</script>
