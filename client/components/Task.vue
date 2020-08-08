<template>
  <div>
    <drop @drop="dropTaskOnTask">
      <drag @drag="draggingTask" @dragend="dragTaskCleanUp" :transfer-data="{task: task, index: index, type: 'task', status: ''}">
        <div class="task" :class="[{'is-active': isShowing, 'is-inactive': !isShowing, 'is-new fade-in animated': isNew}, 'is-' + statusKebab]">
          <span v-if="task.BoardTitle" class="task-board-title">
            {{ task.BoardTitle }}
          </span>
          <div class="task-desc level -u-m" @click="selectTask">
            <div class="level-left">
              <div class="level-item">
                <title-input ref="taskTitle" type="text" v-model="task.Title" @blur="save" />
              </div>
            </div>
            <div class="level-right">
              <div class="level-item">
                <div class="status-control field has-addons">
                  <p class="control">
                    <a class="button is-small bs-todo" @click.stop="statusChange('Todo')" :class="{'is-active-bs': task.Status === 'Todo'}">
                      To Do
                    </a>
                  </p>
                  <p class="control">
                    <a class="button is-small bs-progress" @click.stop="statusChange('In Progress')" :class="{'is-active-bs': task.Status === 'In Progress'}">
                      In Progress
                    </a>
                  </p>
                  <p class="control">
                    <a class="button is-small bs-feedback" @click.stop="statusChange('Needs Feedback')" :class="{'is-active-bs': task.Status === 'Needs Feedback'}">
                      Needs Feeback
                    </a>
                  </p>
                  <p class="control">
                    <a class="button is-small bs-done" @click.stop="statusChange('Done')" :class="{'is-active-bs': task.Status === 'Done'}">
                      Done
                    </a>
                  </p>
                </div>
              </div>
            </div>
            <a tabindex="-1" href="#" class="remove-button" @click.prevent.stop="removeTask">
              <i class="far fa-times" />
            </a>
          </div>
        </div>
      </drag>
    </drop>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import { RECORD_TASK } from '~/helpers/consts.js'
import { fmtKebab } from '~/helpers/format.js'
import { getFileFromEvent, getBase64 } from '~/helpers/file.js'
import TitleInput from '~/components/TitleInput'
// import Comments from '~/components/Comments'
// import Point from '~/components/Point'

export default {
  name: 'Task',
  components: {
    TitleInput
  },
  props: {
    index: Number,
    task: Object,
    selected: String,
    forceUpdate: Boolean,
    isAgenda: Boolean
  },
  data () {
    return {
      isShowing: false,
      commentBoxHeight: 380,
      selectedPointIndex: -1
    }
  },
  computed: {
    ...mapGetters({
      authToken: 'auth/usertoken'
    }),
    isNew () {
      return this.task.TaskID === 0
    },
    statusKebab () {
      return fmtKebab(this.task.Status)
    },
    checklist () {
      if (this.task && this.task.Points) {
        return this.task.points.filter((pt) => { return pt.Type === 'Check' })
      }
      return []
    }
  },
  watch: {
    selected (selULID) {
      if (selULID === this.task.TaskULID) {
        this.wasSelected()
      } else if (this.isShowing) { // close it because its not the current one
        this.toggle()
      }
    }
  },
  created () {
    if (this.task._IsNew) {
      this.focusTask()
    }
    // console.log('created', this.task._IsNew)
  },
  updated () {
  },
  beforeDestroy () {
  },
  methods: {
    ...mapActions({
      playDing: 'app/playDing'
    }),
    dropTaskOnTask (xferData, ev) {
      ev.preventDefault()
      let file = getFileFromEvent(ev)
      if (file) {
        let efile = this.$service.new('efile')
        getBase64(file).then((base64) => {
          efile.Name = file.name
          efile.Size = file.size
          efile.Type = file.type
          efile.LastModified = file.LastModified
          // efile.Path =
          efile.Data = base64
          efile.RecordID = this.task.TaskID
          efile.RecordType = RECORD_TASK
          this.$service.save('EFile', efile).then((file) => {
            if (!this.task.Files) {
              this.task.Files = []
            }
            this.task.Files.push(file)
          })
        })
      }

      if (xferData) {
        let task = xferData.task
        let otherTask = task
        if (otherTask.Status === this.task.Status) {
          // let from = index
          let from = otherTask.ULID
          // let to = this.index
          let to = this.task.ULID // we dropped onto this task
          this.$emit('task-sort', { from, to })
        }
      }
    },
    draggingTask (na, ev) {
      if (this.isAgenda) {
        ev.preventDefault()
        return false
      }
      // console.log(ev)
    },
    dragTaskCleanUp (xfer, ev) {
      if (xfer.status === 'moved') {
        this.$emit('task-updated', null, this.task.TaskULID)
      }
    },
    osType () {
      return (window.platform.os.family.Includes('OS X') >= 0 ? 'cmd' : 'ctrl') + '+enter'
    },

    makeActionPoint (message) {
      this.addPoint('Check', message)
    },

    addPoint (type, message) {
      let points = []
      if (this.task.Points) {
        points = this.task.Points
      }
      let point = this.$service.new('point')
      if (message) {
        // point.MessageID = message.MessageID
        point.Description = message.Message
      }
      point.Type = type
      point.TaskID = this.task.TaskID
      points.push(point)
      this.task.Points = points
      this.selectedPointIndex = (points.length - 1)
      // this.recalcCommentsHeight()
    },
    savePoint (point, index) {
      this.$service.save('point', point).then((newPoint) => {
        this.$set(this.task.Points, index, newPoint)
        // this.update(this.task)
      })
    },
    removePoint (point, index) {
      this.$service.remove('point', point).then(() => {
        this.task.Points.splice(index, 1)
      })
    },
    getPointRef () {
      return 'point-' + (this.tabindex)
    },

    selectTask () {
      if (this.isShowing) {
        this.$emit('task-selected', '')
        return
      }
      this.$emit('task-selected', this.task.TaskULID)
    },

    recalcCommentsHeight () {
      setTimeout(() => {
        let actionPoints = this.$refs.actionPoints
        // let notes = this.$refs.notes // jc removed
        // console.log('ap', actionPoints.clientHeight, 'nts', notes.clientHeight)
        // let height = actionPoints.clientHeight + notes.clientHeight - 10
        let height = actionPoints.clientHeight - 10
        if (height > 380) {
          this.commentBoxHeight = height
        }
      }, 10)
    },

    notesChanged (ev, cb) {
      // this.task.Notes = ev.target.innerHTML
      // this.recalcCommentsHeight()
      // if (cb) cb()
    },

    statusChange (status, ev) {
      let sameStatus = status === this.task.Status
      if (sameStatus && status !== 'In Progress') {
        // return early
        return
      }

      // if (status === 'In Progress') {
      //   var x = ev.pageX - ev.target.getBoundingClientRect().left
      //   var width = ev.target.offsetWidth
      //   var percent = Math.round(x / width * 100)
      //   if (this.task.Percent === 0 && !sameStatus) {
      //     this.task.Percent = 10
      //   }

      //   if (percent && percent > 0 && sameStatus) {
      //     this.task.Percent = Math.ceil(percent / 10) * 10
      //   }
      // }

      if (status === 'Done') {
        this.playDing()
      }
      this.task.Status = status
      this.save()
    },

    scrollMessages () {
      setTimeout(() => {
        // this.$refs.commentsScroller.scrollTop = this.$refs.commentsScroller.scrollHeight
      }, 10)
    },
    updateMessages (message) {
      let task = this.task
      if (!task.Actionables) {
        task.Actionables = []
      }
      task.Actionables.push(message)
      this.scrollMessages()
    },

    update (newTask) {
      this.$emit('task-updated', newTask, newTask.ULID)
    },

    save () {
      this.$service.save('task', this.task).then((newTask) => {
        this.update(newTask) // always update even though the published event is gonna happen because we need to update the id
      })
    },

    removeTask () {
      this.$emit('task-removed', this.task, this.task.ULID)
    },

    focusTask () {
      if (this.$refs.taskTitle) {
        this.$refs.taskTitle.select()
      } else {
        this.$nextTick(() => {
          if (this.$refs.taskTitle) {
            this.$refs.taskTitle.select()
          }
        })
      }
    },

    titleClick () {

    },

    // animation
    toggle () {
      // this.recalcCommentsHeight()
      let self = this
      self.isShowing = !self.isShowing
      if (self.isShowing) {
        if (self.task.TaskID > 0) { // not new
          self.$service.retrieve('task', self.task.TaskID).then((newTask) => {
            this.update(newTask)
          })
        }
        // this.subscribe()
      } else {
        // this.unsubscribe()
      }
    },
    cancel (targets) {
    },
    before (target) {
      target.style.maxHeight = '0px'
      target.style.minHeight = '0px'
    },
    enter (target, done) {
      target.style.maxHeight = `${target.scrollHeight}px`
      target.style.opacity = 1
      setTimeout(() => {
        if (this.onEnter) {
          this.onEnter()
        }
        done()
      }, 600)
    },
    after (target, done) {
      target.style.maxHeight = 'initial'
    },
    beforeLeave (target) {
      target.style.maxHeight = `${target.scrollHeight}px`
    },
    leave (target, done) {
      setTimeout(() => {
        target.style.maxHeight = '0px'
      }, 1)
      // target.style.opacity = 0
      setTimeout(() => {
        done()
      }, 600)
    },
    wasSelected () {
      if (this.task.TaskID <= 0) {
      } else {
        this.onEnter = function () {
          // this.$refs.comments.focus()
          this.scrollMessages()
        }
      }
      this.toggle()
    }
  }
}
</script>
