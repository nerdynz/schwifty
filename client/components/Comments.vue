<template>
  <edit-form title="Comments" size="small" no-margin class="comments-box">
    <div class="comments-scroller">
      <div class="comment" v-for="message in messages" :key="message.MessageID">
        <div class="comment-details">
          <a href="#" class="make-action" @click.prevent="makeActionPoint(message)">
            <b-icon icon="reply" size="is-small" />
          </a>
          <div class="author">
            {{ message.Author ? message.Author : 'Anonymous' }}
          </div>
          <div class="thin-label">
            {{ time(message) }}
          </div>
        </div>
        <div v-html="addLineBreaks(message.Comment)" />
      </div>
    </div>
    <div class="comment-input">
      <div class="comment-bg">
        <drop class="comment-drag" @drop="drop" @dragover="dragOver">
          <textarea
            x-hint="tooltip"
            rows="3"
            ref="commentInput"
            class="subtitle-input"
            @keydown.enter="addComment"
            @focus="focussed"
            @blur="blurred"
            v-model="newMessage.Comment"
            type="text"
          />
          <button class="button is-primary comment-button" @click="addComment">
            Send
          </button>
        </drop>
      </div>
    </div>
  </edit-form>
</template>

<script>
import { mapGetters } from 'vuex'
import { fmtTime } from '~/helpers/format.js'
export default {
  name: 'Board',
  props: {
    task: Object
  },
  data () {
    return {
      newMessage: {
        Comment: ''
      },
      isFocused: false
    }
  },
  computed: {
    ...mapGetters({
      siteID: 'auth/siteID',
      user: 'auth/user'
    }),
    messages () {
      return this.task.Messages
    },
    osType () {
      return (window.platform.os.family.includes('OS X') ? 'cmd' : 'ctrl') + '+enter'
    },
    tooltip () {
      return {
        content: `<strong>Remember</strong><ul>
                  <li>Detailed comments are better</li>
                  <li>Don't spam</li>
                  <li>Refrain from saying 'Hi', 'Thanks', etc.</li>
                  <li>&nbsp;</li>
                  <li><strong>You can also drag attachments here</strong></li>
                 </ul>`,
        display: this.isFocused && this.newMessage.Comment.length < 6,
        className: 'is-note'
      }
    }
  },
  mounted () {
    this.makeNewMessage()
  },
  beforeDestroy () {
  },
  methods: {
    makeActionPoint (message) {
      this.$emit('make-from', message)
    },

    time (msg) {
      return fmtTime(msg.DateCreated)
    },

    makeNewMessage () {
      this.newMessage = this.$service.new('Message', this.siteID)
      this.newMessage.PersonULID = this.user.ULID
    },

    addComment (ev) {
      let message = this.newMessage
      if (!message.Comment) return // leave early if no message

      if (ev.metaKey || ev.ctrlKey || ev.type === 'click') {
        message.BoardULID = this.task.BoardULID
        message.TaskULID = this.task.TaskULID
        this.$service.save('~Message', message).then((message) => {
          this.$emit('message-added', message)
        })
        this.makeNewMessage()
      }
    },

    labelClicked () {
      this.focus()
    },
    focus () {
      this.$refs.commentInput.focus()
    },
    focussed () {
      this.isFocused = true
    },
    blurred () {
      this.isFocused = false
    },
    addLineBreaks (txt) {
      txt = txt + ''
      return txt.replace(/\n/g, '<br />')
    },
    dragOver (na, ev) {
    },
    drop (na, ev) {
      ev.preventDefault()
    }
  }
}
</script>
