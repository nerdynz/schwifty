<template>
  <header class="u-dp app-navbar animated dont-print" :class="{ slideInDown: show, slideOutDown: !show }" :style="{'background-color': settings.NavColor}">
    <div class="navbar-logo">
      <nuxt-link class="" :to="{path: '/'}">
        <img :src="logo" alt="">
      </nuxt-link>
    </div>
    <nav class="navbar main-nav">
      <div class="navbar-start is-flex-touch">
        <div v-if="checkAllowed(item)" v-for="(item, index) in menu" :key="index" class="navbar-item">
          <nuxt-link class="button is-primary" :to="{path: item.path}" v-if="item.path" :class="{'only-if-exact': item.meta && item.meta.isExactOnly}">
            <span class="icon"><fa :icon="['far', item.icon]" /></span>
            <span class="menu-label-text">{{ getLabel(item) }}</span>
          </nuxt-link>
          <div v-else-if="item.menuLabel" class="menu-label">
            {{ item.menuLabel }}
          </div>
        </div>
      </div>
      <div class="navbar-end" />
    </nav>
    <nav class="navbar sub-nav" v-if="parent" :style="{'background-color': settings.NavColor}">
      <div class="navbar-start is-flex-touch">
        <div v-if="checkAllowed(item)" v-for="(item, index) in parent.children" :key="index" class="navbar-item">
          <nuxt-link class="button is-small" :to="{name: item.name}">
            <span class="menu-label-text">{{ getLabel(item) }}</span>
          </nuxt-link>
        </div>
      </div>
      <div class="navbar-end">
        <div v-for="(item, index) in specialButtons" :key="index" class="level-item">
          <button class="button is-small" :class="'is-' + item.kind" :type="item.kind" @click="(ev) => { item.click(ev) }" :disabled="item.isDisabled || !item.click">
            {{ item.text }}
          </button>
        </div>
      </div>
    </nav>
  </header>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  components: {
  },

  props: {
    show: Boolean
  },

  computed: {
    ...mapGetters({
      current: 'app/current',
      parent: 'app/parent',
      isDev: 'app/isDev',
      settings: 'app/settings',
      logo: 'app/logo',
      pkginfo: 'pkg',
      sidebar: 'app/sidebar',
      specialButtons: 'app/specialButtons'
    }),
    menu () {
      return this.$menu.items
    }
  },
  created () {
  },
  methods: {
    ...mapActions({
      toggleSidebar: 'app/toggleSidebar'
    }),

    checkAllowed (item) {
      let allowed = true
      if (item) {
        if (item.isRouteOnly) {
          return false
        }
        if (item.role) {
          allowed = item.role === this.user.Role
        }
        if (allowed && typeof (item.isAllowed) === 'function') {
          return item.isAllowed(this)
        }
        if (typeof (item.path) === 'undefined' && typeof (item.label) === 'undefined') {
          return false
        }
      }
      return allowed
    },

    getLabel (item) {
      if (item.title) {
        if (typeof (item.title) === 'function') {
          return item.title(this)
        }
        return item.title
      }
    }
  }
}
</script>
