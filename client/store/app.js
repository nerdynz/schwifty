export const state = () => ({
  sitename: 'Schwifty',
  buttons: [

  ],
  title: 'Generic Title',
  subtitle: '',
  breadcrumb: null,
  previousRoute: {
    name: '',
    path: ''
  },
  settings: {
  },
  fileQueue: {
  },
  isLoadingAxios: false,
  isLoadingRoute: false,
  isLoadingGeneral: false
})

export const mutations = {
  BLOCK_HOVERED (state, isHovered) {
    state.blockIsHovered = isHovered
  },
  CURRENT_WINDOW_HEIGHT (state, h) {
    state.windowHeight = h
  },
  CURRENT_WINDOW_WIDTH (state, w) {
    state.windowWidth = w
  },
  SET_PREVIOUS_ROUTE (state, prevRoute) {
    if (prevRoute) {
      state.previousRoute = prevRoute
    } else {
      state.previousRoute = {
        name: '',
        path: ''
      }
    }
  },

  SET_TITLE (state, title) {
    state.title = title
  },

  SET_SUBTITLE (state, subtitle) {
    state.subtitle = subtitle
  },

  SET_SETTINGS (state, settings) {
    let root = document.documentElement
    root.style.setProperty('--primary-color', settings.PrimaryColor)
    root.style.setProperty('--nav-bg-color', settings.PrimaryColor)
    root.style.setProperty('--sec-nav-bg-color', settings.SecondaryColor)
    state.settings = settings
  },

  SET_BREADCRUMB (state, breadcrumb) {
    state.breadcrumb = breadcrumb
  },

  LOADING_AXIOS (state, isLoading) {
    if (!isLoading) {
      state.isLoadingGeneral = false // unload this the first chance we get
    }
    state.isLoadingAxios = isLoading
  },

  LOADING_ROUTE (state, isLoading) {
    if (!isLoading) {
      state.isLoadingGeneral = false // unload this the first chance we get
    }
    state.isLoadingRoute = isLoading
  },

  LOADING_GENERAL (state, isLoading) {
    state.isLoadingGeneral = isLoading
  }
}

export const actions = {
  setButtons ({ commit }, buttons) {
    throw new Error('this sucked use edit / list placeholders')
  }
}

export const getters = {
  isLoading (state) {
    if (state.isLoadingGeneral) {
      return true
    }
    if (state.isLoadingRoute) {
      return true
    }
    if (state.isLoadingAxios) {
      return true
    }
    return false
  },
  settings (state) {
    return state.settings
  },
  breadcrumb (state) {
    return state.breadcrumb
  },
  current (state) {
    if (state.breadcrumb) {
      return state.breadcrumb[state.breadcrumb.length - 1]
    }
    return null
  },
  parent (state) {
    if (state.breadcrumb) {
      if (state.breadcrumb.length === 2) {
        return state.breadcrumb[state.breadcrumb.length - 1]
      }
      return state.breadcrumb[state.breadcrumb.length - 2]
    }
    return null
  },
  title (state) {
    if (state.breadcrumb && state.breadcrumb.length > 0) {
      return state.breadcrumb[state.breadcrumb.length - 1].title
    }
    if (state.title) {
      return state.title
    }
    return 'Generic Title'
  },
  subtitle (state) {
    return state.subtitle
  },
  isDev () {
    return (process.env.NODE_ENV === 'development')
  },
  sitename (state) {
    return state.sitename
  },
  logo (state) {
    let logo = state.settings.LogoPicture || ''
    // if (logo && (process.env.NODE_ENV === 'development')) {
    //   logo = logo.replace('/attachments/', `https://cdn.nerdy.co.nz/attachments/${state.sitename}/`)
    // }
    return logo
  },
  specialButtons () {
    return []
  }
}
