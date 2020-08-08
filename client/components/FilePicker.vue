<template>
  <b-field class="file">
    <b-upload :value="files" @input="onFileChange">
      <a class="button is-primary">
        <b-icon icon="upload" />
        <span>Upload</span>
      </a>
    </b-upload>
    <a v-if="files && files.length && files[0]" :href="value" target="_blank" class="file-name">
      {{ prettyValue }}
    </a>
    <span v-else class="file-name">
      nothing uploaded
    </span>
  </b-field>
</template>

<script>
export default {
  props: {
    value: String,
    index: Number,
    target: {
      type: String,
      default: '/api/v1/upload/'
    },
    type: {
      type: String,
      default: 'file'
    },
    action: {
      type: String,
      default: 'POST'
    }
  },
  data () {
    return {
      isSelected: false,
      password: ''
    }
  },
  computed: {
    files () {
      return [this.value]
    },
    prettyValue () {
      if (this.value) {
        let valSplit = this.value.split('/attachments/')
        if (valSplit.length >= 2) {
          let newVal = valSplit[valSplit.length - 1]
          if (newVal.indexOf('___') > 0) {
            let parts = newVal.split('___')
            parts.shift()
            newVal = parts.join('')
          }
          return newVal
        }
        return this.value
      }
      return ''
    }
  },
  methods: {
    startUpload () {
      this.$emit('upload-status-changed', 'started')
    },
    finishUpload (resp) {
      if (resp && resp.status === 200) {
        let data = resp.data
        if (data && data.URL) {
          this.$emit('input', data.URL)
        }
      }
      this.$emit('upload-status-changed', 'finished')
    },
    add () {
      this.$refs.uploader.click()
      return false
    },
    change () {
      this.$refs.uploader.click()
      return false
    },
    remove () {
      this.$emit('input', '')
      return false
    },

    onFileChange (file) {
      // console.log(file)
      // let vm = this
      if (!this.target || this.target === '') {
        console.warn('Please provide the target url')
      } else if (!this.action || this.action === '') {
        console.warn('Please provide file upload action ( POST / PUT )')
      } else if (this.action !== 'POST' && this.action !== 'PUT') {
        console.warn('File upload component only allows POST and PUT Actions')
      } else {
        if (!file) {
          return
        };

        /* global FormData:true */
        /* eslint no-undef: "error" */
        this.file = file
        let formData = new FormData()
        formData.append('file', this.file)

        this.startUpload()
        this.$axios.post(this.target + this.type, formData)
          .then((data) => {
            this.finishUpload(data)
          })

        // var xhr = new XMLHttpRequest()
        // xhr.open(this.action, this.target)
        // xhr.onloadstart = function (e) {
        //   vm.startUpload(e)
        // }
        // xhr.onloadend = function (e) {
        //   vm.finishUpload(e)
        // }
        // xhr.send(formData)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
  input[name="fileUpload"] {
    width: 1px;
    height: 1px;
    opacity: 0;
  }
  .with-border{
    border: 1px solid #e6e6e6;
  }

  .image {
    max-width: 120px;
  }

  .image-label {
    font-size: 11px;
    padding-top: 4px;
  }
</style>
