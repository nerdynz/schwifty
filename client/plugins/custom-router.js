export default function ({ app, store }, inject) {
  logRoutes(app)
  // add this to a router, make it automatically set bread and current on change
  app.router.beforeEach((to, from, next) => {
    console.log('route changed')
    if (to.path === '/logout') {
      app.$toast.open({ message: 'You have been logged out', type: 'is-info' })
      app.router.replace({ path: 'login' })
      store.commit('auth/LOGOUT')
    }

    // store.commit('app/SET_BUTTONS', [])
    store.commit('app/LOADING_ROUTE', true)
    store.commit('app/SET_SUBTITLE', '')
    store.commit('app/SET_SUBTITLE', null)
    if (from && from.path && from.name) {
      store.commit('app/SET_PREVIOUS_ROUTE', {
        name: from.name,
        path: from.path
      })
    }
    next()
  })

  app.router.afterEach((to, from) => {
    setTimeout(() => {
      store.commit('app/LOADING_ROUTE', false)
    }, 300)
  })
}

function logRoutes (app) {
  let routeInfo = '=== APPLICATION ROUTES ===\n'
  app.router.options.routes.forEach((route) => {
    routeInfo += `${route.name} => ${route.path}\n`
  })
  console.log(routeInfo) // eslint-disable-line
}
