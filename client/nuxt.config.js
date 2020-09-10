const path = require('path') // eslint-disable-line
const projectRoot = path.resolve(__dirname, '../')

export default {
  mode: 'spa',
  /*
  ** Headers of the page
  */
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress-bar color
  */
  loading: { color: 'blue', height: '400px' },
  /*
  ** Global CSS
  */
  css: [
  ],
  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
    '~/plugins/custom-controls',
    '~/plugins/custom-router',
    '~/plugins/axios-interceptors',
    '~/plugins/menu'
  ],
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
    // Doc: https://github.com/nuxt-community/eslint-module
    '@nuxtjs/eslint-module'
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
    ['specky-service', { isUsingULIDAsPK: true }],
    // Doc: https://buefy.github.io/#/documentation
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    '@nuxtjs/proxy',
    '@nuxtjs/pwa',
    // '@nuxtjs/apollo',
    'nuxt-fontawesome'
  ],
  fontawesome: {
    component: 'fa',
    imports: [
      // import 2 icons from set
      // please note this is PRO set in this example,
      // you must have it in your node_modules to actually import
      {
        set: '@fortawesome/pro-regular-svg-icons',
        icons: ['faBriefcase', 'faArchive']
      }
    ]
  },
  server: {
    port: 3000
  },
  build: {
    babel: {
      configFile: './babel.config.js'
    },
    /*
    ** Run ESLint on save
    */
    extend (config, ctx) {
      config.resolve.alias['~assets'] = path.resolve(projectRoot, '/assets')
    }
  },
  env: {
    // apiURL: (process.env.NODE_ENV === 'production') ? 'https://api.schwifty.co.nz/api' : 'https://server.dev.nerdy.co.nz/api',
    // schemaURL: (process.env.NODE_ENV === 'production') ? 'https://api.schwifty.co.nz/api/v1/schema' : 'https://server.dev.nerdy.co.nz/api/v1/schema',
    isUsingULIDAsPK: true
  },
  axios: {
    browserBaseURL: '/'
  },
  proxy: [
    ['/api/', {
      target: 'http://localhost:5000/api/',
      changeOrigin: true,
      pathRewrite: {
        '^/api': ''
      }
    }],
    ['/fonts/', {
      target: 'http://localhost:5000/fonts/',
      changeOrigin: true,
      pathRewrite: {
        '^/fonts': ''
      }
    }],
    ['/attachments/', {
      target: 'http://localhost:5000/attachments/',
      changeOrigin: true,
      pathRewrite: {
        '^/attachments': ''
      }
    }],
    ['/firebase-messaging-sw.js', {
      target: 'http://localhost:5000/',
      changeOrigin: true
    }]
  ]
}
