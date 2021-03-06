module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    parser: 'babel-eslint'
  },
  extends: [
    '@nuxtjs',
    'plugin:nuxt/recommended',
    'standard'
  ],
  // add your custom rules here
  rules: {
    'vue/singleline-html-element-content-newline': 0,
    'prefer-const': 0,
    'vue/attributes-order': 0,
    'vue/require-default-prop': 0,
    'no-console': 0
  }
}
