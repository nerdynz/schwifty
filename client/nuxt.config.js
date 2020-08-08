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
    '@nuxtjs/apollo',
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
  apollo: {
    tokenName: 'schwifty-apollo-token', // optional, default: apollo-token
    cookieAttributes: {
      /**
        * Define when the cookie will be removed. Value can be a Number
        * which will be interpreted as days from time of creation or a
        * Date instance. If omitted, the cookie becomes a session cookie.
        */
      expires: 7, // optional, default: 7 (days)

      /**
        * Define the path where the cookie is available. Defaults to '/'
        */
      path: '/', // optional
      /**
        * Define the domain where the cookie is available. Defaults to
        * the domain of the page where the cookie was created.
        */
      domain: 'example.com', // optional

      /**
        * A Boolean indicating if the cookie transmission requires a
        * secure protocol (https). Defaults to false.
        */
      secure: false
    },
    includeNodeModules: true, // optional, default: false (this includes graphql-tag for node_modules folder)
    authenticationType: 'Basic', // optional, default: 'Bearer'
    // (Optional) Default 'apollo' definition
    defaultOptions: {
      // See 'apollo' definition
      // For example: default query options
      $query: {
        loadingKey: 'loading',
        fetchPolicy: 'cache-and-network'
      }
    },
    // optional
    // watchLoading: '~/plugins/apollo-watch-loading-handler.js',
    // optional
    // errorHandler: '~/plugins/apollo-error-handler.js',
    // required
    clientConfigs: {
      default: {
        // required
        httpEndpoint: 'https://hasura.schwifty.co.nz/v1/graphql',
        // optional
        // override HTTP endpoint in browser only
        browserHttpEndpoint: 'https://hasura.schwifty.co.nz/v1/graphql'
        // optional
        // See https://www.apollographql.com/docs/link/links/http.html#options
        // httpLinkOptions: {
        //   credentials: 'same-origin'
        // },
        // // You can use `wss` for secure connection (recommended in production)
        // // Use `null` to disable subscriptions
        // wsEndpoint: 'ws://localhost:4000', // optional
        // // LocalStorage token
        // tokenName: 'apollo-token', // optional
        // // Enable Automatic Query persisting with Apollo Engine
        // persisting: false, // Optional
        // // Use websockets for everything (no HTTP)
        // // You need to pass a `wsEndpoint` for this to work
        // websocketsOnly: false // Optional
      }
      // test: {
      //   httpEndpoint: 'http://localhost:5000',
      //   wsEndpoint: 'ws://localhost:5000',
      //   tokenName: 'apollo-token'
      // }
    }
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
    apiURL: (process.env.NODE_ENV === 'production') ? 'https://api.schwifty.co.nz/api' : 'https://server.dev.nerdy.co.nz/api',
    schemaURL: (process.env.NODE_ENV === 'production') ? 'https://api.schwifty.co.nz/api/v1/schema' : 'https://server.dev.nerdy.co.nz/api/v1/schema',
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
