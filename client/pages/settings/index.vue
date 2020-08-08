<template>
  <setting-edit />
</template>
<script>
import SettingEdit from './settingEdit.vue'

export default {
  components: {
    SettingEdit
  },
  data () {
    return {
      setting: {}
    }
  }
}
</script>
