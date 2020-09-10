import { library } from '@fortawesome/fontawesome-svg-core'
import { far } from '@fortawesome/pro-regular-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import * as Buefy from 'buefy'
import Vue from 'vue'
import VueDragDrop from 'vue-drag-drop'
import Multiselect from 'vue-multiselect'
import VueQuill from 'vue-quill'
import EditForm from '~/components/EditForm'
import EditHeader from '~/components/EditHeader'
import Field from '~/components/Field'
import { fmtDate, fmtDateTime } from '~/helpers/format'
import { focusElement } from '~/helpers/helpers'

Vue.component('fa-icon', FontAwesomeIcon)
// library.add(fas)
library.add(far)

Vue.component('multiselect', Multiselect)
Vue.use(VueDragDrop)
Vue.use(VueQuill)

Vue.mixin({
  methods: {
    fmtDate,
    fmtDateTime,
    focusElement
  }
})

Vue.use(require('vue-prevent-parent-scroll'))
// Vue.component('event-calendar', Calendar)
Vue.component('field', Field)
Vue.component('edit-header', EditHeader)
Vue.component('edit-form', EditForm)

Vue.directive('label', {
  bind (el, binding) {
    let text = ''
    let isFloat = false
    let isStuck = false
    if (binding.modifiers && binding.modifiers.float) {
      isFloat = true
    }
    if (binding.modifiers && binding.modifiers.stuck) {
      isStuck = true
    }
    // value passed in directive is string
    if (binding && typeof binding.value === 'string') {
      text = binding.value
    }

    if (binding && typeof binding.value === 'undefined') {
      text = el.getAttribute('label') || ''
    }

    // value passed in directive is object
    if (binding && typeof binding.value === 'object') {
      text = binding.value.label || ''
      if (binding.value.isFloat) {
        isFloat = true
      }
    }

    // do this here to get the right label
    el = getCorrectTag(el)

    Vue.nextTick(() => {
      // anchor span
      let anchorElement = document.createElement('span')
      anchorElement.className = 'label-anchor'
      // anchorElement.style.width = el.clientWidth + 'px'
      if (isFloat || isStuck) {
        anchorElement.classList.add('float-anchor')
      }

      // label
      let label = document.createElement('label')
      label.classList.add('label')
      if (isFloat) {
        label.classList.add('float-label')
      }
      label.innerHTML = text
      label.addEventListener('click', () => {
        el.focus()
      })

      if (isStuck) {
        label.classList.add('float-label')
        label.classList.add('is-pulled-above-stuck')
      }

      if (el.value) {
        label.classList.add('is-pulled-above')
        if (binding.modifiers.hidden) {
          label.classList.add('hidden')
        }
      }

      el.addEventListener('blur', function () {
        setTimeout(() => {
          label.classList.remove('is-highlighted')
          if (el.value) {
            if (
              binding.modifiers.hidden &&
              el.labelEl.className.includes('hidden')
            ) {
              label.classList.add('hidden')
            }
          } else {
            label.classList.remove('is-pulled-above')
            label.classList.remove('hidden')
          }
        }, 10)
      })
      el.addEventListener('focus', function () {
        label.classList.add('is-highlighted')
        label.classList.add('is-pulled-above')
        label.classList.remove('hidden')
      })
      el.labelEl = label
      anchorElement.appendChild(label)
      let parent = el.parentNode
      // parent.insertBefore(anchorElement, el)
      parent.insertBefore(anchorElement, el.nextSibling)
    })
  },
  update (el, binding) {
    el = getCorrectTag(el)
    if (el.labelEl) {
      // lets hope this reference sticks otherwise we have to go find it in the dom????
      setTimeout(() => {
        if (el.value) {
          if (!el.labelEl.className.includes('is-pulled-above')) {
            el.labelEl.classList.add('is-pulled-above')
          }
          if (
            binding.modifiers.hidden &&
            !el.labelEl.className.includes('hidden')
          ) {
            el.labelEl.classList.add('hidden')
          }
        } else {
          el.labelEl.classList.remove('is-pulled-above')
          if (binding.modifiers.hidden) {
            el.labelEl.classList.remove('hidden')
          }
        }
      }, 10)
    }
  }
})

function getCorrectTag (el) {
  if (el.tagName.toLowerCase() === 'input' || el.tagName.toLowerCase() === 'textarea' || el.className.includes('is-float-anchor')) {
    // good we have the right input
  } else {
    let els = el.getElementsByTagName('input')
    if (els.length > 0) {
      el = els[0]
      return el
    }
    els = el.getElementsByTagName('textarea')
    if (els.length > 0) {
      el = els[0]
      return el
    }
  }
  return el
}

export default function ({ app }, inject) {
  Vue.use(Buefy.default, {
    defaultIconComponent: 'fa-icon',
    defaultIconPack: 'far',
    customIconPacks: {
      fas: {
        sizes: {
          default: 'lg',
          'is-small': '1x',
          'is-medium': '2x',
          'is-large': '3x'
        },
        iconPrefix: ''
      }
    }
  })

  inject('dialog', Buefy.Dialog)
  inject('loader', Buefy.LoadingProgrammatic)
  inject('modal', Buefy.ModalProgrammatic)
  inject('snackbar', Buefy.SnackbarProgrammatic)
  inject('toast', Buefy.ToastProgrammatic)
}

// const apolloClient = new ApolloClient({
//   // You should use an absolute URL here
//   uri: 'https://hasura.schwifty.co.nz/v1/graphql',
//   'x-hasura-admin-secret': 'iKzHiovXQ9vU-7YP'
// })

// const apolloProvider = new VueApollo({
//   defaultClient: apolloClient
// })

// Vue.use(VueApollo)
