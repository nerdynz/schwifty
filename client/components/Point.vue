<template>
  <div class="point" :class="klass">
    <span v-if="isCheck" class="check-item" @click="checkkkkkk">
      <transition>
        <span v-show="point.IsDone" class="icon ">
          <i class="fa fa-check" aria-hidden="true" />
        </span>
      </transition>
    </span>

    <textarea
      v-model="point.Description"
      contenteditable
      ref="checkInput"
      cols="2"
      class="subtitle-input point-input note-input"
      @input="adjustNoteHeight"
      @blur="blur"
      @keydown.enter="addPoint"
    />

    <a tabindex="-1" class="remove-button" @click.prevent="remove">
      <i class="fa fa-times" />
    </a>
  </div>
</template>

<script>
// import {PointService} from 'services'
// import TitleInput from 'components/TitleInput.vue'

export default {
  name: 'Board',
  props: {
    index: Number,
    point: Object,
    selectedIndex: Number
  },
  data () {
    return {

    }
  },
  computed: {
    klass () {
      return this.isNote ? 'is-note' : 'is-check'
    },
    isNote () {
      return this.point.Type === 'Note'
    },
    isCheck () {
      return this.point.Type === 'Check'
    }
  },

  mounted () {
    if (this.isCheck && this.point.PointID === 0) {
      this.$refs.checkInput.focus()
    } else if (this.isNote) {
      this.$refs.noteInput.focus()
    }
    this.adjustNoteHeight()
  },
  methods: {
    addPoint (ev) {
      if (ev.metaKey || ev.ctrlKey) {
        this.$emit('add')
      }
    },
    checkkkkkk () {
      this.point.IsDone = !this.point.IsDone
      this.save()
    },
    blur () {
      if (this.point.Description) {
        this.save()
      } else {
        this.remove()
      }
    },
    save () {
      this.$emit('save', this.point, this.index)
    },
    remove () {
      this.$emit('remove', this.point, this.index)
    },
    adjustNoteHeight () {
      let input = this.$refs.checkInput
      input.style.height = 0
      let height = input.scrollHeight
      if (height < 28) {
        height = 28
      }
      input.style.height = height + 'px'
      this.$emit('redraw')
    },
    selected () {
    },
    deselected () {
    }
  }
}
</script>

<style lang="scss">

</style>
