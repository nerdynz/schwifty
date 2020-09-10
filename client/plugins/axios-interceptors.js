export default function ({ app, store, env }) {
  app.$axios.interceptors.request.use(function (config) {
    if (env.apiURL) {
      if (config.url.indexOf(env.apiURL) !== 0) {
        config.url = config.url.replace('/api', env.apiURL)
      }
      config.url = config.url.replace('API_URL', env.apiURL)
    } else {
      config.url = config.url.replace('API_URL', '/api')
    }

    if (store.state.auth.isValid) {
      let token = store.state.auth.details.token
      config.headers.Authorization = `Basic ${token}`
    }

    // let isSilent = false
    if (config.url.indexOf('~') > 0) {
      // indicates silence flag
      config.url = config.url.replace('~', '')
      // isSilent = true
    } else {
      // SHHHHHH
      store.commit('app/LOADING_AXIOS', true)
    }
    // Do something before request is sent
    // if (config.showProgress === false) { // if undefined then we still do it
    //   nprogress.start()
    // }
    return config
  }, function (error) {
    store.commit('app/LOADING_AXIOS', false)
    // console.log(error)
    // Do something with request error
    return Promise.reject(error)
  })

  app.$axios.interceptors.response.use(function (response) {
    // console.log(response)
    // Do something with response data
    // nprogress.done()
    store.commit('app/LOADING_AXIOS', false)

    return response
  }, function (error) {
    console.log(error) // eslint-disable-line
    store.commit('app/LOADING_AXIOS', false)
    if (!error.response) {
      app.$toast.open({ message: error.message, type: 'is-danger', duration: 5000 })
      return
    }

    let response = error.response
    if (response && (response.status === 403 || (response.data && response.data.indexOf && response.data.includes('ciphertext too short')))) {
      // todo notify
      // app.$notify({title: `${response.status} ${response.statusText}`, message: response.body, type: 'error'})
      app.router.replace({ name: 'login' })
    }
    if (response.status !== 200) {
      let errorData = response.body
      if (response.data) {
        errorData = response.data
      }

      // Vue.rollbar.debug(errorData)
      try {
        errorData = JSON.parse(errorData)
      } catch (e) {
      }

      // console.log(errorData)
      let message = errorData
      if (errorData && errorData.Friendly) {
        message = errorData.Friendly
      }

      app.$snackbar.open({ message, type: 'is-danger', duration: 5000, position: 'is-top' })
    }
    // Do something with request error
    return Promise.reject(error)
  })
}
