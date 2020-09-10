export default function () {
  return {
    index: {
      title: 'Home',
      icon: 'home',
      isRouteOnly: true
    },
    pages: {
      title: 'Pages',
      icon: 'file',
      children: {
        'pages-ID-pageEdit': {
          title (instance) {
            return instance.$route.params.pageID === 0 ? 'Create Page' : 'Edit Page'
          }
        },
        'pages-ID-pageLayout': {
          title: 'Edit Page Layout'
        },
        'pages-ID-pageSlides': {
          title: 'Edit Page Slides'
        }
      }
    },
    jobs: {
      title: 'Jobs',
      icon: 'briefcase',
      children: {
        'jobs-ID-jobEdit': {
          isRouteOnly: true,
          title (instance) {
            return instance.$route.params.ID === 0 ? 'Create Job' : 'Edit Job'
          }
        }
      }
    },
    timeentries: {
      title: 'Timesheets',
      icon: 'clock',
      children: {
        'timeentries-ID-timeEntryEdit': {
          title (instance) {
            return instance.$route.params.ID === 0 ? 'Create TimeEntry' : 'Edit TimeEntry'
          }
        }
      }
    },
    boards: {
      title: 'Tasks',
      icon: 'tasks',
      children: {
        // 'boards-ID-boardEdit': {
        //   title (instance) {
        //     return instance.$route.params.ID === 0 ? 'Create Board' : 'Edit Board'
        //   }
        // }
      }
    },
    clients: {
      title: 'Clients',
      icon: 'user',
      children: {
        'clients-ID-clientEdit': {
          isRouteOnly: true,
          title (instance) {
            return instance.$route.params.ID === 0 ? 'Create Client' : 'Edit Client'
          }
        }
      }
    },
    invoices: {
      title: 'Invoice',
      icon: 'money-bill',
      children: {
        'invoices-ID-invoiceEdit': {
          isRouteOnly: true,
          title (instance) {
            return instance.$route.params.ID === 0 ? 'Create Invoice' : 'Edit Invoice'
          }
        }
      }
    },
    people: {
      title: 'Users',
      icon: 'users',
      children: {
        'people-ID-personEdit': {
          isRouteOnly: true,
          title (instance) {
            return instance.$route.params.ID === 0 ? 'Create Person' : 'Edit Person'
          }
        }
      }
    },
    settings: {
      title: 'Settings',
      icon: 'cog'
    },
    login: {
      isRouteOnly: true
    },
    logout: {
      name: 'logout',
      title: 'Logout',
      icon: 'sign-out',
      path: '/logout' // custom path
    }
  }
}
