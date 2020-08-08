<template>
  <div>
    <b-loading :active="isLoading" />
    <div v-if="board.BoardULID">
      <div class="board-settings u-p">
        <div class="level">
          <div class="level-left">
            <div class="level-item">
              <title-input ref="BoardTitle" v-model="board.Title" @blur="update" />
            </div>
            <div class="level-item">
              <color-picker v-model="board.Color" @input="colorChanged" />
            </div>
          </div>
          <div class="level-right">
            <div v-if="false" class="level-item">
              <multiselect
                :value="peopleOnThisBoard"
                @input="updatePeopleOnBoard"
                placeholder="Users"
                :options="users"
                track-by="PersonID"
                label="Initials"
                multiple
                class="person-selector"
              />
            </div>
          </div>
        </div>
      </div>
      <div class="columns">
        <div class="column is-8">
          <div ref="taskSorter" class="task-list">
            <div v-if="board && board.Tasks && board.Tasks.length > 0">
              <task
                v-for="(task, index) in board.Tasks"
                :task="task"
                :ref="task.TaskULID"
                :index="index"
                :selected="selectedTask"
                :key="task.TaskULID"
                @task-updated="updateTasks"
                @task-removed="removeTask"
                @task-selected="selectTask"
                @task-sort="sortTasks"
                @task-moved="taskMoved"
              />
            </div>
            <div v-else class="no-data">
              <svg viewBox="0 0 80 80" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
                <defs />
                <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
                  <g id="folder-smile" fill-rule="nonzero" fill="#000000">
                    <path d="M69.999,6.667 L40,6.667 L33.333,0 L10,0 C8.167,0 6.667,1.497 6.667,3.333 L6.667,13.333 L73.333,13.333 L73.333,10 C73.332,8.164 71.832,6.667 69.999,6.667 Z" id="Shape" />
                    <path d="M74.338,20 L5.661,20 C1.995,20 -0.511,22.956 0.091,26.575 L8.994,80 L71.005,80 L79.908,26.575 C80.514,22.956 78.007,20 74.338,20 Z M49.999,37.5 C52.301,37.5 54.166,39.362 54.166,41.667 C54.166,43.965 52.301,45.833 49.999,45.833 C47.697,45.833 45.832,43.965 45.832,41.667 C45.832,39.362 47.697,37.5 49.999,37.5 Z M30,37.5 C32.302,37.5 34.167,39.362 34.167,41.667 C34.167,43.965 32.302,45.833 30,45.833 C27.698,45.833 25.833,43.965 25.833,41.667 C25.833,39.362 27.698,37.5 30,37.5 Z M56.666,58.9658752 L53.131,62.5008752 C49.511,58.8808752 44.755,57.0708752 39.999,57.0708752 C35.244,57.0708752 30.488,58.8808752 26.868,62.5008752 L23.333,58.9658752 C32.542,49.7598752 47.467,49.7668752 56.666,58.9658752 Z" id="Shape" />
                  </g>
                </g>
              </svg>
              <h3 class="title is-3">
                Oops. There is nothing here!
              </h3>
              <button @click="addTask" class="button is-primary">
                Add New Task
              </button>
            </div>
          </div>
        </div>
        <div class="column is-4">
          SIDEBAR TODO
        </div>
      </div>

      <b-modal
        :active="hasSelectedTask"
        @close="closeTaskModal"
        has-modal-card
        trap-focus
        :destroy-on-hide="true"
        aria-role="dialog"
        aria-modal
      >
        <task-detail :task-ulid="selectedTask" />
      </b-modal>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
// import { Collapse, Item as CollapseItem } from 'vue-bulma-collapse'
// import PersonSelector from '~/components/PersonSelector.vue'
import TitleInput from '~/components/TitleInput.vue'
import ColorPicker from '~/components/ColorPicker.vue'
import Task from '~/components/Task.vue'
import TaskDetail from '~/components/TaskDetail.vue'
import { indexByULID, changeSortReverse } from '~/helpers/filters.js'

// import ListPlaceholder from '~/components/layout/ListPlaceholder.vue'

export default {
  components: {
    // PersonSelector,
    // ListPlaceholder,
    // Collapse,
    // CollapseItem,
    TitleInput,
    Task,
    TaskDetail,
    ColorPicker
  },
  props: {
    board: Object,
    tabindex: Number,
    selected: Number,
    redraw: Boolean,
    forceUpdate: Boolean,
    selectedTaskUlid: String,
    selectedBoardUlid: String
  },
  data () {
    return {
      isSelected: false,
      selectedTask: '',
      originalTitle: ''
      // board: {
      //   ulid: ''
      // }
    }
  },
  computed: {
    hasSelectedTask () {
      return this.selectedTask !== ''
    },
    ...mapGetters({
      users: 'general/users',
      currentBoardULID: 'general/currentBoardULID',
      currentBoardTaskULID: 'general/currentBoardTaskULID',
      isLoading: 'app/isLoading'
    }),
    // todoTasks () {
    //   return tasksByStatus(this.board.Tasks, 'Todo')
    // },
    // progressTasks () {
    //   return tasksByStatus(this.board.Tasks, 'In Progress')
    // },
    // feedbackTasks () {
    //   return tasksByStatus(this.board.Tasks, 'Needs Feedback')
    // },
    // doneTasks () {
    //   return tasksByStatus(this.board.Tasks, 'Done')
    // },
    peopleOnThisBoard () {
      if (this.board && this.board.People) {
        return this.users.filter((u) => {
          return this.board.People.includes(u.PersonID)
        })
      }
      return []
    }
  },
  created () {
    // this.loadBoard()
  },
  methods: {
    closeTaskModal () {
      this.selectedTask = ''
    },
    updatePeopleOnBoard (selectedPeeps) {
      let people = []
      selectedPeeps.forEach((p) => {
        people.push(p.PersonID)
      })
      this.$axios.put(`/api/v1/board/people/${this.board.BoardID}?ids=${people.join('&ids=')}`).then(() => {
      })
      this.board.People = people
    },
    taskMoved (task) {
    },
    focusTitle () {
      this.$refs.BoardTitle.focus()
    },
    // boardSelected () {
    //   let curSelectedTask = this.selectedTask
    //   if (!curSelectedTask) {
    //     curSelectedTask = this.currentBoardTaskULID
    //   }
    //   if (this.isNew) {
    //     this.isSelected = true
    //     this.focusTitle()
    //     this.$emit('board-loaded')
    //   } else {
    //     this.originalTitle = this.board.title
    //     this.$nextTick(() => {
    //       this.$service.retrieve('board', this.board.BoardID, 'task', curSelectedTask).then((newBoard) => {
    //         this.$emit('board-loaded')
    //         this.$emit('board-updated', newBoard, this.actualTabIndex)
    //         this.isSelected = true
    //       })
    //     })
    //   }
    //   // this.subscribeTask('task', this.publishedTask)
    // },

    // boardDeselected () {
    //   this.isSelected = false
    //   // this.unsubscribeTask('task', this.publishedTask)
    // },

    getBoardRef () {
      return 'board-' + (this.tabindex)
    },

    update (val) {
      this.$emit('board-updated', this.board)
      this.saveBoard()
    },
    colorChanged (color) {
      this.board.color = color
      this.$emit('board-updated', this.board)
      this.saveBoard()
    },
    saveBoard (ev, cb) {
      this.$emit('board-save', this.board)
    },

    // tasks
    addTask (status, title) {
      this.$emit('add-task', { status, title })
    },

    getTaskIndex (taskID) {
      return this.board.Tasks.findIndex((task) => {
        return task.TaskID === taskID
      })
    },

    // subscribeTask () {
    //   this.$service.subscribe('task', this.publishedTask)
    // },

    // publishedTask ({data}) {
    //   if (data && data.Type === 'task-updated') {
    //     let taskIndex = this.getTaskIndex(data.Task.TaskID)
    //     if (taskIndex !== -1) {
    //       this.$set(this.board.Tasks, taskIndex, data.Task)
    //     }
    //   }
    //   if (data && data.Type === 'message-added') {
    //     let taskIndex = this.getTaskIndex(data.Message.TaskID)
    //     if (taskIndex !== -1) {
    //       let task = this.board.Tasks[taskIndex]
    //       if (!task.Messages) {
    //         task.Messages = []
    //       }
    //       task.Messages.push(data.Message)
    //       this.$set(this.board.Tasks, taskIndex, task)
    //     }
    //   }
    // },

    // unsubscribeTask () {
    //   this.$service.unsubscribe('task', this.publishedTask)
    // },

    selectTask (ulid) {
      this.selectedTask = ulid
      // this.$store.commit('general/CURRENT_SELECTED_BOARD_TASK', ulid)
      // this.$store.commit('general/CURRENT_SELECTED_BOARD', ulid)
      // this.$router.push({ name: 'boards-ulid', params: { ulid } })
      // window.history.replaceState(null, '', `/boards/${ulid}`)
    },

    updateTasks (task, ulid) {
      let index = indexByULID(this.board.Tasks, ulid)
      if (task) {
        this.$set(this.board.Tasks, index, task)
      } else {
        this.board.Tasks.splice(index, 1)
      }
    },

    removeTask (task, ulid) {
      let index = indexByULID(this.board.Tasks, ulid)
      this.$service.remove('task', task)
      this.board.Tasks.splice(index, 1)
    },

    sortTasks ({ from, to }) {
      if (from !== to) {
        let fromIndex = indexByULID(this.board.Tasks, from)
        let toIndex = indexByULID(this.board.Tasks, to)
        let newRecords = changeSortReverse(this.board.Tasks, fromIndex, toIndex)
        this.board.Tasks = newRecords
        this.$axios.put('/api/v1/task/sort', newRecords)
      }
    }
  }
}
</script>

<style lang="scss">

</style>
