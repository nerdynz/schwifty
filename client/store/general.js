// GENERAL IS STORED IN LOCALSTORAGE
// import moment from 'moment'

export const state = () => ({
  // agendaTasks: [],
  // clients: [],
  // users: [],
  currentBoardULID: '',
  currentBoardTaskULID: ''
  // currentAgendaTaskULID: '',
  // currentAgendaDate: moment().format('YYYYMMDD')
})

export const mutations = {
  // CURRENT_AGENDA_TASKS (state, tasks) {
  //   state.agendaTasks = tasks
  // },
  // CURRENT_AGENDA_DATE (state, dateStr) {
  //   state.currentAgendaDate = dateStr
  // },
  CURRENT_SELECTED_BOARD_TASK (state, ulid) {
    state.currentBoardTaskULID = ulid
  },
  CURRENT_SELECTED_AGENDA_TASK (state, ulid) {
    state.currentAgendaTaskULID = ulid
  },
  CURRENT_SELECTED_BOARD (state, ulid) {
    state.currentBoardULID = ulid
  }
  // UPDATE_CLIENTS (state, c) {
  //   state.clients = c
  // },
  // UPDATE_USERS (state, u) {
  //   state.users = u
  // }
}

export const getters = {
  // currentAgendaTasks (state) {
  //   return state.agendaTasks
  // },
  // currentAgendaDate (state) {
  //   return state.currentAgendaDate
  // },
  currentBoardULID (state) {
    return state.currentBoardULID
  },
  currentBoardTaskULID (state) {
    return state.currentBoardTaskULID
  },
  currentAgendaTaskULID (state) {
    return state.currentAgendaTaskULID
  }
}

export const actions = {
  // loadAgenda ({commit, state}) {
  //   var qs = ''
  //   var today = moment().format('YYYYMMDD')
  //   if (moment(state.currentAgendaDate).isBefore(moment(today))) {
  //     qs = '?istoday=1'
  //     commit('CURRENT_AGENDA_DATE', today)
  //   } else if (state.currentAgendaDate === today) {
  //     qs = '?istoday=1'
  //   }
  //   this.$axios.get(`/api/v1/task/retrieve/agenda/${state.currentAgendaDate}` + qs).then(({data}) => {
  //     if (data) {
  //       commit('CURRENT_AGENDA_TASKS', data)
  //     } else {
  //       commit('CURRENT_AGENDA_TASKS', [])
  //     }
  //     // this.$nextTick(() => {
  //     //   this.selectTask(this.currentAgendaTaskULID)
  //     // })
  //   })
  // },
  // loadClients ({commit}) {
  //   this.$axios.get('/api/v1/client/retrieve').then(({data}) => {
  //     commit('UPDATE_CLIENTS', data)
  //   })
  // },
  // loadUsers ({commit}) {
  //   this.$axios.get('/api/v1/people/retrieve').then(({data}) => {
  //     commit('UPDATE_USERS', data)
  //   })
  // }
}
