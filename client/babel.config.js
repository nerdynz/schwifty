module.exports = function (api) {
  api.cache(true)

  return {
    presets: [
      [
        '@nuxt/babel-preset-app',
        {
          corejs: { version: 3 }
        }
      ]
    ]
  }
}
