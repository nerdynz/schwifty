// import * as Buefy from 'buefy'
// import Vue from 'vue'

// export default function ({ app }, inject) {
//   Vue.use(Buefy, {
//     defaultIconComponent: 'fa',
//     defaultIconPack: 'fas',
//     customIconPacks: {
//       fas: {
//         sizes: {
//           default: 'lg',
//           'is-small': '1x',
//           'is-medium': '2x',
//           'is-large': '3x'
//         },
//         iconPrefix: ''
//       }
//     }
//   })

//   inject('dialog', Buefy.Dialog)
//   inject('loader', Buefy.LoadingProgrammatic)
//   inject('modal', Buefy.ModalProgrammatic)
//   inject('snackbar', Buefy.SnackbarProgrammatic)
//   inject('toast', Buefy.ToastProgrammatic)
// }
import * as Buefy from 'buefy'
import Vue from 'vue'

export default function ({ app }, inject) {
  Vue.use(Buefy.default, {
    defaultIconComponent: 'fa',
    defaultIconPack: 'far'
  })
  inject('dialog', Buefy.Dialog)
  inject('loader', Buefy.LoadingProgrammatic)
  inject('modal', Buefy.ModalProgrammatic)
  inject('snackbar', Buefy.SnackbarProgrammatic)
  inject('toast', Buefy.ToastProgrammatic)
}
