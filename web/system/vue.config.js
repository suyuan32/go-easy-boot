const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
      /* auto open browser */
      open: true,
      /* use proxy to support cross-origin requests*/
      proxy: {
          '/api': {
              target: 'http://localhost:8500',
              /* allow cross-origin requests*/
              changeOrigin: true,
              ws: false,
              pathRewrite: {
                  '^/api': '/'
            }
          },
      },
  },

})
