export const state = () => ({
  isCheckingLogin: false,
  isValid: false, // always true because we allow people to view a read only copy of the admin
  details: {
    token: '',
    expiration: '',
    cache: '',
    user: {
      ID: 0,
      Name: 'Anonymous',
      Role: '',
      SiteID: 0
    }
  }
})

export const mutations = {
  CHECKING (state, checking) {
    state.isCheckingLogin = checking
  },
  LOGIN (state, userData) {
    state.details = userData
    state.isValid = true
  },
  LOGOUT (state) {
    state.details = {
      token: '',
      expiration: '',
      cache: '',
      user: {
        Name: 'Anonymous',
        Role: ''
      }
    }
    // state.isValid = false
    state.isCheckingLogin = false
  }
}

export const actions = {
  login ({ commit }, details) {
    let self = this.app
    commit('CHECKING', true)
    self.$axios.post('API_URL/v1/login', details)
      .then(({ data }) => {
        // console.log(userDetails)
        // if (userDetails.Person.Role === 'Technician') {
        //   commit(types.SET_JOB_STATUS, 'Workshop')
        // }
        self.router.replace('/')
        commit('LOGIN', data)
        commit('CHECKING', false)
      })
      .catch(({ data }) => {
        commit('CHECKING', false)
      })
  },

  logout ({ commit }, details) {
    commit('LOGOUT')
  }
}

export const getters = {
  siteID: (state) => {
    return '01EDG1D97AWN9V0Q87E4SJ13C7'
    // return state.details.user.SiteID
  },
  userID: (state) => {
    return state.details.user.ID
  },
  user: (state) => {
    return state.details.user
  },
  userIsDev: (state) => {
    let role = ''
    role += state.details.user.Role
    role = role.toLowerCase()
    return role.includes('developer')
  }
}
